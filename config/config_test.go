// Copyright 2015 The Golang.hr Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package config ...
package config

import (
	"fmt"
	"net"
	"net/http"
	"testing"
	"time"

	etcdc "github.com/coreos/etcd/client"
	. "github.com/smartystreets/goconvey/convey"
)

var (
	testEnv        = "test_local"
	testEtcdFolder = "golanghr-test"
)

// GetManager - Helper
func getTestManager() (Manager, error) {
	return New(map[string]interface{}{
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
}

// TestManagerConfigDefaults -
func TestManagerConfigDefaults(t *testing.T) {
	Convey("Test If Env Is Required", t, func() {
		manager, err := New(map[string]interface{}{})
		So(err.Error(), ShouldEqual, fmt.Errorf(ErrorInvalidEnv, map[string]interface{}{}).Error())
		So(manager, ShouldBeNil)
	})

	Convey("Test If Folder Is Required", t, func() {

	})

	Convey("Test If Etcd Is Required", t, func() {

	})
}

// TestNewEventCreation -
func TestNewManagerCreation(t *testing.T) {
	Convey("Test If Manager/Etcd", t, func() {
		manager, err := getTestManager()

		So(err, ShouldBeNil)
		So(manager.Etcd(), ShouldNotBeNil)
		So(manager, ShouldHaveSameTypeAs, &ManagerInstance{})
	})

}

// TestCustomTransport -
func TestCustomTransport(t *testing.T) {
	Convey("Test If Custom Cancellable Transport", t, func() {
		var CustomHTTPTransport etcdc.CancelableTransport = &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			Dial: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).Dial,
			TLSHandshakeTimeout: 10 * time.Second,
		}

		So(CustomHTTPTransport, ShouldHaveSameTypeAs, &http.Transport{})

		manager, err := New(map[string]interface{}{
			"env":                testEnv,
			"folder":             testEtcdFolder,
			"auto_sync":          true,
			"auto_sync_interval": 10 * time.Second,
			"etcd": map[string]interface{}{
				"version":                    "v2",
				"endpoints":                  []string{"http://127.0.0.1:2379"},
				"transport":                  CustomHTTPTransport,
				"username":                   "",
				"password":                   "",
				"header_timeout_per_request": time.Second,
			},
		})

		So(err, ShouldBeNil)

		So(manager.Etcd(), ShouldNotBeNil)
		So(manager, ShouldHaveSameTypeAs, &ManagerInstance{})
	})
}

// TestAggregation -
func TestAggregation(t *testing.T) {
	Convey("Test If Custom Cancellable Transport", t, func() {
		manager, err := getTestManager()

		So(err, ShouldBeNil)
		So(manager.Etcd(), ShouldNotBeNil)
		So(manager, ShouldHaveSameTypeAs, &ManagerInstance{})

		value, err := manager.SetTTL("platform", "Test Golang.hr Platform", 10*time.Second)
		So(value, ShouldHaveSameTypeAs, &Value{})
		So(err, ShouldBeNil)

		gvalue, gerr := manager.Get("platform")
		So(gvalue, ShouldHaveSameTypeAs, &Value{})
		So(gerr, ShouldBeNil)
		So(gvalue.Value(), ShouldEqual, "Test Golang.hr Platform")

	})
}
