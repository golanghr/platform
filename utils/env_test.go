// Copyright 2015 The Golang.hr Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package utils ...
package utils

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetFromEnvOr(t *testing.T) {
	Convey("Test Get From Env Or Without Environment Variables", t, func() {
		So(GetFromEnvOr("GHR_TEST_GET_ENV_OR", "default value"), ShouldEqual, "default value")
	})

	Convey("Test Get From Env Or With Environment Variables", t, func() {
		os.Setenv("GHR_TEST_GET_ENV_OR", "5")
		So(GetFromEnvOr("GHR_TEST_GET_ENV_OR", ""), ShouldEqual, "5")
	})
}
