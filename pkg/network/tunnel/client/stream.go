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

package client

import (
	"encoding/json"

	"emperror.dev/errors"
	"github.com/xtaci/smux"

	"github.com/cisco-open/nasp/pkg/network/tunnel/api"
	"github.com/cisco-open/nasp/pkg/network/tunnel/common"
)

type ctrlStream struct {
	api.ControlStream

	client *client
}

func NewControlStream(client *client, str *smux.Stream) api.ControlStream {
	cs := common.NewControlStream(str, client.logger)

	s := &ctrlStream{
		ControlStream: cs,
		client:        client,
	}

	cs.AddMessageHandler(api.AddPortResponseMessageType, s.addPortResponse)
	cs.AddMessageHandler(api.RequestConnectionMessageType, s.requestConnection)
	cs.AddMessageHandler(api.PingMessageType, s.ping)

	return s
}

func (s *ctrlStream) ping(msg []byte) error {
	s.client.logger.V(3).Info("ping arrived, send pong")

	_, _, err := api.SendMessage(s, api.PongMessageType, nil)
	if err != nil {
		return err
	}

	return nil
}

func (s *ctrlStream) requestConnection(msg []byte) error {
	var req api.RequestConnectionMessage
	if err := json.Unmarshal(msg, &req); err != nil {
		return errors.WrapIf(err, "could not unmarshal requestConnection message")
	}

	var mp *managedPort
	if v, ok := s.client.managedPorts.Load(req.PortID); ok {
		if p, ok := v.(*managedPort); ok {
			mp = p
		}
	}

	if mp == nil {
		return errors.WithStackIf(api.ErrInvalidPort)
	}

	conn, err := s.client.session.OpenTCPStream(req.Identifier)
	if err != nil {
		return errors.WrapIfWithDetails(err, "could not open tcp stream", "portID", req.PortID, "id", req.Identifier)
	}

	if conn, err := newconn(conn, "tcp", req.LocalAddress, req.RemoteAddress); err != nil {
		return errors.WrapIf(err, "could not create wrapped connection")
	} else {
		s.client.logger.V(3).Info("put stream into the connection channel", "portID", mp.id, "requestedPort", mp.requestedPort, "remoteAddress", conn.RemoteAddr())

		mp.connChan <- conn
	}

	return nil
}

func (s *ctrlStream) addPortResponse(msg []byte) error {
	var resp api.AddPortResponseMessage
	if err := json.Unmarshal(msg, &resp); err != nil {
		return errors.WrapIf(err, "could not unmarshal addPortResponse message")
	}

	if v, ok := s.client.managedPorts.Load(resp.ID); ok {
		if mp, ok := v.(*managedPort); ok {
			if resp.AssignedPort == 0 {
				_ = mp.Close()

				s.client.managedPorts.Delete(resp.ID)

				return errors.NewWithDetails("could not assign port", "portID", resp.ID)
			}

			mp.remoteAddress = resp.Address
			mp.initialized = true

			s.client.logger.V(2).Info("port added", "portID", resp.ID, "remoteAddress", resp.Address)

			return nil
		}
	}

	return errors.WithStackIf(api.ErrInvalidPort)
}