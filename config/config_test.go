// Copyright 2015 The Golang.hr Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package config ...
package config

import (
	"testing"
	"time"

	etcdc "github.com/coreos/etcd/client"
	. "github.com/smartystreets/goconvey/convey"
)

var (
	testEnv        = "test_local"
	testEtcdFolder = "golanghr-test"
)

// TestNewEventCreation -
func TestNewManagerCreation(t *testing.T) {
	Convey("Test If Manager Instance", t, func() {
		manager, err := NewManager(map[string]interface{}{
			"env": testEnv,
			"etcd": map[string]interface{}{
				"folder":                     testEtcdFolder,
				"endpoints":                  []string{"127.0.0.1:2379"},
				"transport":                  etcdc.DefaultTransport,
				"username":                   "",
				"password":                   "",
				"header_timeout_per_request": time.Second,
			},
		})

		So(err, ShouldHaveSameTypeAs, nil)
		So(manager, ShouldHaveSameTypeAs, &ManagerInstance{})
	})
}
