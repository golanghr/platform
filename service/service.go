// Copyright 2015 The Golang.hr Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package service ...
package service

import "github.com/golanghr/platform/options"

// Service - Service instance wrapper
type Service struct {
	Options options.Options

	Quit chan bool
}

// GetQuitChan -
func (s *Service) GetQuitChan() chan bool {
	return s.Quit
}

// GetOptions - Will return back options manager
func (s *Service) GetOptions() options.Options {
	return s.Options
}

// Name - Will return name of the service
func (s *Service) Name() string {
	name, _ := s.Options.Get("service-name")
	return name.String()
}

// Description - Will return description of the service
func (s *Service) Description() string {
	name, _ := s.Options.Get("service-description")
	return name.String()
}

// Version - Will return version of the service
func (s *Service) Version() string {
	name, _ := s.Options.Get("service-version")
	return name.String()
}

// New -
func New(opts options.Options) (s Servicer, err error) {
	s = Servicer(&Service{
		Options: opts,
		Quit:    make(chan bool),
	})

	return
}
