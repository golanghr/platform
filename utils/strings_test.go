// Copyright 2015 The Golang.hr Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package utils ...
package utils

import (
	"fmt"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestStringInSlice(t *testing.T) {
	Convey("Test String Is In Slice", t, func() {
		So(StringInSlice("test", []string{"test", "testing"}), ShouldBeTrue)
	})

	Convey("Test String Is Not In Slice", t, func() {
		So(StringInSlice("test-not-in-slice", []string{"test", "testing"}), ShouldBeFalse)
	})
}

func TestKeyInSlice(t *testing.T) {
	kvmap := map[string]interface{}{
		"test-one": "lorem",
		"test-two": "lorem-two",
	}

	Convey("Test Key Is In Slice", t, func() {
		So(KeyInSlice("test-one", kvmap), ShouldBeTrue)
	})

	Convey("Test String Is Not In Slice", t, func() {
		So(KeyInSlice("test-not-in-slice", kvmap), ShouldBeFalse)
	})
}

func TestGetStringFromMap(t *testing.T) {
	kvmap := map[string]interface{}{
		"test-one": "lorem",
	}

	Convey("Test Get Existing String From", t, func() {
		So(GetStringFromMap(kvmap, "test-one"), ShouldEqual, kvmap["test-one"])
	})

	Convey("Test Get Non-Existing String From", t, func() {
		So(GetStringFromMap(kvmap, "test-one-not-here"), ShouldEqual, "")
	})
}

func TestGetStringsFromMap(t *testing.T) {
	kvmap := map[string]interface{}{
		"test-one": []string{"lorem", "ipsum", "dolor"},
	}

	Convey("Test Get Existing Strings From", t, func() {

		// Due to some unknown (yet to me) reason So([]string, ShouldEqual, []string) do not
		// match and are not equal ...
		So(
			fmt.Sprintf("%v", GetStringsFromMap(kvmap, "test-one")),
			ShouldEqual,
			fmt.Sprintf("%v", kvmap["test-one"].([]string)),
		)
	})

	Convey("Test Get Non-Existing Strings From", t, func() {
		So(GetStringsFromMap(kvmap, "test-one-not-here"), ShouldBeNil)
	})
}

func TestGetDurationFromMap(t *testing.T) {
	kvmap := map[string]interface{}{
		"test-one": 10 * time.Second,
	}

	Convey("Test Get Existing Bool From", t, func() {
		So(GetDurationFromMap(kvmap, "test-one"), ShouldEqual, kvmap["test-one"])
	})

	Convey("Test Get Non-Existing Bool From", t, func() {
		So(GetDurationFromMap(kvmap, "test-one-not-here"), ShouldEqual, 0)
	})
}

func TestGetBoolFromMap(t *testing.T) {
	kvmap := map[string]interface{}{
		"test-one": true,
	}

	Convey("Test Get Existing Bool From", t, func() {
		So(GetBoolFromMap(kvmap, "test-one"), ShouldBeTrue)
	})

	Convey("Test Get Non-Existing Bool From", t, func() {
		So(GetBoolFromMap(kvmap, "test-one-not-here"), ShouldBeFalse)
	})
}
