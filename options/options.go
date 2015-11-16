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

// Package options ...
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
