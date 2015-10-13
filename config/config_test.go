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
		cnf := map[string]interface{}{
			"env": testEnv,
		}

		manager, err := New(cnf)
		So(err.Error(), ShouldEqual, fmt.Errorf(ErrorInvalidFolder, cnf).Error())
		So(manager, ShouldBeNil)
	})

	Convey("Test If Etcd Is Required", t, func() {
		cnf := map[string]interface{}{
			"env":    testEnv,
			"folder": testEtcdFolder,
		}

		manager, err := New(cnf)
		So(err.Error(), ShouldEqual, fmt.Errorf(ErrorInvalidEtcdConfig, cnf).Error())
		So(manager, ShouldBeNil)
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

// TestCreateGetDeleteKeys -
func TestCreateGetDeleteKeys(t *testing.T) {
	manager, err := getTestManager()

	Convey("Test If SetTTL and Get works", t, func() {
		So(err, ShouldBeNil)
		So(manager.Etcd(), ShouldNotBeNil)
		So(manager, ShouldHaveSameTypeAs, &ManagerInstance{})
	})

	Convey("Test If SetTTL and Get works", t, func() {
		value, err := manager.SetTTL("platform", "Test Golang.hr Platform", 10*time.Second)
		So(value, ShouldHaveSameTypeAs, &Value{})
		So(err, ShouldBeNil)

		gvalue, gerr := manager.Get("platform")
		So(gvalue, ShouldHaveSameTypeAs, &Value{})
		So(gerr, ShouldBeNil)
		So(gvalue.Value(), ShouldEqual, "Test Golang.hr Platform")

	})

	Convey("Test If Set and Get works", t, func() {
		value, err := manager.Set("platform", "Test Golang.hr Platform")
		So(value, ShouldHaveSameTypeAs, &Value{})
		So(err, ShouldBeNil)

		gvalue, gerr := manager.Get("platform")
		So(gvalue, ShouldHaveSameTypeAs, &Value{})
		So(gerr, ShouldBeNil)
		So(gvalue.Value(), ShouldEqual, "Test Golang.hr Platform")
	})

	Convey("Test If Set, Delete and Get works", t, func() {
		value, err := manager.Set("platform-delete", "Test Golang.hr Platform")
		So(value, ShouldHaveSameTypeAs, &Value{})
		So(err, ShouldBeNil)

		dvalue, derr := manager.Delete("platform-delete")
		So(dvalue, ShouldHaveSameTypeAs, &Value{})
		So(derr, ShouldBeNil)
		So(dvalue.Value(), ShouldBeBlank)

		gvalue, gerr := manager.Get("platform-delete")
		So(gvalue, ShouldHaveSameTypeAs, &Value{})
		So(gerr, ShouldNotBeNil)
	})

	Convey("Test If GetOrSet works", t, func() {
		value, err := manager.GetOrSet("platform-get-or-set", "Test Golang.hr Platform GetOrSet")
		So(value, ShouldHaveSameTypeAs, &Value{})
		So(err, ShouldBeNil)
		So(value.Value(), ShouldEqual, "Test Golang.hr Platform GetOrSet")
	})
}
