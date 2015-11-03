// Copyright 2015 The Golang.hr Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package service ...
package service

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	testEnv        = "test_local"
	testEtcdFolder = "golanghr-test"
)

func TestNewServiceCreation(t *testing.T) {
	Convey("Test Required Logging Type", t, func() {

		serv, err := getService(t)

		So(err, ShouldBeNil)
		So(serv, ShouldHaveSameTypeAs, Instance{})
	})
}
