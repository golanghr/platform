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

// Options - A options container and management struct
type Options struct {
	Collection map[string]*Option
}

// Exists - Will check whenever key exists in options collection
func (o *Options) Exists(key string) bool {
	return OptionExists(key, o.Collection)
}

// Get - Will retreive option from options collection or return nil in case that
// nothing is found.
func (o *Options) Get(key string) *Option {
	if !o.Exists(key) {
		return nil
	}
	return o.Collection[key]
}

// Set - Will set key with appropriate value in options collection
// @TODO - Figure out how to address error here
func (o *Options) Set(key string, value interface{}) error {
	o.Collection[key] = &Option{key, value}
	return nil
}

// SetMany - Will execute Set() method recursively
func (o *Options) SetMany(opts map[string]interface{}) (err error) {
	for optk, optv := range opts {
		if err = o.Set(optk, optv); err != nil {
			return
		}
	}
	return
}

// Unset - Will attempt to delete option from the collection. Will return boolean
// depending on if it's deleted or not.
func (o *Options) Unset(key string) bool {
	if !o.Exists(key) {
		return false
	}

	delete(o.Collection, key)
	return true
}

// New -
func New(opts map[string]interface{}) (*Options, error) {
	opt := &Options{Collection: make(map[string]*Option)}

	if err := opt.SetMany(opts); err != nil {
		return nil, err
	}

	return opt, nil
}
