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
	"testing"
	"time"

	"github.com/golanghr/platform/logging"
	"github.com/golanghr/platform/options"
	"github.com/golanghr/platform/services"
	. "github.com/smartystreets/goconvey/convey"
)

var (
	serviceName        = "Grpc Test"
	serviceDescription = "Testing grpc server"
	serviceVersion     = 0.1
	grpcAddr           = ""
	grpcTLS            = false
	grpcTLSCert        = "test_data/server.crt"
	grpcTLSKey         = "test_data/server.key"
	grpcMaxStreams     = uint32(20)
	grpcListenForever  = true

	testOptions options.Options
	testService services.Servicer
	testLogging logging.Logging
)

func getOptions() options.Options {
	opts, _ := options.New("memo", map[string]interface{}{
		"service-name":                serviceName,
		"service-description":         serviceDescription,
		"service-version":             serviceVersion,
		"grpc-listen-forever":         grpcListenForever,
		"grpc-addr":                   grpcAddr,
		"grpc-tls":                    grpcTLS,
		"grpc-tls-cert":               grpcTLSCert,
		"grpc-tls-key":                grpcTLSKey,
		"grpc-max-concurrent-streams": grpcMaxStreams,
	})

	return opts
}

func getService() service.Servicer {
	service, _ := service.New(getOptions())
	return service
}

func TestGrpcInterface(t *testing.T) {
	grpcserv, err := NewGrpcServer(getService(), getOptions(), logging.New(getOptions()))

	Convey("By initializing GRPC server we are getting proper *server.Grpc without any errors", t, func() {
		So(grpcserv, ShouldHaveSameTypeAs, &Grpc{})
		So(err, ShouldBeNil)
	})

	Convey("By accessing Interface() we are getting &Grpc interface", t, func() {
		So(grpcserv.Interface(), ShouldHaveSameTypeAs, &Grpc{})
	})

}

func TestGrpcOptions(t *testing.T) {
	opts, err := options.New("memo", map[string]interface{}{
		"service-name":                serviceName,
		"service-description":         serviceDescription,
		"service-version":             serviceVersion,
		"grpc-listen-forever":         grpcListenForever,
		"grpc-addr":                   grpcAddr,
		"grpc-tls":                    grpcTLS,
		"grpc-tls-cert":               grpcTLSCert,
		"grpc-tls-key":                grpcTLSKey,
		"grpc-max-concurrent-streams": grpcMaxStreams,
	})

	Convey("By initializing options we are getting proper options memo interface without any errors", t, func() {
		So(opts, ShouldHaveSameTypeAs, &options.Memo{})
		So(err, ShouldBeNil)
	})

	Convey("By NOT providing grpc-addr we are getting error", t, func() {
		opts, _ := options.New("memo", map[string]interface{}{})
		grpcserv, err := NewGrpcServer(getService(), opts, logging.New(getOptions()))
		So(grpcserv, ShouldBeNil)
		So(err.Error(), ShouldContainSubstring, "You must provide `grpc-addr`")
	})

	Convey("By providing invalid grpc-addr we are getting failed to listen error", t, func() {
		opts, _ := options.New("memo", map[string]interface{}{
			"grpc-addr": "I Am Invalid",
		})
		grpcserv, err := NewGrpcServer(getService(), opts, logging.New(getOptions()))
		So(grpcserv, ShouldBeNil)
		So(err.Error(), ShouldContainSubstring, "Failed to listen: listen tcp")
	})

	Convey("By providing valid grpc-tls but no grpc-tls-cert we are getting error", t, func() {
		opts, _ := options.New("memo", map[string]interface{}{
			"grpc-addr": grpcAddr,
			"grpc-tls":  true,
		})
		grpcserv, err := NewGrpcServer(getService(), opts, logging.New(getOptions()))
		So(grpcserv, ShouldBeNil)
		So(err.Error(), ShouldContainSubstring, "You must provide `grpc-tls-cert`")
	})

	Convey("By providing valid grpc-tls but no grpc-tls-key we are getting error", t, func() {
		opts, _ := options.New("memo", map[string]interface{}{
			"grpc-addr":     grpcAddr,
			"grpc-tls":      true,
			"grpc-tls-cert": grpcTLSCert,
		})
		grpcserv, err := NewGrpcServer(getService(), opts, logging.New(getOptions()))
		So(grpcserv, ShouldBeNil)
		So(err.Error(), ShouldContainSubstring, "You must provide `grpc-tls-key`")
	})

	Convey("By providing invalid grpc-tls-key we are getting gRPC credentials error", t, func() {
		opts, _ := options.New("memo", map[string]interface{}{
			"grpc-addr":     grpcAddr,
			"grpc-tls":      true,
			"grpc-tls-cert": grpcTLSCert,
			"grpc-tls-key":  "/tmp/i-am-invalid",
		})
		grpcserv, err := NewGrpcServer(getService(), opts, logging.New(getOptions()))
		So(grpcserv, ShouldBeNil)
		So(err.Error(), ShouldContainSubstring, "Failed to generate gRPC credentials")
	})

	Convey("By providing valid grpc tls information we get no errors", t, func() {
		opts, _ := options.New("memo", map[string]interface{}{
			"grpc-addr":     grpcAddr,
			"grpc-tls":      true,
			"grpc-tls-cert": grpcTLSCert,
			"grpc-tls-key":  grpcTLSKey,
		})
		grpcserv, err := NewGrpcServer(getService(), opts, logging.New(getOptions()))
		So(grpcserv, ShouldHaveSameTypeAs, &Grpc{})
		So(err, ShouldBeNil)
	})
}

func TestGrpcConnectivityState(t *testing.T) {
	grpcserv, err := NewGrpcServer(getService(), getOptions(), logging.New(getOptions()))
	grpcstate := grpcserv.State()

	Convey("By initializing GRPC server we are getting proper *server.Grpc without any errors", t, func() {
		So(grpcserv, ShouldHaveSameTypeAs, &Grpc{})
		So(err, ShouldBeNil)
	})

	Convey("By manipulating runtime connection state is changing.", t, func() {
		So(grpcstate.GetCurrentState(), ShouldEqual, grpcstate.GetStateByName("down"))

		go grpcserv.Start()

		time.Sleep(100 * time.Microsecond)
		So(grpcstate.GetCurrentState(), ShouldEqual, grpcstate.GetStateByName("ready"))

		err := grpcserv.Stop()
		So(err, ShouldBeNil)
		So(grpcstate.GetCurrentState(), ShouldEqual, grpcstate.GetStateByName("down"))
	})
}
