// Copyright (c) 2022 Cisco and/or its affiliates. All rights reserved.
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//       https://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package grpc

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"strings"

	"github.com/go-logr/logr"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/cisco-open/libnasp/pkg/istio/discovery"
	"github.com/cisco-open/libnasp/pkg/network"
	"github.com/cisco-open/libnasp/pkg/proxywasm/api"
	"github.com/cisco-open/libnasp/pkg/proxywasm/http"
)

type GRPCDialer struct {
	http.MiddlewareHandler

	tlsConfig       *tls.Config
	streamHandler   api.StreamHandler
	discoveryClient discovery.DiscoveryClient
	logger          logr.Logger

	connectionState network.ConnectionState
}

func NewGRPCDialer(streamHandler api.StreamHandler, tlsConfig *tls.Config, discoveryClient discovery.DiscoveryClient, logger logr.Logger) *GRPCDialer {
	dialer := GRPCDialer{
		MiddlewareHandler: http.NewMiddlewareHandler(),

		tlsConfig:       tlsConfig,
		streamHandler:   streamHandler,
		discoveryClient: discoveryClient,
		logger:          logger,
	}

	return &dialer
}

func (g *GRPCDialer) Dial(ctx context.Context, addr string) (net.Conn, error) {
	tlsConfig := g.tlsConfig.Clone()

	if prop, _ := g.discoveryClient.GetHTTPClientPropertiesByHost(ctx, addr); prop != nil {
		g.logger.V(3).Info("discovered overrides", "overrides", prop)
		if endpointAddr, err := prop.Address(); err != nil {
			return nil, err
		} else {
			addr = endpointAddr.String()
		}
		g.logger.V(3).Info("address override", "address", addr)
		tlsConfig.ServerName = prop.ServerName()
	}

	ctx = network.NewConnectionStateHolderToContext(ctx)

	opts := []network.DialerOption{
		network.DialerWithConnectionOptions(network.ConnectionWithCloserWrapper(g.discoveryClient.NewConnectionCloseWrapper())),
		network.DialerWithDialerWrapper(g.discoveryClient.NewDialWrapper()),
	}

	conn, err := network.NewDialerWithTLSConfig(tlsConfig, opts...).DialTLSContext(ctx, "tcp", addr)
	if err != nil {
		return nil, err
	}

	if state, ok := network.ConnectionStateFromContext(ctx); ok {
		g.connectionState = state
	}

	return conn, nil
}

func (g *GRPCDialer) RequestInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	g.logger.Info("intercepted request", "target", cc.Target(), "method", method, "req", req)

	if g.discoveryClient != nil {
		g.discoveryClient.IncrementActiveRequestsCount(cc.Target())
		defer g.discoveryClient.DecrementActiveRequestsCount(cc.Target())
	}

	stream, err := g.streamHandler.NewStream(api.ListenerDirectionOutbound)
	if err != nil {
		g.logger.Error(err, "could not get new stream")
		return err
	}
	defer stream.Close()

	headers := metadata.New(map[string]string{
		"content-type": "application/grpc",
	})

	if md, found := metadata.FromOutgoingContext(ctx); found {
		headers = metadata.Join(headers, md)
	}

	var responseHeaders, responseTrailers metadata.MD
	opts = append(opts, grpc.Header(&responseHeaders), grpc.Trailer(&responseTrailers))

	wrappedRequest := WrapGRPCRequest(fmt.Sprintf("https://%s/%s", cc.Target(), strings.TrimLeft(method, "/")), headers, g.connectionState)

	g.BeforeRequest(wrappedRequest, stream)

	if err = stream.HandleHTTPRequest(wrappedRequest); err != nil {
		return err
	}

	g.AfterRequest(wrappedRequest, stream)

	ctx = metadata.NewOutgoingContext(ctx, headers)

	err = invoker(ctx, method, req, reply, cc, opts...)
	if err != nil {
		return err
	}

	if h, ok := network.ConnectionStateHolderFromContext(ctx); ok {
		h.Set(g.connectionState)
	}

	wrappedResponse := WrapGRPCResponse(status.Code(err), responseHeaders, responseTrailers, g.connectionState)
	stream.Set("grpc.status", status.Code(err))

	g.BeforeResponse(wrappedResponse, stream)

	if err = stream.HandleHTTPResponse(wrappedResponse); err != nil {
		return err
	}

	g.AfterResponse(wrappedResponse, stream)

	g.logger.Info("intercepted reply", "method", method, "reply", reply)

	return nil
}
