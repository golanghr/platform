// Copyright 2015 The Golang.hr Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package logging ...
package logging

import (
	"testing"

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
