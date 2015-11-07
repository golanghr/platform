// Package options ...
// Copyright 2015 The Golang.hr Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package options

import "fmt"

// New - Will return options based on adapter name and provided options.
// This is a shortcut towards multiple adapters that are supported within this package
//
// Example
// 	opt, err := options.New("map", map[string]interface{}{
// 	  "test_option": "test_string_value",
// 	})
func New(adapter string, opts map[string]interface{}) (Options, error) {
	var opt Options

	if _, exists := AvailableAdapters[adapter]; !exists {
		return nil, fmt.Errorf("Invalid options adapter provided. We do not support '%s'", adapter)
	}

	switch adapter {
	case "memo":
		opt = &Memo{
			Adapter:    &Adapter{Name: "map"},
			Collection: make(map[string]*Option),
		}

		if err := opt.SetMany(opts); err != nil {
			return nil, err
		}
	}

	return opt, nil
}
