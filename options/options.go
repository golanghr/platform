// Package options ...
// Copyright 2015 The Golang.hr Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package options

// Option - Option container and type helper
type Option struct {
	Key   string      `option:"key"`
	Value interface{} `option:"value"`
}

// New - Will return options based on adapter name and provided options.
// This is a shortcut towards multiple adapters that are supported within this package
func New(adapter string, opts map[string]interface{}) (Options, error) {
	var opt Options

	switch adapter {
	case ADAPTER_BASE:
		opt = Options(&AdapterBase{
			Adapter:    &Adapter{Name: ADAPTER_BASE},
			Collection: make(map[string]*Option),
		})

		if err := opt.SetMany(opts); err != nil {
			return nil, err
		}
	}

	return opt, nil
}
