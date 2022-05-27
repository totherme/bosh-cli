// Code generated by counterfeiter. DO NOT EDIT.
package licensefakes

import (
	"sync"

	"github.com/cloudfoundry/bosh-cli/v6/release/license"
)

type FakeDirReader struct {
	ReadStub        func(string) (*license.License, error)
	readMutex       sync.RWMutex
	readArgsForCall []struct {
		arg1 string
	}
	readReturns struct {
		result1 *license.License
		result2 error
	}
	readReturnsOnCall map[int]struct {
		result1 *license.License
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDirReader) Read(arg1 string) (*license.License, error) {
	fake.readMutex.Lock()
	ret, specificReturn := fake.readReturnsOnCall[len(fake.readArgsForCall)]
	fake.readArgsForCall = append(fake.readArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Read", []interface{}{arg1})
	fake.readMutex.Unlock()
	if fake.ReadStub != nil {
		return fake.ReadStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.readReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeDirReader) ReadCallCount() int {
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	return len(fake.readArgsForCall)
}

func (fake *FakeDirReader) ReadCalls(stub func(string) (*license.License, error)) {
	fake.readMutex.Lock()
	defer fake.readMutex.Unlock()
	fake.ReadStub = stub
}

func (fake *FakeDirReader) ReadArgsForCall(i int) string {
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	argsForCall := fake.readArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeDirReader) ReadReturns(result1 *license.License, result2 error) {
	fake.readMutex.Lock()
	defer fake.readMutex.Unlock()
	fake.ReadStub = nil
	fake.readReturns = struct {
		result1 *license.License
		result2 error
	}{result1, result2}
}

func (fake *FakeDirReader) ReadReturnsOnCall(i int, result1 *license.License, result2 error) {
	fake.readMutex.Lock()
	defer fake.readMutex.Unlock()
	fake.ReadStub = nil
	if fake.readReturnsOnCall == nil {
		fake.readReturnsOnCall = make(map[int]struct {
			result1 *license.License
			result2 error
		})
	}
	fake.readReturnsOnCall[i] = struct {
		result1 *license.License
		result2 error
	}{result1, result2}
}

func (fake *FakeDirReader) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDirReader) recordInvocation(key string, args []interface{}) {
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

var _ license.DirReader = new(FakeDirReader)
