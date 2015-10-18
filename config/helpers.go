// Copyright 2015 The Golang.hr Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package config ...
package config

import (
	"strconv"
	"time"
)

// GetBool - Helper function that will try to fetch key from etcd and parse it as bool
// In case of any issues it will return error
func (mi *ManagerInstance) GetBool(key string) (bool, error) {
	var val *Value
	var err error
	var result bool

	if val, err = mi.Get(key); err != nil {
		return false, err
	}

	if result, err = strconv.ParseBool(val.Value()); err != nil {
		return false, err
	}

	return result, nil
}

// GetString - Helper function that will try to fetch key from etcd and parse it as string
// In case of any issues it will return error
func (mi *ManagerInstance) GetString(key string) (string, error) {
	var val *Value
	var err error

	if val, err = mi.Get(key); err != nil {
		return "", err
	}

	return val.Value(), nil
}

// GetDuration - Helper function that will try to fetch key from etcd and parse it as time.Duration
// In case of any issues it will return error
func (mi *ManagerInstance) GetDuration(key string) (time.Duration, error) {
	var val *Value
	var err error
	var i int

	if val, err = mi.Get(key); err != nil {
		return 0, err
	}

	if i, err = strconv.Atoi(val.Value()); err != nil {
		return 0, err
	}

	return time.Duration(i), nil
}
