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

// GetConfig - Will return back configuration manager
func (i *Instance) GetConfig() config.Manager {
	return i.Config
}

// Name - Will return name of the service
func (i *Instance) Name() string {
	var value *config.Value
	var err error

	if value, err = i.Config.Get(ServiceNameTag); err != nil {
		return ""
	}

	return value.Value()
}

// Description - Will return description of the service
func (i *Instance) Description() string {
	var value *config.Value
	var err error

	if value, err = i.Config.Get(ServiceDescriptionTag); err != nil {
		return ""
	}

	return value.Value()
}

// Version - Will return version of the service
func (i *Instance) Version() string {
	var value *config.Value
	var err error

	if value, err = i.Config.Get(ServiceVersionTag); err != nil {
		return ""
	}

	return value.Value()
}

// New -
func New(cnf config.Manager) (s Service, err error) {
	s = Service(&Instance{
		Config: cnf,
	})

	if _, err = cnf.Get(ServiceNameTag); err != nil {
		return
	}

	if _, err = cnf.Get(ServiceDescriptionTag); err != nil {
		return
	}

	if _, err = cnf.Get(ServiceVersionTag); err != nil {
		return
	}

	return
}
