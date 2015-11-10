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

// Package manager ...
package manager

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

// GrpcServer -
type GrpcServer struct {
	options.Options
	net.Listener
	logging.Logging
	service.Servicer

	*grpc.Server
}

// Interface - Will return current server as interface. This is useful in case you
// need to access portion of the struct that is not acccessibble via interface
func (g *GrpcServer) Interface() interface{} {
	return g
}

// Start - Proxy for grpc.Serve(), http.Serve()
// @TODO - Add auto-restart. Now it will keep running until error
func (g *GrpcServer) Start() error {
	g.Infof("Starting `%s` platform service ...", g.Name())
	return g.Server.Serve(g.Listener)
}

// Stop - Proxy for grpc.Stop(), http.Stop()
// @TODO - Make sure error can be returned. If not, kill error from stop entirely...
func (g *GrpcServer) Stop() error {
	g.Infof("Stopping `%s` platform service ...", g.Name())
	g.Server.Stop()
	return nil
}

// Restart - Will initiate full service restart ...
func (g *GrpcServer) Restart() error {
	g.Infof("Restarting `%s` platform service ...", g.Name())

	// If status is connected
	// Initiate stop... (use waitgroup)

	// Initiate start...
	return g.Start()
}

// Status -
func (g *GrpcServer) Status() (int64, error) {
	return 0, nil
}

// NewGrpcServer -
func NewGrpcServer(serv service.Servicer, opts options.Options, logger logging.Logging) (Managerer, error) {

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

	return Managerer(&GrpcServer{
		Options:  opts,
		Server:   grpcServer,
		Listener: listener,
		Logging:  logger,
		Servicer: serv,
	}), nil
}
