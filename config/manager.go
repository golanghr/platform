// Copyright 2015 The Golang.hr Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package config ...
package config

import (
	"fmt"

	etcdc "github.com/coreos/etcd/client"
	"github.com/golanghr/platform/utils"
)

// ManagerInstance -
type ManagerInstance struct {
	Env        string
	EtcdFolder string

	EtcdClient etcdc.Client
}

// Etcd -
func (mi *ManagerInstance) Etcd() etcdc.Client {
	return mi.EtcdClient
}

// NewManager -
func NewManager(cnf map[string]interface{}) (Manager, error) {

	if !utils.KeyInSlice("env", cnf) {
		return nil, fmt.Errorf("Could not find (key: env) within (config: %q). Plase make sure to read package documentation.", cnf)
	}

	if !utils.KeyInSlice("etcd", cnf) {
		return nil, fmt.Errorf("Could not find (key: etcd) within (config: %q). Plase make sure to read package documentation.", cnf)
	}

	etcdconf := cnf["etcd"].(map[string]interface{})

	etcdcli, err := etcdc.New(etcdc.Config{
		Endpoints:               utils.GetStrings(etcdconf, "endpoints"),
		Transport:               etcdconf["transport"].(etcdc.CancelableTransport),
		HeaderTimeoutPerRequest: utils.GetDuration(etcdconf, "header_timeout_per_request"),
		Username:                utils.GetString(etcdconf, "username"),
		Password:                utils.GetString(etcdconf, "password"),
	})

	if err != nil {
		return nil, err
	}

	manager := Manager(&ManagerInstance{
		Env:        utils.GetString(cnf, "env"),
		EtcdFolder: utils.GetString(etcdconf, "folder"),
		EtcdClient: etcdcli,
	})

	return manager, nil
}
