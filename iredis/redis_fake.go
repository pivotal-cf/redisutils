// This file was generated by counterfeiter
// counterfeiter -o iredis/redis_fake.go --fake-name Fake iredis/redis.go Redis

package iredis

import (
	"sync"

	redis "gopkg.in/redis.v5"
)

//Fake ...
type Fake struct {
	NewClientStub        func(*redis.Options) Client
	newClientMutex       sync.RWMutex
	newClientArgsForCall []struct {
		arg1 *redis.Options
	}
	newClientReturns struct {
		result1 Client
	}
	NewScriptStub        func(src string) Script
	newScriptMutex       sync.RWMutex
	newScriptArgsForCall []struct {
		src string
	}
	newScriptReturns struct {
		result1 Script
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

//NewFake is the preferred way to initialise a Fake
func NewFake() *Fake {
	return new(Fake)
}

//NewClient ...
func (fake *Fake) NewClient(arg1 *redis.Options) Client {
	fake.newClientMutex.Lock()
	fake.newClientArgsForCall = append(fake.newClientArgsForCall, struct {
		arg1 *redis.Options
	}{arg1})
	fake.recordInvocation("NewClient", []interface{}{arg1})
	fake.newClientMutex.Unlock()
	if fake.NewClientStub != nil {
		return fake.NewClientStub(arg1)
	}
	return fake.newClientReturns.result1
}

//NewClientCallCount ...
func (fake *Fake) NewClientCallCount() int {
	fake.newClientMutex.RLock()
	defer fake.newClientMutex.RUnlock()
	return len(fake.newClientArgsForCall)
}

//NewClientArgsForCall ...
func (fake *Fake) NewClientArgsForCall(i int) *redis.Options {
	fake.newClientMutex.RLock()
	defer fake.newClientMutex.RUnlock()
	return fake.newClientArgsForCall[i].arg1
}

//NewClientReturns ...
func (fake *Fake) NewClientReturns(result1 Client) {
	fake.NewClientStub = nil
	fake.newClientReturns = struct {
		result1 Client
	}{result1}
}

//NewScript ...
func (fake *Fake) NewScript(src string) Script {
	fake.newScriptMutex.Lock()
	fake.newScriptArgsForCall = append(fake.newScriptArgsForCall, struct {
		src string
	}{src})
	fake.recordInvocation("NewScript", []interface{}{src})
	fake.newScriptMutex.Unlock()
	if fake.NewScriptStub != nil {
		return fake.NewScriptStub(src)
	}
	return fake.newScriptReturns.result1
}

//NewScriptCallCount ...
func (fake *Fake) NewScriptCallCount() int {
	fake.newScriptMutex.RLock()
	defer fake.newScriptMutex.RUnlock()
	return len(fake.newScriptArgsForCall)
}

//NewScriptArgsForCall ...
func (fake *Fake) NewScriptArgsForCall(i int) string {
	fake.newScriptMutex.RLock()
	defer fake.newScriptMutex.RUnlock()
	return fake.newScriptArgsForCall[i].src
}

//NewScriptReturns ...
func (fake *Fake) NewScriptReturns(result1 Script) {
	fake.NewScriptStub = nil
	fake.newScriptReturns = struct {
		result1 Script
	}{result1}
}

//Invocations ...
func (fake *Fake) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.newClientMutex.RLock()
	defer fake.newClientMutex.RUnlock()
	fake.newScriptMutex.RLock()
	defer fake.newScriptMutex.RUnlock()
	return fake.invocations
}

func (fake *Fake) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ Redis = new(Fake)
