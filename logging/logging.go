// Copyright 2015 The Golang.hr Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package logging ...
package logging

import (
	"github.com/Sirupsen/logrus"
	logstashf "github.com/Sirupsen/logrus/formatters/logstash"
	"github.com/golanghr/platform/utils"
)

// Logging - A small wrapper around logrus.Logger
type Logging struct {
	*logrus.Logger
	Config map[string]interface{}
}

// New - Will create new Logging instance
func New(config map[string]interface{}) Logging {

	logger := Logging{
		Logger: logrus.New(),
		Config: config,
	}

	if utils.KeyInSlice("formatter", config) {
		switch config["formatter"].(string) {
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

	if utils.KeyInSlice("level", config) {
		logger.Level = config["level"].(logrus.Level)
	}

	return logger
}
