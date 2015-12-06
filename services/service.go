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

// Package services ...
package services

import (
	"os"

	"github.com/golanghr/platform/options"
)

// Service - Service instance wrapper
type Service struct {
	Options options.Options

	Interrupt chan os.Signal
}

// GetInterruptChan -
func (s *Service) GetInterruptChan() chan os.Signal {
	return s.Interrupt
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
func (s *Service) Version() float64 {
	name, _ := s.Options.Get("service-version")
	return name.Float()
}

// New -
func New(opts options.Options) (s Servicer, err error) {
	s = Servicer(&Service{
		Options:   opts,
		Interrupt: make(chan os.Signal, 1),
	})

	return
}
