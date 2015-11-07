// Package options ...
// Copyright 2015 The Golang.hr Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package options

// Options -
type Options interface {
	Get(key string) (*Option, bool)
	GetMany(keys []string) []*Option
	Set(key string, value interface{}) error
	SetMany(opts map[string]interface{}) error
	Unset(key string) bool
	Interface() interface{}
	String() string
}

// Adapter -
type Adapter struct {
	Name string `option:"adapter_name"`
}

// Adapter - Name of current, initialized adapter
func (a *Adapter) String() string {
	return a.Name
}
