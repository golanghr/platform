// Copyright 2015 The Golang.hr Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package config ...
package config

import "github.com/golanghr/platform/utils"

// GetBool - Helper function that will try to fetch key from etcd and parse it as bool
// In case of any issues it will return error
func (mi *ManagerInstance) GetBool(key string) (bool, error) {
	var val *Value
	var err error

	if val, err = mi.Get(key); err != nil {
		return false, err
	}

	if utils.StringInSlice(val.Value(), []string{"yes", "1", "true"}) {
		return true, nil
	}

	return false, nil
}
