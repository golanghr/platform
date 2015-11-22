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

// Package handlers ...
package handlers

import (
	"github.com/gengo/grpc-gateway/runtime"
	"github.com/golanghr/platform/logging"
	"github.com/golanghr/platform/service"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// HttpGrpcHandler -
type HttpGrpcHandler struct {
	ctx context.Context
	*runtime.ServeMux

	// Servicer - Is a Servicer interface
	service.Servicer

	*logging.Entry
}

// ServeHTTP -
func (hgh *HttpGrpcHandler) RegisterHandler(handler func(context.Context, *runtime.ServeMux, *grpc.ClientConn) error) error {
	var opts []grpc.DialOption
	var creds credentials.TransportAuthenticator

	certFile, _ := hgh.GetOptions().Get("grpc-tls-cert")
	domain, _ := hgh.GetOptions().Get("grpc-tls-domain")
	grpcAddr, _ := hgh.GetOptions().Get("grpc-addr")

	creds, err := credentials.NewClientTLSFromFile(certFile.String(), domain.String())

	if err != nil {
		return err
	}

	opts = append(opts, grpc.WithTransportCredentials(creds))

	conn, err := grpc.Dial(grpcAddr.String(), opts...)

	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				hgh.Errorf("Failed to close conn to %s: %v", grpcAddr.String(), cerr)
			}
			return
		}
		go func() {
			<-hgh.ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				hgh.Errorf("Failed to close conn to %s: %v", grpcAddr.String(), cerr)
			}
		}()
	}()

	return handler(hgh.ctx, hgh.ServeMux, conn)
}

// NewHttpGrpcHandler -
func NewHttpGrpcHandler(serv service.Servicer, logger *logging.Entry, handler func(context.Context, *runtime.ServeMux, *grpc.ClientConn) error) (*HttpGrpcHandler, error) {
	httpgrpc := &HttpGrpcHandler{
		ctx:      context.Background(),
		ServeMux: runtime.NewServeMux(),
		Servicer: serv,
		Entry:    logger,
	}

	//
	if err := httpgrpc.RegisterHandler(handler); err != nil {
		return nil, err
	}

	/**
	ctx := context.Background()
	mux := runtime.NewServeMux()


	grpcAddr, _ := opts.Get("grpc-addr")
	certFile, _ := opts.Get("grpc-tls-cert")
	certDomain, _ := opts.Get("grpc-tls-domain")

	if err := RegisterHelloHandlerFromEndpointTLS(ctx, mux, grpcAddr.String(), certFile.String(), certDomain.String()); err != nil {
		return nil, err
	}


	**/

	return httpgrpc, nil
}
