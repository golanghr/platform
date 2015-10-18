// Copyright 2015 The Golang.hr Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package service ...
package service

import (
	"github.com/golanghr/platform/config"
	"github.com/golanghr/platform/logging"
)

// Service -
type Service interface {
	Name() string
	Description() string
	Version() string

	GetConfig() config.Manager
	GetLogger() logging.Logging
	GetQuitChan() chan bool
}
