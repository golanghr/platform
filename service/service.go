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

// New -
func New(cnf config.Manager) (s Service, err error) {
	s = Service(Instance{
		Config: cnf,
	})

	return
}
