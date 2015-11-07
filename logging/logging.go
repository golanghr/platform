// Copyright 2015 The Golang.hr Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package logging ...
package logging

import (
	"io"

	"github.com/Sirupsen/logrus"
	logstashf "github.com/Sirupsen/logrus/formatters/logstash"
	"github.com/golanghr/platform/options"
)

// Logging - A small wrapper around logrus.Logger
type Logging struct {
	*logrus.Logger
	options.Options
}

// New - Will create new Logging instance
func New(opts options.Options) Logging {

	logger := Logging{
		Logger:  logrus.New(),
		Options: opts,
	}

	formatter, exists := opts.Get("formatter")
	level, levelExists := opts.Get("level")

	if exists {
		switch formatter.String() {
		case "text":
			logger.Formatter = new(logrus.TextFormatter)
		case "json":
			logger.Formatter = new(logrus.JSONFormatter)
		case "logstash":
			logger.Formatter = new(logstashf.LogstashFormatter)
		default:
			logger.Formatter = new(logrus.TextFormatter)
		}
	}

	if levelExists {
		logger.Level = level.Interface().(logrus.Level)
	}

	return logger
}

// SetOutput of logger
func (l *Logging) SetOutput(w io.Writer) {
	l.Logger.Out = w
}
