// Package options ...
// Copyright 2015 The Golang.hr Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package options

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestOptions(t *testing.T) {
	opt, err := New(map[string]interface{}{
		"option_string":     "test",
		"option_int":        22,
		"option_uint":       uint(22),
		"option_int8":       int8(1),
		"option_uint8":      uint8(1),
		"option_int16":      int16(1),
		"option_uint16":     uint16(1),
		"option_int32":      int32(1),
		"option_uint32":     uint32(1),
		"option_int64":      int64(1),
		"option_uint64":     uint64(1),
		"option_float32":    float32(1.1),
		"option_float64":    float64(1.2),
		"option_complex64":  complex64(1.1),
		"option_complex128": complex128(1.2),
		"option_bool":       true,
	})

	Convey("Should return proper options without any errors", t, func() {
		So(*opt, ShouldHaveSameTypeAs, Options{})
		So(err, ShouldBeNil)
	})

	Convey("Should exist in collection", t, func() {
		So(opt.Exists("option_string"), ShouldBeTrue)
	})

	Convey("Should  notexist in collection", t, func() {
		So(opt.Exists("option_not_set"), ShouldBeFalse)
	})

	Convey("Should be able to get existing option", t, func() {
		option := opt.Get("option_string")
		So(*option, ShouldHaveSameTypeAs, Option{})
		So(option.Key, ShouldEqual, "option_string")
		So(option.Value, ShouldEqual, "test")
	})

	Convey("Should be able to set new option", t, func() {
		opt.Set("option_float", 2.2)

		option := opt.Get("option_float")
		So(*option, ShouldHaveSameTypeAs, Option{})
		So(option.Key, ShouldEqual, "option_float")
		So(option.Value, ShouldEqual, 2.2)
	})

	Convey("Should be able to set many options", t, func() {
		opt.SetMany(map[string]interface{}{
			"option_four": 11,
			"option_five": true,
		})

		optionInt := opt.Get("option_four")
		So(*optionInt, ShouldHaveSameTypeAs, Option{})
		So(optionInt.Key, ShouldEqual, "option_four")
		So(optionInt.Value, ShouldEqual, 11)

		optionBool := opt.Get("option_five")
		So(*optionBool, ShouldHaveSameTypeAs, Option{})
		So(optionBool.Key, ShouldEqual, "option_five")
		So(optionBool.Value, ShouldEqual, true)
	})

	Convey("Should be able to set new option", t, func() {
		opt.Unset("option_float")
		So(opt.Get("option_float"), ShouldBeNil)
	})

	Convey("Should be type of string", t, func() {
		So(opt.Get("option_string").String(), ShouldHaveSameTypeAs, "")
	})

	Convey("Should be type of bool", t, func() {
		So(opt.Get("option_bool").Bool(), ShouldHaveSameTypeAs, false)
	})

	Convey("Should be type of int", t, func() {
		So(opt.Get("option_int").Int(), ShouldHaveSameTypeAs, 1)
	})

	Convey("Should be type of uint", t, func() {
		So(opt.Get("option_uint").UInt(), ShouldHaveSameTypeAs, uint(1))
	})

	Convey("Should be type of int8", t, func() {
		So(opt.Get("option_int8").Int8(), ShouldHaveSameTypeAs, int8(1))
	})

	Convey("Should be type of uint8", t, func() {
		So(opt.Get("option_uint8").UInt8(), ShouldHaveSameTypeAs, uint8(1))
	})

	Convey("Should be type of int16", t, func() {
		So(opt.Get("option_int16").Int16(), ShouldHaveSameTypeAs, int16(1))
	})

	Convey("Should be type of uint16", t, func() {
		So(opt.Get("option_uint16").UInt16(), ShouldHaveSameTypeAs, uint16(1))
	})

	Convey("Should be type of int32", t, func() {
		So(opt.Get("option_int32").Int32(), ShouldHaveSameTypeAs, int32(1))
	})

	Convey("Should be type of uint16", t, func() {
		So(opt.Get("option_uint32").UInt32(), ShouldHaveSameTypeAs, uint32(1))
	})

	Convey("Should be type of int64", t, func() {
		So(opt.Get("option_int64").Int64(), ShouldHaveSameTypeAs, int64(1))
	})

	Convey("Should be type of uint64", t, func() {
		So(opt.Get("option_uint64").UInt64(), ShouldHaveSameTypeAs, uint64(1))
	})

	Convey("Should be type of float - *float64*", t, func() {
		So(opt.Get("option_float64").Float(), ShouldHaveSameTypeAs, float64(1.2))
	})

	Convey("Should be type of float32", t, func() {
		So(opt.Get("option_float32").Float32(), ShouldHaveSameTypeAs, float32(1.5))
	})

	Convey("Should be type of float64", t, func() {
		So(opt.Get("option_float64").Float64(), ShouldHaveSameTypeAs, float64(1.6))
	})

	Convey("Should be type of complex64", t, func() {
		So(opt.Get("option_complex64").Complex64(), ShouldHaveSameTypeAs, complex64(1.5))
	})

	Convey("Should be type of complex128", t, func() {
		So(opt.Get("option_complex128").Complex128(), ShouldHaveSameTypeAs, complex128(1.6))
	})
}
