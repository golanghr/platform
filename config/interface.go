// Copyright 2015 The Golang.hr Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package config ...
package config

import (
	"time"

	etcdc "github.com/coreos/etcd/client"
)

// Config -
type Config interface {
	Set(key, value string) (*Value, error)
	SetTTL(key, value string, ttl time.Duration) (*Value, error)

	Get(key string) (*Value, error)
}

// Manager -
type Manager interface {
	Config

	Etcd() etcdc.Client
	ShouldAutoSyncNodes() bool
	SyncNodes(interval time.Duration) error
	Init() error
}
