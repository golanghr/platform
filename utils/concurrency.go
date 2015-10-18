// Copyright 2015 The Golang.hr Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package utils ...
package utils

import (
	"os"
	"runtime"
	"strconv"
)

// GetProcessCount - Get Process count defined by ENV or by NumCPU()
func GetProcessCount(env string) int {

	envName := "GH_GO_MAX_PROCS"

	if env != "" {
		envName = env
	}

	pc, err := strconv.Atoi(os.Getenv(envName))

	if err != nil {
		pc = runtime.NumCPU()
	}

	return int(pc)
}
