// Copyright 2015 The Golang.hr Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package service ...
package service

import (
	"testing"

	"github.com/golanghr/platform/options"
	. "github.com/smartystreets/goconvey/convey"
)

var (
	testEnv        = "test_local"
	testEtcdFolder = "golanghr-test"
)

func TestNewServiceCreation(t *testing.T) {
	Convey("By passing proper details we are getting valid service", t, func() {

		opts, err := options.New("memo", map[string]interface{}{})
		So(err, ShouldBeNil)

		serv, err := New(opts)

		So(err, ShouldBeNil)
		So(serv, ShouldHaveSameTypeAs, Service{})
	})
}
