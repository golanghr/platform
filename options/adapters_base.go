// Package options ...
// Copyright 2015 The Golang.hr Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package options

// Memo - A options container and management struct
type Memo struct {
	*Adapter
	Collection map[string]*Option
}

// Exists - Will check whenever key exists in options collection
func (m *Memo) Exists(key string) bool {
	if _, exists := m.Collection[key]; !exists {
		return false
	}

	return true
}

// Get - Will retreive option from options collection or return nil in case that
// nothing is found.
func (m *Memo) Get(key string) *Option {
	if !m.Exists(key) {
		return nil
	}
	return m.Collection[key]
}

// GetMany - Will attempt to get many keys. In case that key does not exist it will
// ommit that key
func (m *Memo) GetMany(keys []string) []*Option {
	var opts []*Option

	for _, key := range keys {
		if m.Exists(key) {
			opts = append(opts, m.Get(key))
		}
	}

	return opts
}

// Set - Will set key with appropriate value in options collection
// @TODO - Figure out how to address error here
func (m *Memo) Set(key string, value interface{}) error {
	m.Collection[key] = &Option{key, value}
	return nil
}

// SetMany - Will execute Set() method recursively
func (m *Memo) SetMany(opts map[string]interface{}) (err error) {
	for optk, optv := range opts {
		if err = m.Set(optk, optv); err != nil {
			return
		}
	}
	return
}

// Unset - Will attempt to delete option from the collection. Will return boolean
// depending on if it's deleted or not.
func (m *Memo) Unset(key string) bool {
	if !m.Exists(key) {
		return false
	}

	delete(m.Collection, key)
	return true
}

// Interface - Will return Base Adapter as interface. Not useful for anything
// but is useful for adapters such as etcd
func (m *Memo) Interface() interface{} {
	return m
}
