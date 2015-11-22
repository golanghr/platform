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

// Package logging ...
package logging

import (
	"bytes"
	"testing"

	"github.com/MatejB/reactLog"
	"github.com/Sirupsen/logrus"
	logstashf "github.com/Sirupsen/logrus/formatters/logstash"

	"github.com/golanghr/platform/options"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNewLoggingInstance(t *testing.T) {

	opts, err := options.New("base", map[string]interface{}{
		"formatter": "text",
		"level":     logrus.FatalLevel,
	})

	Convey("Logging options should be successfully created without any errors", t, func() {
		So(err, ShouldBeNil)
	})

	log := New(opts)

	Convey("Test Required Logging Type", t, func() {
		So(log, ShouldHaveSameTypeAs, Logging{})
	})

	Convey("Test Required Logging Level", t, func() {
		So(log.Level, ShouldEqual, logrus.FatalLevel)
	})

	Convey("Test Logger Formatter", t, func() {
		So(log.Formatter, ShouldHaveSameTypeAs, &logrus.TextFormatter{})
	})

	Convey("Test Logger JSON Formatter", t, func() {

		opts, err := options.New("base", map[string]interface{}{
			"formatter": "json",
			"level":     logrus.FatalLevel,
		})
		So(err, ShouldBeNil)

		log := New(opts)
		So(log.Formatter, ShouldHaveSameTypeAs, &logrus.JSONFormatter{})
	})

	Convey("Test Logger Logstash Formatter", t, func() {
		opts, err := options.New("base", map[string]interface{}{
			"formatter": "logstash",
			"level":     logrus.FatalLevel,
		})
		So(err, ShouldBeNil)

		log := New(opts)
		So(log.Formatter, ShouldHaveSameTypeAs, &logstashf.LogstashFormatter{})
	})
}

func TestMiddlewareIntegration(t *testing.T) {
	generalLogContainer := &bytes.Buffer{}
	logContainerForUser107 := &bytes.Buffer{}

	rlog := reactLog.New(generalLogContainer)
	rlog.AddReaction("user ID 107", &reactLog.Redirect{logContainerForUser107})

	opts, err := options.New("base", map[string]interface{}{
		"formatter": "text",
		"level":     logrus.FatalLevel,
	})

	Convey("Middleware logging options should be successfully created without any errors", t, func() {
		So(err, ShouldBeNil)
	})

	log := New(opts)

	log.SetOutput(rlog)

	log.Info("This is normal log")
	log.Info("This is log that concers user ID 107 with important data")

	Convey("Test logger middleware redirect", t, func() {
		So(generalLogContainer.String(), ShouldContainSubstring, "This is normal log")
	})

	Convey("Test logger middleware redirect", t, func() {
		So(logContainerForUser107.String(), ShouldContainSubstring, "This is log that concers user ID 107 with important data")
	})
}
