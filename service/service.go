// Copyright 2015 The Golang.hr Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package service ...
package service

import "github.com/golanghr/platform/config"

// Instance - Service instance wrapper
type Instance struct {
	Config config.Manager
}

// Name - Will return name of the service
func (i *Instance) Name() string {
	var value *config.Value
	var err error

	if value, err = i.Config.Get("service-name"); err != nil {
		return ""
	}

	return value.Value()
}

// Description - Will return description of the service
func (i *Instance) Description() string {
	var value *config.Value
	var err error

	if value, err = i.Config.Get("service-description"); err != nil {
		return ""
	}

	return value.Value()
}

// Version - Will return version of the service
func (i *Instance) Version() string {
	var value *config.Value
	var err error

	if value, err = i.Config.Get("service-version"); err != nil {
		return ""
	}

	return value.Value()
}

// New -
func New(cnf config.Manager) (s Service, err error) {
	s = Service(Instance{
		Config: cnf,
	})

	if _, err = cnf.Get("service-name"); err != nil {
		return
	}

	if _, err = cnf.Get("service-description"); err != nil {
		return
	}

	if _, err = cnf.Get("service-version"); err != nil {
		return
	}

	return
}
