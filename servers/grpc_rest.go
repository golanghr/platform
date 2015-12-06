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
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/golanghr/platform/logging"
	"github.com/golanghr/platform/options"
	"github.com/golanghr/platform/services"
)

// GrpcRest -
type GrpcRest struct {
	options.Options
	*logging.Entry
	services.Servicer

	// ConnectivityState indicates the state of a http connection.
	*ConnectivityState
	*http.Server

	// ListenForever if set to true will restart it immediately if listening fails
	ListenForever bool

	// Limit the number of outstanding requests
	ListenLimit int

	// TLS Certificate
	TLSCertFile string

	// TLS Key
	TLSKeyFile string

	// Interrupt signals the listener to stop serving connections,
	// and the server to shut down.
	Interrupt chan os.Signal

	// Mutex is used to protect against concurrent calls to Stop
	sync.Mutex

	// wg is used ...
	wg sync.WaitGroup

	connections map[string]net.Conn
}

// Interface -
func (gr *GrpcRest) Interface() interface{} {
	return gr
}

// Start - Will attempt to start HTTP server depending on actual configuration
// In case that ListenForever is enabled we are going to try to restart server until server starts
// In case of interruption we are going to stop http server
func (gr *GrpcRest) Start() error {
	errors := make(chan error)

	// Debug message to help us understand from logging level if we're listening forever
	// or until the first error
	if gr.ListenForever {
		gr.Debug("GRPC-REST server will listen forever ( will auto restart on error ) ...")
	}

	go func() {
		for {
			gr.SetStateConnecting()
			gr.Infof("GRPC-REST (Re)Starting platform server (addr: %s) ...", gr.Addr)

			// In case that server fails for whatever reason, error bellow will ensure
			// this changes as soon as we can. There's really no reason for us to
			// over complicate this part of the code.
			go func() {
				time.Sleep(2 * time.Second)
				if gr.GetCurrentState() != gr.GetStateByName("failed") {
					gr.SetStateReady()
				}
			}()

			err := gr.ListenAndServe()

			if err != nil {
				gr.SetStateFailed()

				if gr.ListenForever {
					gr.Errorf("GRPC-REST server listener crashed with: %s. Restarting server now ....", err)
					time.Sleep(1 * time.Second)
					continue
				}

				errors <- err
				return
			}
		}
	}()

	select {
	case err := <-errors:
		return err
	case <-gr.Interrupt:
		gr.Warn("GRPC-REST server received interrupt signal. Stopping server now...")
		return nil
	}
}

// Stop -
func (gr *GrpcRest) Stop() error {
	gr.Info("Stopping platform GRPC-REST server...")

	if gr.GetCurrentState() != gr.GetStateByName("ready") {
		return fmt.Errorf("Could not stop GRPC-REST server as it is NOT running...")
	}

	gr.SetStateShutdown()

	gr.Mutex.Lock()
	gr.Server.SetKeepAlivesEnabled(false)
	gr.SetStateDown()
	gr.Mutex.Unlock()

	return nil
}

// Restart -
func (gr *GrpcRest) Restart() error {
	return nil
}

// State - State of the GRPC REST API connection
// @TODO
func (gr *GrpcRest) State() *ConnectivityState {
	return gr.ConnectivityState
}

// SetHandler -
func (gr *GrpcRest) SetHandler(uri string, handler http.Handler) {
	gr.Server.Handler = handler
}

// NewGrpcRestServer -
func NewGrpcRestServer(serv services.Servicer, opts options.Options, logger *logging.Entry) (Serverer, error) {

	addr, addrOk := opts.Get("grpc-rest-addr")

	if !addrOk {
		return nil, errors.New("You must provide `grpc-rest-addr` in order to create HTTP server...")
	}

	certFile, certOk := opts.Get("grpc-rest-tls-cert")
	certKeyFile, keyOk := opts.Get("grpc-rest-tls-key")

	s := &GrpcRest{
		Server:            &http.Server{Addr: addr.String()},
		Options:           opts,
		Entry:             logger,
		Servicer:          serv,
		ConnectivityState: &ConnectivityState{},
		Interrupt:         serv.GetInterruptChan(),
	}

	s.Server.SetKeepAlivesEnabled(true)

	if useTLS, ok := opts.Get("grpc-rest-tls"); ok && useTLS.Bool() {
		if !certOk {
			return nil, errors.New("You must provide `grpc-rest-tls-cert` in order to create HTTP server...")
		}

		if !keyOk {
			return nil, errors.New("You must provide `grpc-rest-tls-key` in order to create HTTP server...")
		}

		s.TLSCertFile = certFile.String()
		s.TLSKeyFile = certKeyFile.String()
	}

	if listenForever, lfOk := opts.Get("grpc-rest-listen-forever"); lfOk {
		s.ListenForever = listenForever.Bool()
	}

	return Serverer(s), nil
}
