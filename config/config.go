// Copyright 2015 The Golang.hr Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package config ...
package config

import (
	"fmt"
	"time"

	etcdc "github.com/coreos/etcd/client"
	"golang.org/x/net/context"
)

// Value -
type Value struct {
	*etcdc.Response
}

// Value - Helper for etcd Node.Value
func (v *Value) Value() string {
	return v.Node.Value
}

// Set -
func (mi *ManagerInstance) Set(key, value string) (*Value, error) {
	return nil, nil
}

// SetTTL -
func (mi *ManagerInstance) SetTTL(key, value string, ttl time.Duration) (*Value, error) {
	res, err := mi.Kapi.Set(context.Background(), key, value, &etcdc.SetOptions{
		TTL: ttl,
	})

	if err != nil {
		if err == context.Canceled {
			// ctx is canceled by another routine
		} else if err == context.DeadlineExceeded {
			// ctx is attached with a deadline and it exceeded
		} else if cerr, ok := err.(*etcdc.ClusterError); ok {
			// process (cerr.Errors)
			fmt.Println(cerr)
		} else {
			// bad cluster endpoints, which are not etcd servers
		}
	}

	return &Value{res}, err
}

// Get -
func (mi *ManagerInstance) Get(key string) (*Value, error) {
	res, err := mi.Kapi.Get(context.Background(), key, &etcdc.GetOptions{Quorum: true})

	if err != nil {
		if err == context.Canceled {
			// ctx is canceled by another routine
		} else if err == context.DeadlineExceeded {
			// ctx is attached with a deadline and it exceeded
		} else if cerr, ok := err.(*etcdc.ClusterError); ok {
			// process (cerr.Errors)
			fmt.Println(cerr.Detail())
		} else {
			// bad cluster endpoints, which are not etcd servers

		}
	}

	return &Value{res}, err
}
