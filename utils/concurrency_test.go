// Copyright 2015 The Golang.hr Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package utils ...
package utils

import (
	"os"
	"runtime"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetProcessCount(t *testing.T) {
	Convey("Test Get Process Count Without Environment Variables", t, func() {
		So(GetProcessCount(""), ShouldEqual, runtime.NumCPU())
	})

	Convey("Test Get Process Count With Environment Variables", t, func() {
		os.Setenv("GHR_TEST_PROCESS_COUNT", "4")
		So(GetProcessCount("GHR_TEST_PROCESS_COUNT"), ShouldEqual, 4)
	})
}
