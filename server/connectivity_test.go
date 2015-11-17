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

// Package server ...
package server

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestConnectivityStates(t *testing.T) {
	state := ConnectivityState{}

	Convey("Default connection state should match DOWN", t, func() {
		So(state.GetCurrentState(), ShouldEqual, state.GetStateByName("down"))
	})

	Convey("By setting invalid state we get error", t, func() {
		err := state.SetState("some-invalid-state")
		So(err, ShouldNotBeNil)
	})

	Convey("By invoke of GetStateByName() we get proper name", t, func() {
		So(state.GetStateByName("failed"), ShouldEqual, connectivityStates["failed"])
	})

	Convey("State successfully switches once we set it to down", t, func() {
		state.SetStateDown()
		So(state.GetCurrentState(), ShouldEqual, state.GetStateByName("down"))
	})

	Convey("State successfully switches once we set it to idle", t, func() {
		state.SetStateIdle()
		So(state.GetCurrentState(), ShouldEqual, state.GetStateByName("idle"))
	})

	Convey("State successfully switches once we set it to connecting", t, func() {
		state.SetStateConnecting()
		So(state.GetCurrentState(), ShouldEqual, state.GetStateByName("connecting"))
	})

	Convey("State successfully switches once we set it to ready", t, func() {
		state.SetStateReady()
		So(state.GetCurrentState(), ShouldEqual, state.GetStateByName("ready"))
	})

	Convey("State successfully switches once we set it to shutdown", t, func() {
		state.SetStateShutdown()
		So(state.GetCurrentState(), ShouldEqual, state.GetStateByName("shutdown"))
	})

	Convey("State successfully switches once we set it to failed", t, func() {
		state.SetStateFailed()
		So(state.GetCurrentState(), ShouldEqual, state.GetStateByName("failed"))
	})

	Convey("By invoking String() we are getting name of the current state", t, func() {
		So(state.String(), ShouldEqual, "failed")
	})
}
