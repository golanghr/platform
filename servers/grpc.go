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

// Package servers ...
package servers

import (
	"errors"
	"fmt"
	"net"
	"os"
	"sync"
	"time"

	"github.com/golanghr/platform/logging"
	"github.com/golanghr/platform/options"
	"github.com/golanghr/platform/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// Grpc - Wrapper on top of googles grpc.io server
type Grpc struct {

	// IF grpc fails, it will restart it immediately...
	ListenForever bool

	// ConnectivityState indicates the state of a grpc connection.
	*ConnectivityState

	// Interrupt signals the listener to stop serving connections,
	// and the server to shut down.
	Interrupt chan os.Signal

	options.Options
	net.Listener
	*logging.Entry
	services.Servicer
	mu sync.Mutex
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
	// Debug message to help us from logging level understand if we're listening forever
	// or until the first error
	if g.ListenForever {
		g.Debug("GRPC Service will listen forever ( will auto restart on error ) ...")
	}

	for {
		g.Infof("GRPC (Re)Starting platform service (addr: %s) ...", g.Addr().String())

		g.SetStateConnecting()

		// This is really just a hack how to notify every one else that things have been started.
		// In case that there are some failures, it will be altered later on.
		go func() {
			time.Sleep(3 * time.Second)
			if g.GetCurrentState() != g.GetStateByName("failed") {
				g.SetStateReady()
			}
		}()

		if err := g.Server.Serve(g.Listener); err != nil {
			g.SetStateFailed()

			if g.ListenForever {
				g.Errorf(
					"HTTP server listener crashed due to (err: %s). Restarting server now...",
					err, g.ListenForever,
				)

				time.Sleep(1 * time.Second)
				continue
			}

			return err
		}
	}
}

// Stop - Will attempt to stop gRPC server. In case that server is not started
// we will return error
func (g *Grpc) Stop() error {
	g.Info("Stopping platform GRPC server...")

	if g.GetCurrentState() != g.GetStateByName("ready") {
		return fmt.Errorf("Could not stop GRPC server as it is NOT running...")
	}

	g.SetStateShutdown()

	g.mu.Lock()
	// Make sure that listen forever goes down...
	g.ListenForever = false
	g.Server.Stop()
	g.mu.Unlock()

	g.SetStateDown()
	return nil
}

// Restart - Will initiate full service restart ...
func (g *Grpc) Restart() error {
	g.Info("Restarting platform service ...")

	g.mu.Lock()
	if g.State().GetCurrentState() == g.GetStateByName("ready") {
		g.Stop()
	}
	// Make sure we return back listen-forever state back to one from options
	if listenForever, lfOk := g.Options.Get("grpc-listen-forever"); lfOk {
		g.ListenForever = listenForever.Bool()
	}
	g.mu.Unlock()

	// Initiate start...
	return g.Start()
}

// State - indicates the state of a grpc connection.
func (g *Grpc) State() *ConnectivityState {
	return g.ConnectivityState
}

// NewGrpcServer -
func NewGrpcServer(serv services.Servicer, opts options.Options, logger *logging.Entry) (Serverer, error) {

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
		Entry:             logger,
		Servicer:          serv,
		ConnectivityState: &ConnectivityState{},
		Interrupt:         serv.GetInterruptChan(),
	}

	if listenForever, lfOk := opts.Get("grpc-listen-forever"); lfOk {
		s.ListenForever = listenForever.Bool()
	}

	return Serverer(s), nil
}
