// Copyright 2015 The Golang.hr Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package config ...
package config

import (
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

func TestManagerConfigDefaults(t *testing.T) {

}

// TestNewEventCreation -
func TestNewManagerCreation(t *testing.T) {
	Convey("Test If Manager/Etcd", t, func() {
		manager, err := NewManager(map[string]interface{}{
			"env":                testEnv,
			"auto_sync":          true,
			"auto_sync_interval": 10 * time.Second,
			"etcd": map[string]interface{}{
				"folder":                     testEtcdFolder,
				"endpoints":                  []string{"127.0.0.1:2379"},
				"transport":                  etcdc.DefaultTransport,
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

		manager, err := NewManager(map[string]interface{}{
			"env":                testEnv,
			"auto_sync":          true,
			"auto_sync_interval": 10 * time.Second,
			"etcd": map[string]interface{}{
				"folder":                     testEtcdFolder,
				"endpoints":                  []string{"127.0.0.1:2379"},
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
