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
	"errors"

	"github.com/gengo/grpc-gateway/runtime"
	"github.com/golanghr/platform/logging"
	"github.com/golanghr/platform/options"
	"github.com/golanghr/platform/services"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// HttpGrpcHandler -
type HttpGrpcHandler struct {
	ctx context.Context
	*runtime.ServeMux

	// Servicer - Is a Servicer interface
	services.Servicer

	*logging.Entry

	CertFile   *options.Option
	CertDomain *options.Option
	GrpcAddr   *options.Option
}

// RegisterHandlerTLS - Will register new TLS
func (hgh *HttpGrpcHandler) RegisterHandlerTLS(handler func(context.Context, *runtime.ServeMux, *grpc.ClientConn) error) error {
	var opts []grpc.DialOption
	var creds credentials.TransportAuthenticator

	creds, err := credentials.NewClientTLSFromFile(hgh.CertFile.String(), hgh.CertDomain.String())

	if err != nil {
		return err
	}

	opts = append(opts, grpc.WithTransportCredentials(creds))

	conn, err := grpc.Dial(hgh.GrpcAddr.String(), opts...)

	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				hgh.Errorf("Failed to close conn to %s: %v", hgh.GrpcAddr.String(), cerr)
			}
			return
		}

		go func() {
			<-hgh.ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				hgh.Errorf("Failed to close conn to %s: %v", hgh.GrpcAddr.String(), cerr)
			}
		}()

	}()

	return handler(hgh.ctx, hgh.ServeMux, conn)
}

// NewHttpGrpcHandler - Will return back new HTTP GRPC handler
func NewHttpGrpcHandler(serv services.Servicer, logger *logging.Entry, handler func(context.Context, *runtime.ServeMux, *grpc.ClientConn) error) (*HttpGrpcHandler, error) {
	grpcAddr, ok := serv.GetOptions().Get("grpc-addr")

	if !ok {
		return nil, errors.New("In order to register new grpc http handler `grpc-addr` must be set.")
	}

	certFile, ok := serv.GetOptions().Get("grpc-tls-cert")

	if !ok {
		return nil, errors.New("In order to register new grpc http handler `grpc-tls-cert` must be set.")
	}

	certDomain, ok := serv.GetOptions().Get("grpc-tls-domain")
	if !ok {
		return nil, errors.New("In order to register new grpc http handler `grpc-tls-domain` must be set.")
	}

	httpgrpc := &HttpGrpcHandler{
		ctx:        context.Background(),
		ServeMux:   runtime.NewServeMux(),
		Servicer:   serv,
		Entry:      logger,
		GrpcAddr:   grpcAddr,
		CertFile:   certFile,
		CertDomain: certDomain,
	}

	// Will register provider handler (protocol buffer client) and start grpc.Dial
	if err := httpgrpc.RegisterHandlerTLS(handler); err != nil {
		return nil, err
	}

	return httpgrpc, nil
}
