// Copyright 2015 The Golang.hr Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package config ...
package config

// Config -
type Config interface {
	Set(key string, value interface{}, ttl int64) error
	Get(key string) (Value, error)
}

// Manager -
type Manager interface {
}
