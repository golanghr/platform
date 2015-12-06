/*
Copyright (c) 2015 Golang Croatia
All rights reserved.

The MIT License (MIT)

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

// Package options ...
package options

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestOptionsBaseAdapter(t *testing.T) {
	opt, err := New("memo", map[string]interface{}{
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
		So(opt, ShouldHaveSameTypeAs, &Memo{})
		So(err, ShouldBeNil)
	})

	Convey("Should exist in collection", t, func() {
		_, ok := opt.Get("option_string")
		So(ok, ShouldBeTrue)
	})

	Convey("Should  notexist in collection", t, func() {
		option, ok := opt.Get("option_not_set")
		So(option, ShouldBeNil)
		So(ok, ShouldBeFalse)
	})

	Convey("Should be able to get existing option", t, func() {
		option, _ := opt.Get("option_string")
		So(*option, ShouldHaveSameTypeAs, Option{})
		So(option.Key, ShouldEqual, "option_string")
		So(option.Value, ShouldEqual, "test")
	})

	Convey("Should be able to set new option", t, func() {
		opt.Set("option_float", 2.2)

		option, _ := opt.Get("option_float")
		So(*option, ShouldHaveSameTypeAs, Option{})
		So(option.Key, ShouldEqual, "option_float")
		So(option.Value, ShouldEqual, 2.2)
	})

	Convey("Should be able to set many options", t, func() {
		opt.SetMany(map[string]interface{}{
			"option_four": 11,
			"option_five": true,
		})

		optionInt, _ := opt.Get("option_four")
		So(*optionInt, ShouldHaveSameTypeAs, Option{})
		So(optionInt.Key, ShouldEqual, "option_four")
		So(optionInt.Value, ShouldEqual, 11)

		optionBool, _ := opt.Get("option_five")
		So(*optionBool, ShouldHaveSameTypeAs, Option{})
		So(optionBool.Key, ShouldEqual, "option_five")
		So(optionBool.Value, ShouldEqual, true)
	})

	Convey("Should be able to set new option", t, func() {
		opt.Unset("option_float")
		option, _ := opt.Get("option_float")
		So(option, ShouldBeNil)
	})

	Convey("Should be type of string", t, func() {
		option, _ := opt.Get("option_string")
		So(option.String(), ShouldHaveSameTypeAs, "")
	})

	Convey("Should be type of bool", t, func() {
		option, _ := opt.Get("option_bool")
		So(option.Bool(), ShouldHaveSameTypeAs, false)
	})

	Convey("Should be type of int", t, func() {
		option, _ := opt.Get("option_int")
		So(option.Int(), ShouldHaveSameTypeAs, 1)
	})

	Convey("Should be type of uint", t, func() {
		option, _ := opt.Get("option_uint")
		So(option.UInt(), ShouldHaveSameTypeAs, uint(1))
	})

	Convey("Should be type of int8", t, func() {
		option, _ := opt.Get("option_int8")
		So(option.Int8(), ShouldHaveSameTypeAs, int8(1))
	})

	Convey("Should be type of uint8", t, func() {
		option, _ := opt.Get("option_uint8")
		So(option.UInt8(), ShouldHaveSameTypeAs, uint8(1))
	})

	Convey("Should be type of int16", t, func() {
		option, _ := opt.Get("option_int16")
		So(option.Int16(), ShouldHaveSameTypeAs, int16(1))
	})

	Convey("Should be type of uint16", t, func() {
		option, _ := opt.Get("option_uint16")
		So(option.UInt16(), ShouldHaveSameTypeAs, uint16(1))
	})

	Convey("Should be type of int32", t, func() {
		option, _ := opt.Get("option_int32")
		So(option.Int32(), ShouldHaveSameTypeAs, int32(1))
	})

	Convey("Should be type of uint16", t, func() {
		option, _ := opt.Get("option_uint32")
		So(option.UInt32(), ShouldHaveSameTypeAs, uint32(1))
	})

	Convey("Should be type of int64", t, func() {
		option, _ := opt.Get("option_int64")
		So(option.Int64(), ShouldHaveSameTypeAs, int64(1))
	})

	Convey("Should be type of uint64", t, func() {
		option, _ := opt.Get("option_uint64")
		So(option.UInt64(), ShouldHaveSameTypeAs, uint64(1))
	})

	Convey("Should be type of float - *float64*", t, func() {
		option, _ := opt.Get("option_float64")
		So(option.Float(), ShouldHaveSameTypeAs, float64(1.2))
	})

	Convey("Should be type of float32", t, func() {
		option, _ := opt.Get("option_float32")
		So(option.Float32(), ShouldHaveSameTypeAs, float32(1.5))
	})

	Convey("Should be type of float64", t, func() {
		option, _ := opt.Get("option_float64")
		So(option.Float64(), ShouldHaveSameTypeAs, float64(1.6))
	})

	Convey("Should be type of complex64", t, func() {
		option, _ := opt.Get("option_complex64")
		So(option.Complex64(), ShouldHaveSameTypeAs, complex64(1.5))
	})

	Convey("Should be type of complex128", t, func() {
		option, _ := opt.Get("option_complex128")
		So(option.Complex128(), ShouldHaveSameTypeAs, complex128(1.6))
	})
}
