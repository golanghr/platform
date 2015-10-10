// Copyright 2015 The Golang.hr Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package utils ...
package utils

import "time"

// StringInSlice - Will check if string in list.
// This is equivalent to python if x in []
func StringInSlice(str string, list []string) bool {
	for _, value := range list {
		if value == str {
			return true
		}
	}
	return false
}

// KeyInSlice - Will check if key in list.
// This is equivalent to python if x in []
func KeyInSlice(str string, list map[string]interface{}) bool {
	for key, _ := range list {
		if key == str {
			return true
		}
	}
	return false
}

// GetString - Will return string from key-value based storage
func GetString(store map[string]interface{}, key string) string {
	return store[key].(string)
}

// GetStrings - Will return strings from key-value based storage
func GetStrings(store map[string]interface{}, key string) []string {
	return store[key].([]string)
}

// GetDuration - Will return time.Duration from key-value based storage
func GetDuration(store map[string]interface{}, key string) time.Duration {
	return store[key].(time.Duration)
}
