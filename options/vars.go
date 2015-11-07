// Package options ...
// Copyright 2015 The Golang.hr Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package options

var (

	// AvailableAdapters - List of available adapters that are currently supported
	AvailableAdapters = map[string]string{
		"memo": "Memo",
		"etcd": "Etcd",
	}
)
