// Copyright 2015 The Golang.hr Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package utils ...
package utils

import "time"

// StringInSlice - Will check if string in list. This is equivalent to python if x in []
func StringInSlice(str string, list []string) bool {
	for _, value := range list {
		if value == str {
			return true
		}
	}
	return false
}

// KeyInSlice - Will check if key in list. This is equivalent to python if x in []
func KeyInSlice(str string, list map[string]interface{}) bool {
	for key, _ := range list {
		if key == str {
			return true
		}
	}
	return false
}

// GetStringFromMap - Will return string value from key-value based storage
func GetStringFromMap(store map[string]interface{}, key string) string {
	if !KeyInSlice(key, store) {
		return ""
	}

	return store[key].(string)
}

// GetStringsFromMap - Will return strings value from key-value based storage
func GetStringsFromMap(store map[string]interface{}, key string) []string {
	if !KeyInSlice(key, store) {
		return nil
	}

	return store[key].([]string)
}

// GetDurationFromMap - Will return time.Duration value from key-value based storage
func GetDurationFromMap(store map[string]interface{}, key string) time.Duration {
	if !KeyInSlice(key, store) {
		return 0
	}

	return store[key].(time.Duration)
}

// GetBoolFromMap - Will return bool value from key-value based storage
func GetBoolFromMap(store map[string]interface{}, key string) bool {
	if !KeyInSlice(key, store) {
		return false
	}

	return store[key].(bool)
}
