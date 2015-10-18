// Copyright 2015 The Golang.hr Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package config ...
package config

import (
	"time"

	etcdc "github.com/coreos/etcd/client"
)

// Config - Basic CRUD operations against Etcd configuration path
type Config interface {
	Get(key string) (*Value, error)
	GetOrSet(key, defaults string) (*Value, error)

	GetBool(key string) (bool, error)
	GetString(key string) (string, error)
	GetDuration(key string) (time.Duration, error)

	Set(key, value string) (*Value, error)
	SetTTL(key, value string, ttl time.Duration) (*Value, error)

	EnsureSet(key, defaults string) error
	EnsureSetMany(map[string]string) error

	Exists(key string) error

	Delete(key string) (*Value, error)
}

// Manager -
type Manager interface {
	Config

	Etcd() etcdc.Client
	ShouldAutoSyncNodes() bool
	SyncNodes(interval time.Duration) error
	Init() error
}
