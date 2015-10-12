// Copyright 2015 The Golang.hr Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package config ...
package config

import (
	"time"

	etcdc "github.com/coreos/etcd/client"
	"golang.org/x/net/context"
)

// Value -
type Value struct {
	*etcdc.Response
}

// Set -
func (mi *ManagerInstance) Set(key, value string) (*Value, error) {
	return nil, nil
}

// SetTTL -
func (mi *ManagerInstance) SetTTL(key, value string, ttl time.Duration) (*Value, error) {
	response, err := mi.Kapi.Set(context.Background(), key, value, &etcdc.SetOptions{
		TTL: ttl,
	})

	return &Value{response}, err
}

// Get -
func (mi *ManagerInstance) Get(key string) (*Value, error) {
	return nil, nil
}
