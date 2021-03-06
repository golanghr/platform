// Copyright 2015 The Golang.hr Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package config ...
package config

import (
	"fmt"
	"strings"
	"time"

	etcdc "github.com/coreos/etcd/client"
	"golang.org/x/net/context"
)

// Value -  Etcd response proxy struct designed to help out normalize how we
// access data
type Value struct {
	*etcdc.Response
}

// Key - Helper for etcd Node.Key
func (v *Value) Key() string {
	return v.Node.Key
}

// Value - Helper for etcd Node.Value
func (v *Value) Value() string {
	return v.Node.Value
}

// Set - Will set new etcd key. If there are no keys, will return error
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

// SetTTL - Will set new expiration etcd key. If there are no keys, will return error
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

// Get - Will return etcd response about specified key. If there are no keys, will return error
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

// GetOrSet - Will attempt to retreive key from etcd and if one is not there, will
// attempt to create it
func (mi *ManagerInstance) GetOrSet(key, defaults string) (*Value, error) {
	res, err := mi.Kapi.Get(context.Background(), key, &etcdc.GetOptions{Quorum: true})

	if err != nil {
		if err == context.Canceled {
			return nil, fmt.Errorf("Could not get configuration (key: %s) due to context being canceled", key)
		} else if err == context.DeadlineExceeded {
			return nil, fmt.Errorf("Could not get configuration (key: %s) due to context deadline being exceeded", key)
		} else if cerr, ok := err.(*etcdc.ClusterError); ok {
			return nil, fmt.Errorf("Could not get configuration (key: %s) due to (cluster_err: %s)", key, cerr.Detail())
		}

		if !strings.Contains(err.Error(), "Key not found") {
			// bad cluster endpoints, which are not etcd servers
			return nil, fmt.Errorf("Could not get configuration (key: %s) due to (err: %s)", key, err)
		}

		return mi.Set(key, defaults)
	}

	return &Value{res}, err
}

// Exists - Will check if key exists or not
func (mi *ManagerInstance) Exists(key string) error {
	if _, err := mi.Get(key); err != nil {
		return err
	}

	return nil
}

// EnsureSet - Will ensure that key is set. If it's not set, it will set key with provided defaults.
func (mi *ManagerInstance) EnsureSet(key, defaults string) error {
	if _, err := mi.Kapi.Get(context.Background(), key, &etcdc.GetOptions{Quorum: true}); err != nil {
		if err == context.Canceled {
			return fmt.Errorf("Could not get configuration (key: %s) due to context being canceled", key)
		} else if err == context.DeadlineExceeded {
			return fmt.Errorf("Could not get configuration (key: %s) due to context deadline being exceeded", key)
		} else if cerr, ok := err.(*etcdc.ClusterError); ok {
			return fmt.Errorf("Could not get configuration (key: %s) due to (cluster_err: %s)", key, cerr.Detail())
		}

		if !strings.Contains(err.Error(), "Key not found") {
			// bad cluster endpoints, which are not etcd servers
			return fmt.Errorf("Could not get configuration (key: %s) due to (err: %s)", key, err)
		}

		if _, err := mi.Set(key, defaults); err != nil {
			return err
		}
	}

	return nil
}

// EnsureSetMany - Will recursively go over map of strings and ensure each of them is set.
func (mi *ManagerInstance) EnsureSetMany(cnf map[string]string) error {
	for key, value := range cnf {
		if err := mi.EnsureSet(key, value); err != nil {
			return err
		}
	}

	return nil
}

// Delete - Will delete etcd key. If there are no keys, will return error
func (mi *ManagerInstance) Delete(key string) (*Value, error) {
	res, err := mi.Kapi.Delete(context.Background(), key, &etcdc.DeleteOptions{Recursive: true})

	if err != nil {
		if err == context.Canceled {
			return nil, fmt.Errorf("Could not delete configuration (key: %s) due to context being canceled", key)
		} else if err == context.DeadlineExceeded {
			return nil, fmt.Errorf("Could not delete configuration (key: %s) due to context deadline being exceeded", key)
		} else if cerr, ok := err.(*etcdc.ClusterError); ok {
			return nil, fmt.Errorf("Could not delete configuration (key: %s) due to (cluster_err: %s)", key, cerr.Detail())
		}

		// bad cluster endpoints, which are not etcd servers
		return nil, fmt.Errorf("Could not delete configuration (key: %s) due to (err: %s)", key, err)
	}

	return &Value{res}, err
}
