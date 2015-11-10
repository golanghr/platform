// Copyright 2015 The Golang.hr Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package service ...
package service

import "github.com/golanghr/platform/options"

// Servicer -
type Servicer interface {
	Name() string
	Description() string
	Version() float64

	GetOptions() options.Options
	GetQuitChan() chan bool
}
