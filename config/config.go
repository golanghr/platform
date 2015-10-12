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
	res, err := mi.Kapi.Set(context.Background(), key, value, nil)

	if err != nil {
		if err == context.Canceled {
			return nil, fmt.Errorf("Could not set configuration (key: %s) due to context being canceled", key)
		} else if err == context.DeadlineExceeded {
			return nil, fmt.Errorf("Could not set configuration (key: %s) due to context deadline being exceeded", key)
		} else if cerr, ok := err.(*etcdc.ClusterError); ok {
			return nil, fmt.Errorf("Could not set configuration (key: %s) due to (cluster_err: %s)", key, cerr.Detail())
		}

		// bad cluster endpoints, which are not etcd servers
		return nil, fmt.Errorf("Could not set configuration (key: %s) due to (err: %s)", key, err)
	}

	return &Value{res}, err
}

// SetTTL -
func (mi *ManagerInstance) SetTTL(key, value string, ttl time.Duration) (*Value, error) {
	res, err := mi.Kapi.Set(context.Background(), key, value, &etcdc.SetOptions{
		TTL: ttl,
	})

	if err != nil {
		if err == context.Canceled {
			return nil, fmt.Errorf("Could not set configuration (key: %s) due to context being canceled", key)
		} else if err == context.DeadlineExceeded {
			return nil, fmt.Errorf("Could not set configuration (key: %s) due to context deadline being exceeded", key)
		} else if cerr, ok := err.(*etcdc.ClusterError); ok {
			return nil, fmt.Errorf("Could not set configuration (key: %s) due to (cluster_err: %s)", key, cerr.Detail())
		}

		// bad cluster endpoints, which are not etcd servers
		return nil, fmt.Errorf("Could not set configuration (key: %s) due to (err: %s)", key, err)
	}

	return &Value{res}, err
}

// Get -
func (mi *ManagerInstance) Get(key string) (*Value, error) {
	res, err := mi.Kapi.Get(context.Background(), key, &etcdc.GetOptions{Quorum: true})

	if err != nil {
		if err == context.Canceled {
			return nil, fmt.Errorf("Could not get configuration (key: %s) due to context being canceled", key)
		} else if err == context.DeadlineExceeded {
			return nil, fmt.Errorf("Could not get configuration (key: %s) due to context deadline being exceeded", key)
		} else if cerr, ok := err.(*etcdc.ClusterError); ok {
			return nil, fmt.Errorf("Could not get configuration (key: %s) due to (cluster_err: %s)", key, cerr.Detail())
		}

		// bad cluster endpoints, which are not etcd servers
		return nil, fmt.Errorf("Could not get configuration (key: %s) due to (err: %s)", key, err)
	}

	return &Value{res}, err
}
