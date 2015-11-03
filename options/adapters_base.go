// Package options ...
// Copyright 2015 The Golang.hr Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package options

// AdapterBase - A options container and management struct
type AdapterBase struct {
	*Adapter
	Collection map[string]*Option
}

// Exists - Will check whenever key exists in options collection
func (ab *AdapterBase) Exists(key string) bool {
	return OptionExists(key, ab.Collection)
}

// Get - Will retreive option from options collection or return nil in case that
// nothing is found.
func (ab *AdapterBase) Get(key string) *Option {
	if !ab.Exists(key) {
		return nil
	}
	return ab.Collection[key]
}

// GetMany - Will attempt to get many keys. In case that key does not exist it will
// ommit that key
func (ab *AdapterBase) GetMany(keys []string) []*Option {
	var opts []*Option

	for _, key := range keys {
		if ab.Exists(key) {
			opts = append(opts, ab.Get(key))
		}
	}

	return opts
}

// Set - Will set key with appropriate value in options collection
// @TODO - Figure out how to address error here
func (ab *AdapterBase) Set(key string, value interface{}) error {
	ab.Collection[key] = &Option{key, value}
	return nil
}

// SetMany - Will execute Set() method recursively
func (ab *AdapterBase) SetMany(opts map[string]interface{}) (err error) {
	for optk, optv := range opts {
		if err = ab.Set(optk, optv); err != nil {
			return
		}
	}
	return
}

// Unset - Will attempt to delete option from the collection. Will return boolean
// depending on if it's deleted or not.
func (ab *AdapterBase) Unset(key string) bool {
	if !ab.Exists(key) {
		return false
	}

	delete(ab.Collection, key)
	return true
}

// Interface - Will return Base Adapter as interface. Not useful for anything
// but is useful for adapters such as etcd
func (ab *AdapterBase) Interface() interface{} {
	return ab
}
