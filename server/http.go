// Copyright 2015 The Golang.hr Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package server ...
package server

import (
	"github.com/golanghr/platform/config"
	"github.com/golanghr/platform/service"
)

// HTTPServer -
type HTTPServer struct {
	service.Service
	ConfigManager config.Manager

	UseTLS bool

	Done chan bool
}

// Start -
func (hs *HTTPServer) Start(event chan string) (err error) {

	// Notify that we are starting service now ...
	event <- STARTING

	return nil
}

// RunOnAddr -
func (hs *HTTPServer) RunOnAddr() error {
	return nil

	/**
	errs := make(chan error)

	addr := s.Config.GetString("addr")
	sslAddr := s.Config.GetString("ssl_addr")
	sslCert := s.Config.GetMapStrings("ssl_cert")

	// Starting HTTP server
	go func() {
		logger.Warning("Staring acknowledge.io HTTP service on %s (%s)\n", addr, martini.Env)

		if err := http.ListenAndServe(addr, s); err != nil {
			logger.Error("Could not start listening HTTP service due to (error: %s)", err)
			errs <- err
		}
	}()

	// Starting HTTPS server
	go func() {
		logger.Warning("Staring acknowledge.io HTTPS service on %s (%s)\n", sslAddr, martini.Env)

		if err := http.ListenAndServeTLS(sslAddr, sslCert["cert"], sslCert["key"], s); err != nil {
			logger.Error("Could not start listening HTTPS service due to (error: %s)", err)
			errs <- err
		}
	}()

	select {
	case err := <-errs:
		return err
	}

	**/
}

// Stop -
func (hs *HTTPServer) Stop() (err error) {
	return
}

// NewHTTPServer -
func NewHTTPServer(serv service.Service) (Server, error) {
	s := HTTPServer{Service: serv, ConfigManager: serv.GetConfig()}

	var err error

	if s.UseTLS, err = s.ConfigManager.GetBool(ServerHttpTls); err != nil {
		return nil, err
	}

	return Server(&s), nil
}
