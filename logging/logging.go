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
	"io"

	"github.com/Sirupsen/logrus"
	logstashf "github.com/Sirupsen/logrus/formatters/logstash"
	"github.com/golanghr/platform/options"
)

// Entry -
type Entry struct {
	*logrus.Entry
}

// Logging - A small wrapper around logrus.Logger
type Logging struct {
	*logrus.Logger
	options.Options
}

// WithFields -
func (l *Logging) WithFields(fields logrus.Fields) *Entry {
	return &Entry{
		l.Logger.WithFields(fields),
	}
}

// New - Will create new Logging instance
func New(opts options.Options) Logging {

	logger := Logging{
		Logger:  logrus.New(),
		Options: opts,
	}

	if formatter, ok := opts.Get("formatter"); ok {
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

	if level, ok := opts.Get("level"); ok {
		logger.Level = level.Interface().(logrus.Level)
	}

	return logger
}

// SetOutput of logger
func (l *Logging) SetOutput(w io.Writer) {
	l.Logger.Out = w
}
