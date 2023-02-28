// Copyright (c) 2023 Cisco and/or its affiliates. All rights reserved.
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
	"net"

	"github.com/cisco-open/nasp/pkg/network/tunnel/api"
)

var _ net.Listener = &managedPort{}

type managedPort struct {
	id            string
	requestedPort int
	remoteAddress string

	connChan chan net.Conn

	initialized bool
}

func NewManagedPort(id string, requestedPort int) *managedPort {
	return &managedPort{
		id:            id,
		requestedPort: requestedPort,

		connChan: make(chan net.Conn, 1),
	}
}

func (p *managedPort) Accept() (net.Conn, error) {
	conn, open := <-p.connChan

	if !open {
		return nil, api.ErrListenerStopped
	}

	if conn == nil {
		return nil, api.ErrInvalidConnection
	}

	return conn, nil
}

func (p *managedPort) Close() error {
	close(p.connChan)

	return nil
}

func (p *managedPort) Addr() net.Addr {
	if p.remoteAddress != "" {
		if addr, err := net.ResolveTCPAddr("tcp", p.remoteAddress); err == nil {
			return addr
		}
	}

	return &net.TCPAddr{
		Port: p.requestedPort,
	}
}