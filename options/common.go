// Package options ...
// Copyright 2015 The Golang.hr Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package options

// OptionExists - Will check if key exists in options map.
func OptionExists(str string, opts map[string]*Option) bool {
	for key := range opts {
		if key == str {
			return true
		}
	}
	return false
}
