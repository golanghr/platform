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
	"testing"

	"github.com/golanghr/platform/logging"
	"github.com/golanghr/platform/options"
	"github.com/golanghr/platform/service"
	. "github.com/smartystreets/goconvey/convey"
)

var (
	httpServiceName        = "HTTP Test"
	httpServiceDescription = "Testing http server"
	httpServiceVersion     = 0.1
	httpTLS                = false
	httpTLSCert            = "test_data/server.crt"
	httpTLSKey             = "test_data/server.key"
	httpListenForever      = true
)

func getHTTPOptions() options.Options {
	opts, _ := options.New("memo", map[string]interface{}{
		"service-name":        httpServiceName,
		"service-description": httpServiceDescription,
		"service-version":     httpServiceVersion,
		"http-addr":           ":7321",
		"http-listen-forever": httpListenForever,
		"grpc-tls":            httpTLS,
		"grpc-tls-cert":       httpTLSCert,
		"grpc-tls-key":        httpTLSKey,
	})

	return opts
}

func getHTTPService() service.Servicer {
	service, _ := service.New(getHTTPOptions())
	return service
}

func TestHttpOptions(t *testing.T) {
	opts, err := options.New("memo", map[string]interface{}{
		"service-name":        httpServiceName,
		"service-description": httpServiceDescription,
		"service-version":     httpServiceVersion,
		"http-addr":           ":7321",
		"http-listen-forever": httpListenForever,
	})

	Convey("By initializing options we are getting proper options memo interface without any errors", t, func() {
		So(opts, ShouldHaveSameTypeAs, &options.Memo{})
		So(err, ShouldBeNil)
	})

	Convey("By NOT providing http-addr we are getting error", t, func() {
		opts, _ := options.New("memo", map[string]interface{}{})
		grpcserv, err := NewHTTPServer(getHTTPService(), opts, logging.New(getHTTPOptions()))
		So(grpcserv, ShouldBeNil)
		So(err.Error(), ShouldContainSubstring, "You must provide `http-addr`")
	})

	Convey("By providing valid http-tls but no http-tls-cert we are getting error", t, func() {
		opts, _ := options.New("memo", map[string]interface{}{
			"http-addr": ":7321",
			"http-tls":  true,
		})
		grpcserv, err := NewHTTPServer(getHTTPService(), opts, logging.New(getHTTPOptions()))
		So(grpcserv, ShouldBeNil)
		So(err.Error(), ShouldContainSubstring, "You must provide `http-tls-cert`")
	})

	Convey("By providing valid http-tls but no http-tls-key we are getting error", t, func() {
		opts, _ := options.New("memo", map[string]interface{}{
			"http-addr":     ":7321",
			"http-tls":      true,
			"http-tls-cert": httpTLSKey,
		})
		grpcserv, err := NewHTTPServer(getHTTPService(), opts, logging.New(getHTTPOptions()))
		So(grpcserv, ShouldBeNil)
		So(err.Error(), ShouldContainSubstring, "You must provide `http-tls-key`")
	})
}

func TestHttpInterface(t *testing.T) {
	httpserv, err := NewHTTPServer(getHTTPService(), getHTTPOptions(), logging.New(getHTTPOptions()))

	Convey("By initializing HTTP server we are getting proper *server.HTTP without any errors", t, func() {
		So(httpserv, ShouldHaveSameTypeAs, &HTTP{})
		So(err, ShouldBeNil)
	})

	Convey("By accessing Interface() we are getting &HTTP interface", t, func() {
		So(httpserv.Interface(), ShouldHaveSameTypeAs, &HTTP{})
	})

}
