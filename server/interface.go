// Copyright 2015 The Golang.hr Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package server ...
package server

// Server -
type Server interface {
	Start(chan string) error
	Stop() error

	RunOnAddr() error
}
