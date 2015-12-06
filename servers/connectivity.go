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

// Package servers ...
package servers

import "fmt"

// ConnectivityState - Designed to keep track of server current connection state
type ConnectivityState struct {
	state int64
}

// GetStateByName - Will attempt to figure out state by name. If state is not
// found it will return 0 which equals to DOWN
func (cs *ConnectivityState) GetStateByName(state string) int64 {
	if res, ok := connectivityStates[state]; ok {
		return res
	}

	return 0
}

// GetCurrentState - Will return current state. If state is not set, it will
// return 0 which equals to state DOWN
func (cs *ConnectivityState) GetCurrentState() int64 {
	return cs.state
}

// SetState - Will connection state. In case that state is not available it will
// return error
func (cs *ConnectivityState) SetState(state string) error {
	if _, ok := connectivityStates[state]; !ok {
		return fmt.Errorf("Invalid connectivity state provided (provided: %s)", state)
	}

	cs.state = connectivityStates[state]
	return nil
}

// SetStateDown - Will set current state to DOWN
func (cs *ConnectivityState) SetStateDown() {
	cs.SetState("down")
}

// SetStateIdle - Will set current state to IDLE
func (cs *ConnectivityState) SetStateIdle() {
	cs.SetState("idle")
}

// SetStateConnecting - Will set current state to CONNECTING
func (cs *ConnectivityState) SetStateConnecting() {
	cs.SetState("connecting")
}

// SetStateReady - Will set current state to READY
func (cs *ConnectivityState) SetStateReady() {
	cs.SetState("ready")
}

// SetStateFailed - Will set current state to FAILED
func (cs *ConnectivityState) SetStateFailed() {
	cs.SetState("failed")
}

// SetStateShutdown - Will set current state to SHUTDOWN
func (cs *ConnectivityState) SetStateShutdown() {
	cs.SetState("shutdown")
}

// String - Will return name of the state. If state cannot be found, empty
// string will be returned
func (cs *ConnectivityState) String() string {
	for state, id := range connectivityStates {
		if id == cs.GetCurrentState() {
			return state
		}
	}

	return ""
}
