// Code generated by counterfeiter. DO NOT EDIT.
package releasedirfakes

import (
	"sync"

	"github.com/cloudfoundry/bosh-cli/releasedir"
)

type FakeGitRepo struct {
	InitStub        func() error
	initMutex       sync.RWMutex
	initArgsForCall []struct {
	}
	initReturns struct {
		result1 error
	}
	initReturnsOnCall map[int]struct {
		result1 error
	}
	LastCommitSHAStub        func() (string, error)
	lastCommitSHAMutex       sync.RWMutex
	lastCommitSHAArgsForCall []struct {
	}
	lastCommitSHAReturns struct {
		result1 string
		result2 error
	}
	lastCommitSHAReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	MustNotBeDirtyStub        func(bool) (bool, error)
	mustNotBeDirtyMutex       sync.RWMutex
	mustNotBeDirtyArgsForCall []struct {
		arg1 bool
	}
	mustNotBeDirtyReturns struct {
		result1 bool
		result2 error
	}
	mustNotBeDirtyReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeGitRepo) Init() error {
	fake.initMutex.Lock()
	ret, specificReturn := fake.initReturnsOnCall[len(fake.initArgsForCall)]
	fake.initArgsForCall = append(fake.initArgsForCall, struct {
	}{})
	fake.recordInvocation("Init", []interface{}{})
	fake.initMutex.Unlock()
	if fake.InitStub != nil {
		return fake.InitStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.initReturns
	return fakeReturns.result1
}

func (fake *FakeGitRepo) InitCallCount() int {
	fake.initMutex.RLock()
	defer fake.initMutex.RUnlock()
	return len(fake.initArgsForCall)
}

func (fake *FakeGitRepo) InitCalls(stub func() error) {
	fake.initMutex.Lock()
	defer fake.initMutex.Unlock()
	fake.InitStub = stub
}

func (fake *FakeGitRepo) InitReturns(result1 error) {
	fake.initMutex.Lock()
	defer fake.initMutex.Unlock()
	fake.InitStub = nil
	fake.initReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeGitRepo) InitReturnsOnCall(i int, result1 error) {
	fake.initMutex.Lock()
	defer fake.initMutex.Unlock()
	fake.InitStub = nil
	if fake.initReturnsOnCall == nil {
		fake.initReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.initReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeGitRepo) LastCommitSHA() (string, error) {
	fake.lastCommitSHAMutex.Lock()
	ret, specificReturn := fake.lastCommitSHAReturnsOnCall[len(fake.lastCommitSHAArgsForCall)]
	fake.lastCommitSHAArgsForCall = append(fake.lastCommitSHAArgsForCall, struct {
	}{})
	fake.recordInvocation("LastCommitSHA", []interface{}{})
	fake.lastCommitSHAMutex.Unlock()
	if fake.LastCommitSHAStub != nil {
		return fake.LastCommitSHAStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.lastCommitSHAReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeGitRepo) LastCommitSHACallCount() int {
	fake.lastCommitSHAMutex.RLock()
	defer fake.lastCommitSHAMutex.RUnlock()
	return len(fake.lastCommitSHAArgsForCall)
}

func (fake *FakeGitRepo) LastCommitSHACalls(stub func() (string, error)) {
	fake.lastCommitSHAMutex.Lock()
	defer fake.lastCommitSHAMutex.Unlock()
	fake.LastCommitSHAStub = stub
}

func (fake *FakeGitRepo) LastCommitSHAReturns(result1 string, result2 error) {
	fake.lastCommitSHAMutex.Lock()
	defer fake.lastCommitSHAMutex.Unlock()
	fake.LastCommitSHAStub = nil
	fake.lastCommitSHAReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeGitRepo) LastCommitSHAReturnsOnCall(i int, result1 string, result2 error) {
	fake.lastCommitSHAMutex.Lock()
	defer fake.lastCommitSHAMutex.Unlock()
	fake.LastCommitSHAStub = nil
	if fake.lastCommitSHAReturnsOnCall == nil {
		fake.lastCommitSHAReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.lastCommitSHAReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeGitRepo) MustNotBeDirty(arg1 bool) (bool, error) {
	fake.mustNotBeDirtyMutex.Lock()
	ret, specificReturn := fake.mustNotBeDirtyReturnsOnCall[len(fake.mustNotBeDirtyArgsForCall)]
	fake.mustNotBeDirtyArgsForCall = append(fake.mustNotBeDirtyArgsForCall, struct {
		arg1 bool
	}{arg1})
	fake.recordInvocation("MustNotBeDirty", []interface{}{arg1})
	fake.mustNotBeDirtyMutex.Unlock()
	if fake.MustNotBeDirtyStub != nil {
		return fake.MustNotBeDirtyStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.mustNotBeDirtyReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeGitRepo) MustNotBeDirtyCallCount() int {
	fake.mustNotBeDirtyMutex.RLock()
	defer fake.mustNotBeDirtyMutex.RUnlock()
	return len(fake.mustNotBeDirtyArgsForCall)
}

func (fake *FakeGitRepo) MustNotBeDirtyCalls(stub func(bool) (bool, error)) {
	fake.mustNotBeDirtyMutex.Lock()
	defer fake.mustNotBeDirtyMutex.Unlock()
	fake.MustNotBeDirtyStub = stub
}

func (fake *FakeGitRepo) MustNotBeDirtyArgsForCall(i int) bool {
	fake.mustNotBeDirtyMutex.RLock()
	defer fake.mustNotBeDirtyMutex.RUnlock()
	argsForCall := fake.mustNotBeDirtyArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeGitRepo) MustNotBeDirtyReturns(result1 bool, result2 error) {
	fake.mustNotBeDirtyMutex.Lock()
	defer fake.mustNotBeDirtyMutex.Unlock()
	fake.MustNotBeDirtyStub = nil
	fake.mustNotBeDirtyReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeGitRepo) MustNotBeDirtyReturnsOnCall(i int, result1 bool, result2 error) {
	fake.mustNotBeDirtyMutex.Lock()
	defer fake.mustNotBeDirtyMutex.Unlock()
	fake.MustNotBeDirtyStub = nil
	if fake.mustNotBeDirtyReturnsOnCall == nil {
		fake.mustNotBeDirtyReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.mustNotBeDirtyReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeGitRepo) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.initMutex.RLock()
	defer fake.initMutex.RUnlock()
	fake.lastCommitSHAMutex.RLock()
	defer fake.lastCommitSHAMutex.RUnlock()
	fake.mustNotBeDirtyMutex.RLock()
	defer fake.mustNotBeDirtyMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeGitRepo) recordInvocation(key string, args []interface{}) {
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

var _ releasedir.GitRepo = new(FakeGitRepo)
