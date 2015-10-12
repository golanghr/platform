// Copyright 2015 The Golang.hr Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package config ...
package config

import (
	"fmt"
	"time"

	etcdc "github.com/coreos/etcd/client"
	"github.com/golanghr/platform/utils"
	"golang.org/x/net/context"
)

// ManagerInstance - Instance that manages configurations
type ManagerInstance struct {
	Config

	AutoSync   bool
	Env        string
	EtcdFolder string

	Client etcdc.Client
}

// Etcd - Will return instance of CoreOS Etcd
func (mi *ManagerInstance) Etcd() etcdc.Client {
	return mi.Client
}

// ShouldAutoSyncNodes - Basically return if we have permission to
// auto synchronize nodes or not. Used on package startup ...
func (mi *ManagerInstance) ShouldAutoSyncNodes() bool {
	return mi.AutoSync
}

// SyncNodes - Will initiate syncronization for client if configuration allows it...
func (mi *ManagerInstance) SyncNodes(interval time.Duration) (err error) {
	if mi.ShouldAutoSyncNodes() {
		go func() error {
			for {
				if err = mi.Client.AutoSync(context.Background(), interval); err != nil {
					return err
				}
			}
		}()
	}

	return
}

// NewManager - Return instance of configuration manager. Will return erorr
// in case of issues
func NewManager(cnf map[string]interface{}) (Manager, error) {

	if !utils.KeyInSlice("env", cnf) {
		return nil, fmt.Errorf("Could not find (key: env) within (config: %q). Plase make sure to read package documentation.", cnf)
	}

	if !utils.KeyInSlice("etcd", cnf) {
		return nil, fmt.Errorf("Could not find (key: etcd) within (config: %q). Plase make sure to read package documentation.", cnf)
	}

	autoSyncNodes := true

	if !utils.KeyInSlice("auto_sync", cnf) {
		autoSyncNodes = utils.GetBoolFromMap(cnf, "auto_load")
	}

	autoSyncInterval := 10 * time.Second

	if !utils.KeyInSlice("auto_sync_interval", cnf) {
		autoSyncInterval = utils.GetDurationFromMap(cnf, "auto_sync_interval")
	}

	etcdconf := cnf["etcd"].(map[string]interface{})

	var etcdcli etcdc.Client
	var err error

	if etcdcli, err = etcdc.New(etcdc.Config{
		Endpoints:               utils.GetStringsFromMap(etcdconf, "endpoints"),
		Transport:               etcdconf["transport"].(etcdc.CancelableTransport),
		HeaderTimeoutPerRequest: utils.GetDurationFromMap(etcdconf, "header_timeout_per_request"),
		Username:                utils.GetStringFromMap(etcdconf, "username"),
		Password:                utils.GetStringFromMap(etcdconf, "password"),
	}); err != nil {
		return nil, err
	}

	manager := Manager(&ManagerInstance{
		AutoSync:   autoSyncNodes,
		Env:        utils.GetStringFromMap(cnf, "env"),
		EtcdFolder: utils.GetStringFromMap(etcdconf, "folder"),
		Client:     etcdcli,
	})

	// This will spawn goroutine ...
	if err := manager.SyncNodes(autoSyncInterval); err != nil {
		return manager, err
	}

	return manager, nil
}
