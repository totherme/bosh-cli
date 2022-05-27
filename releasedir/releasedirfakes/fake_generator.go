// Code generated by counterfeiter. DO NOT EDIT.
package releasedirfakes

import (
	"sync"

	"github.com/cloudfoundry/bosh-cli/v6/releasedir"
)

type FakeGenerator struct {
	GenerateJobStub        func(string) error
	generateJobMutex       sync.RWMutex
	generateJobArgsForCall []struct {
		arg1 string
	}
	generateJobReturns struct {
		result1 error
	}
	generateJobReturnsOnCall map[int]struct {
		result1 error
	}
	GeneratePackageStub        func(string) error
	generatePackageMutex       sync.RWMutex
	generatePackageArgsForCall []struct {
		arg1 string
	}
	generatePackageReturns struct {
		result1 error
	}
	generatePackageReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeGenerator) GenerateJob(arg1 string) error {
	fake.generateJobMutex.Lock()
	ret, specificReturn := fake.generateJobReturnsOnCall[len(fake.generateJobArgsForCall)]
	fake.generateJobArgsForCall = append(fake.generateJobArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GenerateJob", []interface{}{arg1})
	fake.generateJobMutex.Unlock()
	if fake.GenerateJobStub != nil {
		return fake.GenerateJobStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.generateJobReturns
	return fakeReturns.result1
}

func (fake *FakeGenerator) GenerateJobCallCount() int {
	fake.generateJobMutex.RLock()
	defer fake.generateJobMutex.RUnlock()
	return len(fake.generateJobArgsForCall)
}

func (fake *FakeGenerator) GenerateJobCalls(stub func(string) error) {
	fake.generateJobMutex.Lock()
	defer fake.generateJobMutex.Unlock()
	fake.GenerateJobStub = stub
}

func (fake *FakeGenerator) GenerateJobArgsForCall(i int) string {
	fake.generateJobMutex.RLock()
	defer fake.generateJobMutex.RUnlock()
	argsForCall := fake.generateJobArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeGenerator) GenerateJobReturns(result1 error) {
	fake.generateJobMutex.Lock()
	defer fake.generateJobMutex.Unlock()
	fake.GenerateJobStub = nil
	fake.generateJobReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeGenerator) GenerateJobReturnsOnCall(i int, result1 error) {
	fake.generateJobMutex.Lock()
	defer fake.generateJobMutex.Unlock()
	fake.GenerateJobStub = nil
	if fake.generateJobReturnsOnCall == nil {
		fake.generateJobReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.generateJobReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeGenerator) GeneratePackage(arg1 string) error {
	fake.generatePackageMutex.Lock()
	ret, specificReturn := fake.generatePackageReturnsOnCall[len(fake.generatePackageArgsForCall)]
	fake.generatePackageArgsForCall = append(fake.generatePackageArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GeneratePackage", []interface{}{arg1})
	fake.generatePackageMutex.Unlock()
	if fake.GeneratePackageStub != nil {
		return fake.GeneratePackageStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.generatePackageReturns
	return fakeReturns.result1
}

func (fake *FakeGenerator) GeneratePackageCallCount() int {
	fake.generatePackageMutex.RLock()
	defer fake.generatePackageMutex.RUnlock()
	return len(fake.generatePackageArgsForCall)
}

func (fake *FakeGenerator) GeneratePackageCalls(stub func(string) error) {
	fake.generatePackageMutex.Lock()
	defer fake.generatePackageMutex.Unlock()
	fake.GeneratePackageStub = stub
}

func (fake *FakeGenerator) GeneratePackageArgsForCall(i int) string {
	fake.generatePackageMutex.RLock()
	defer fake.generatePackageMutex.RUnlock()
	argsForCall := fake.generatePackageArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeGenerator) GeneratePackageReturns(result1 error) {
	fake.generatePackageMutex.Lock()
	defer fake.generatePackageMutex.Unlock()
	fake.GeneratePackageStub = nil
	fake.generatePackageReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeGenerator) GeneratePackageReturnsOnCall(i int, result1 error) {
	fake.generatePackageMutex.Lock()
	defer fake.generatePackageMutex.Unlock()
	fake.GeneratePackageStub = nil
	if fake.generatePackageReturnsOnCall == nil {
		fake.generatePackageReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.generatePackageReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeGenerator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.generateJobMutex.RLock()
	defer fake.generateJobMutex.RUnlock()
	fake.generatePackageMutex.RLock()
	defer fake.generatePackageMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeGenerator) recordInvocation(key string, args []interface{}) {
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

var _ releasedir.Generator = new(FakeGenerator)
