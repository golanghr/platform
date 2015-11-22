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

// Package manager ...
package manager

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/golanghr/platform/logging"
	"github.com/golanghr/platform/options"
	"github.com/golanghr/platform/server"
	"github.com/golanghr/platform/service"
)

// Manager -
type Manager struct {
	service.Servicer
	options.Options
	*logging.Entry

	Servers map[string]server.Serverer

	// InterruptWaitTimeout
	InterruptWaitTimeout time.Duration

	// Interrupt -
	Interrupt chan os.Signal

	// Used while starting and stopping servers ...
	wg sync.WaitGroup
}

// Attach - Will attach server to the servers map
func (m *Manager) Attach(server string, i server.Serverer) error {
	if _, ok := m.Servers[server]; ok {
		return fmt.Errorf("Could not attach server `%s` as one is already attached.", server)
	}

	m.Servers[server] = i
	return nil
}

// Remove - Will remove server from servers map
func (m *Manager) Remove(server string) error {
	if _, ok := m.Servers[server]; !ok {
		return fmt.Errorf("Could not remove server `%s` as one does not seem to be set.", server)
	}

	delete(m.Servers, server)
	return nil
}

// Available - Will return list of all available servers
func (m *Manager) Available() map[string]server.Serverer {
	return m.Servers
}

// Start - Will loop throught each of servers and attempt to start them up.
// In case that channel receives single error it will kill function otherwise, will listen
// forever
func (m *Manager) Start() error {
	m.Info("Starting service management...")
	errors := make(chan error)

	signal.Notify(m.GetInterruptChan(), os.Interrupt)
	signal.Notify(m.GetInterruptChan(), syscall.SIGTERM)

	go m.HandleInterrupt()

	for _, serv := range m.Servers {
		go func(serv server.Serverer) {
			if err := serv.Start(); err != nil {
				errors <- err
			}
		}(serv)
	}

	select {
	case err := <-errors:
		return err
	case <-m.GetInterruptChan():
		m.Warn("Service management interrupt signal caught. Giving it few seconds before we shut down permanently...")
		close(m.GetInterruptChan())

		time.Sleep(m.InterruptWaitTimeout)
		m.Info("All done :( We go underground now for good...")
		return nil
	}
}

// HandleInterrupt - Will wait for interrupt channel to quit and once it does
// will execute
func (m *Manager) HandleInterrupt() {
	m.Warn("Listening for service interrupt signal ...")

	<-m.GetInterruptChan()
	m.Stop()
}

// Stop - Will gracefully stop all exisitng and running servers
// In case of any errors, we will return one :)
func (m *Manager) Stop() error {
	m.Info("Stopping service management...")
	errors := make(chan error)

	for name, serv := range m.Servers {
		m.Infof("About to start shutting down (server: %s)", name)
		m.wg.Add(1)

		go func(serv server.Serverer) {
			if err := serv.Stop(); err != nil {
				errors <- err
			}
			m.wg.Done()
		}(serv)
	}

	m.wg.Wait()

	select {
	case err := <-errors:
		return err
	}
}

// New -
func New(serv service.Servicer, opts options.Options, logger *logging.Entry) (Managerer, error) {
	waitTimeout := DefaultInterruptWaitTimeout

	if miwt, ok := opts.Get("manager-interrupt-wait-timeout"); ok {
		waitTimeout = miwt.Int()
	}

	return Managerer(&Manager{
		Servicer:             serv,
		Options:              opts,
		Entry:                logger,
		Servers:              make(map[string]server.Serverer),
		InterruptWaitTimeout: time.Duration(waitTimeout) * time.Second,
		Interrupt:            serv.GetInterruptChan(),
	}), nil
}
