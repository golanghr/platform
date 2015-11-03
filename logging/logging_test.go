// Copyright 2015 The Golang.hr Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
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

	opts, err := options.New(map[string]interface{}{
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

		opts, err := options.New(map[string]interface{}{
			"formatter": "json",
			"level":     logrus.FatalLevel,
		})
		So(err, ShouldBeNil)

		log := New(opts)
		So(log.Formatter, ShouldHaveSameTypeAs, &logrus.JSONFormatter{})
	})

	Convey("Test Logger Logstash Formatter", t, func() {
		opts, err := options.New(map[string]interface{}{
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

	opts, err := options.New(map[string]interface{}{
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
