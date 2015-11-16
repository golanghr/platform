/*
Copyright (c) 2015 Golang Croatia
All rights reserved.

The MIT License (MIT)

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

// Package server ...
package server

import (
	"errors"
	"fmt"
	"net"

	"github.com/golanghr/platform/logging"
	"github.com/golanghr/platform/options"
	"github.com/golanghr/platform/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// Grpc -
type Grpc struct {

	// IF grpc fails, it will restart it immediately...
	ListenForever bool

	// ConnectivityState indicates the state of a grpc connection.
	*ConnectivityState

	options.Options
	net.Listener
	logging.Logging
	service.Servicer

	*grpc.Server
}

// Interface - Will return current server as interface. This is useful in case you
// need to access portion of the struct that is not acccessibble via interface
func (g *Grpc) Interface() interface{} {
	return g
}

// Start - Proxy on top of grpc.Server with functionallity to auto restart if
// Grpc.ListenForever is set to be true.
func (g *Grpc) Start() error {
	g.Infof("Starting `%s` platform service ...", g.Name())

	g.ConnectivityState.SetStateConnecting()

	errors := make(chan error)

	go func(errors chan error) {
		for {
			g.ConnectivityState.SetStateReady()

			if err := g.Server.Serve(g.Listener); err != nil {
				g.Errorf(
					"[GRPC] Service `%s` listener failed due to (err: %s). Restarting service ? (%v)...",
					g.Name(), err, g.ListenForever,
				)
				errors <- err
				g.ConnectivityState.SetStateFailed()

				if !g.ListenForever {
					return
				}
			}
		}
	}(errors)

	select {
	case err := <-errors:
		return err
	}
}

// Stop - Proxy for grpc.Stop()
// @TODO - Figure out way how to handle graceful shutdown...
func (g *Grpc) Stop() error {
	g.Infof("Stopping `%s` platform service ...", g.Name())
	g.ConnectivityState.SetStateShutdown()
	g.Server.Stop()
	g.ConnectivityState.SetStateDown()
	return nil
}

// Restart - Will initiate full service restart ...
func (g *Grpc) Restart() error {
	g.Infof("Restarting `%s` platform service ...", g.Name())

	// If status is connected
	// Initiate stop... (use waitgroup)

	// Initiate start...
	return g.Start()
}

// State - indicates the state of a grpc connection.
func (g *Grpc) State() *ConnectivityState {
	return g.ConnectivityState
}

// NewGrpcServer -
func NewGrpcServer(serv service.Servicer, opts options.Options, logger logging.Logging) (Serverer, error) {

	addr, addrOk := opts.Get("grpc-addr")

	if !addrOk {
		return nil, errors.New("You must provide `grpc-addr` in order to create gRPC server...")
	}

	listener, err := net.Listen("tcp", addr.String())

	if err != nil {
		return nil, fmt.Errorf("Failed to listen: %v", err)
	}

	var grpcOpts []grpc.ServerOption

	if useGrpc, ok := opts.Get("grpc-tls"); ok && useGrpc.Bool() {
		certFile, certOk := opts.Get("grpc-tls-cert")
		certKey, keyOk := opts.Get("grpc-tls-key")

		if !certOk {
			return nil, errors.New("You must provide `grpc-tls-cert` in order to create gRPC server...")
		}

		if !keyOk {
			return nil, errors.New("You must provide `grpc-tls-key` in order to create gRPC server...")
		}

		creds, err := credentials.NewServerTLSFromFile(certFile.String(), certKey.String())

		if err != nil {
			return nil, fmt.Errorf("Failed to generate gRPC credentials: %v", err)
		}

		grpcOpts = []grpc.ServerOption{grpc.Creds(creds)}
	}

	if maxStreams, msOk := opts.Get("grpc-max-concurrent-streams"); msOk {
		grpcOpts = []grpc.ServerOption{grpc.MaxConcurrentStreams(maxStreams.UInt32())}
	}

	grpcServer := grpc.NewServer(grpcOpts...)

	s := &Grpc{
		Options:           opts,
		Server:            grpcServer,
		Listener:          listener,
		Logging:           logger,
		Servicer:          serv,
		ConnectivityState: &ConnectivityState{},
	}

	if listenForever, lfOk := opts.Get("grpc-listen-forever"); lfOk {
		s.ListenForever = listenForever.Bool()
	}

	return Serverer(s), nil
}
