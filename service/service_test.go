// Copyright 2015 The Golang.hr Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package service ...
package service

import (
	"testing"
	"time"

	etcdc "github.com/coreos/etcd/client"
	"github.com/golanghr/platform/config"
	. "github.com/smartystreets/goconvey/convey"
)

var (
	testEnv        = "test_local"
	testEtcdFolder = "golanghr-test"
)

func getService(t *testing.T) (Service, error) {
	conf, err := config.New(map[string]interface{}{
		"env":                testEnv,
		"folder":             testEtcdFolder,
		"auto_sync":          true,
		"auto_sync_interval": 10 * time.Second,
		"etcd": map[string]interface{}{
			"version":                    "v2",
			"endpoints":                  []string{"http://localhost:2379"},
			"transport":                  etcdc.DefaultTransport,
			"username":                   "",
			"password":                   "",
			"header_timeout_per_request": time.Second,
		},
	})

	Convey("Test If Configuration Is Available", t, func() {
		So(err, ShouldBeNil)
		So(conf.Etcd(), ShouldNotBeNil)
		So(conf, ShouldHaveSameTypeAs, &config.ManagerInstance{})
	})

	return New(conf)
}

func TestNewServiceCreation(t *testing.T) {
	Convey("Test Required Logging Type", t, func() {

		serv, err := getService(t)

		So(err, ShouldBeNil)
		So(serv, ShouldHaveSameTypeAs, Instance{})
	})
}
