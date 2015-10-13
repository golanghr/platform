// Copyright 2015 The Golang.hr Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package utils ...
package utils

import "os"

// GetFromEnvOr - Will attempt to return environment variable value or
// fail back to provided defaults
func GetFromEnvOr(env string, def string) string {
	if res := os.Getenv(env); res != "" {
		return res
	}
	return def
}
