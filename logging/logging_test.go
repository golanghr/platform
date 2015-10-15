// Copyright 2015 The Golang.hr Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package logging ...
package logging

import (
	"bytes"
	"testing"

	"github.com/MatejB/reactLog"
	"github.com/Sirupsen/logrus"
	logstashf "github.com/Sirupsen/logrus/formatters/logstash"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNewLoggingInstance(t *testing.T) {
	log := New(map[string]interface{}{
		"formatter": "text",
		"level":     logrus.FatalLevel,
	})

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
		log := New(map[string]interface{}{
			"formatter": "json",
		})

		So(log.Formatter, ShouldHaveSameTypeAs, &logrus.JSONFormatter{})
	})

	Convey("Test Logger Logstash Formatter", t, func() {
		log := New(map[string]interface{}{
			"formatter": "logstash",
		})

		So(log.Formatter, ShouldHaveSameTypeAs, &logstashf.LogstashFormatter{})
	})
}

func TestMiddlewareIntegration(t *testing.T) {
	generalLogContainer := &bytes.Buffer{}
	logContainerForUser107 := &bytes.Buffer{}

	rlog := reactLog.New(generalLogContainer)
	rlog.AddReaction("user ID 107", &reactLog.Redirect{logContainerForUser107})

	log := New(map[string]interface{}{
		"formatter": "text",
	})

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
