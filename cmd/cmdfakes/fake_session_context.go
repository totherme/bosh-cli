// Code generated by counterfeiter. DO NOT EDIT.
package cmdfakes

import (
	"sync"

	"github.com/cloudfoundry/bosh-cli/v6/cmd"
	"github.com/cloudfoundry/bosh-cli/v6/cmd/config"
)

type FakeSessionContext struct {
	CACertStub        func() string
	cACertMutex       sync.RWMutex
	cACertArgsForCall []struct {
	}
	cACertReturns struct {
		result1 string
	}
	cACertReturnsOnCall map[int]struct {
		result1 string
	}
	ConfigStub        func() config.Config
	configMutex       sync.RWMutex
	configArgsForCall []struct {
	}
	configReturns struct {
		result1 config.Config
	}
	configReturnsOnCall map[int]struct {
		result1 config.Config
	}
	CredentialsStub        func() config.Creds
	credentialsMutex       sync.RWMutex
	credentialsArgsForCall []struct {
	}
	credentialsReturns struct {
		result1 config.Creds
	}
	credentialsReturnsOnCall map[int]struct {
		result1 config.Creds
	}
	DeploymentStub        func() string
	deploymentMutex       sync.RWMutex
	deploymentArgsForCall []struct {
	}
	deploymentReturns struct {
		result1 string
	}
	deploymentReturnsOnCall map[int]struct {
		result1 string
	}
	EnvironmentStub        func() string
	environmentMutex       sync.RWMutex
	environmentArgsForCall []struct {
	}
	environmentReturns struct {
		result1 string
	}
	environmentReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSessionContext) CACert() string {
	fake.cACertMutex.Lock()
	ret, specificReturn := fake.cACertReturnsOnCall[len(fake.cACertArgsForCall)]
	fake.cACertArgsForCall = append(fake.cACertArgsForCall, struct {
	}{})
	stub := fake.CACertStub
	fakeReturns := fake.cACertReturns
	fake.recordInvocation("CACert", []interface{}{})
	fake.cACertMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeSessionContext) CACertCallCount() int {
	fake.cACertMutex.RLock()
	defer fake.cACertMutex.RUnlock()
	return len(fake.cACertArgsForCall)
}

func (fake *FakeSessionContext) CACertCalls(stub func() string) {
	fake.cACertMutex.Lock()
	defer fake.cACertMutex.Unlock()
	fake.CACertStub = stub
}

func (fake *FakeSessionContext) CACertReturns(result1 string) {
	fake.cACertMutex.Lock()
	defer fake.cACertMutex.Unlock()
	fake.CACertStub = nil
	fake.cACertReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeSessionContext) CACertReturnsOnCall(i int, result1 string) {
	fake.cACertMutex.Lock()
	defer fake.cACertMutex.Unlock()
	fake.CACertStub = nil
	if fake.cACertReturnsOnCall == nil {
		fake.cACertReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.cACertReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeSessionContext) Config() config.Config {
	fake.configMutex.Lock()
	ret, specificReturn := fake.configReturnsOnCall[len(fake.configArgsForCall)]
	fake.configArgsForCall = append(fake.configArgsForCall, struct {
	}{})
	stub := fake.ConfigStub
	fakeReturns := fake.configReturns
	fake.recordInvocation("Config", []interface{}{})
	fake.configMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeSessionContext) ConfigCallCount() int {
	fake.configMutex.RLock()
	defer fake.configMutex.RUnlock()
	return len(fake.configArgsForCall)
}

func (fake *FakeSessionContext) ConfigCalls(stub func() config.Config) {
	fake.configMutex.Lock()
	defer fake.configMutex.Unlock()
	fake.ConfigStub = stub
}

func (fake *FakeSessionContext) ConfigReturns(result1 config.Config) {
	fake.configMutex.Lock()
	defer fake.configMutex.Unlock()
	fake.ConfigStub = nil
	fake.configReturns = struct {
		result1 config.Config
	}{result1}
}

func (fake *FakeSessionContext) ConfigReturnsOnCall(i int, result1 config.Config) {
	fake.configMutex.Lock()
	defer fake.configMutex.Unlock()
	fake.ConfigStub = nil
	if fake.configReturnsOnCall == nil {
		fake.configReturnsOnCall = make(map[int]struct {
			result1 config.Config
		})
	}
	fake.configReturnsOnCall[i] = struct {
		result1 config.Config
	}{result1}
}

func (fake *FakeSessionContext) Credentials() config.Creds {
	fake.credentialsMutex.Lock()
	ret, specificReturn := fake.credentialsReturnsOnCall[len(fake.credentialsArgsForCall)]
	fake.credentialsArgsForCall = append(fake.credentialsArgsForCall, struct {
	}{})
	stub := fake.CredentialsStub
	fakeReturns := fake.credentialsReturns
	fake.recordInvocation("Credentials", []interface{}{})
	fake.credentialsMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeSessionContext) CredentialsCallCount() int {
	fake.credentialsMutex.RLock()
	defer fake.credentialsMutex.RUnlock()
	return len(fake.credentialsArgsForCall)
}

func (fake *FakeSessionContext) CredentialsCalls(stub func() config.Creds) {
	fake.credentialsMutex.Lock()
	defer fake.credentialsMutex.Unlock()
	fake.CredentialsStub = stub
}

func (fake *FakeSessionContext) CredentialsReturns(result1 config.Creds) {
	fake.credentialsMutex.Lock()
	defer fake.credentialsMutex.Unlock()
	fake.CredentialsStub = nil
	fake.credentialsReturns = struct {
		result1 config.Creds
	}{result1}
}

func (fake *FakeSessionContext) CredentialsReturnsOnCall(i int, result1 config.Creds) {
	fake.credentialsMutex.Lock()
	defer fake.credentialsMutex.Unlock()
	fake.CredentialsStub = nil
	if fake.credentialsReturnsOnCall == nil {
		fake.credentialsReturnsOnCall = make(map[int]struct {
			result1 config.Creds
		})
	}
	fake.credentialsReturnsOnCall[i] = struct {
		result1 config.Creds
	}{result1}
}

func (fake *FakeSessionContext) Deployment() string {
	fake.deploymentMutex.Lock()
	ret, specificReturn := fake.deploymentReturnsOnCall[len(fake.deploymentArgsForCall)]
	fake.deploymentArgsForCall = append(fake.deploymentArgsForCall, struct {
	}{})
	stub := fake.DeploymentStub
	fakeReturns := fake.deploymentReturns
	fake.recordInvocation("Deployment", []interface{}{})
	fake.deploymentMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeSessionContext) DeploymentCallCount() int {
	fake.deploymentMutex.RLock()
	defer fake.deploymentMutex.RUnlock()
	return len(fake.deploymentArgsForCall)
}

func (fake *FakeSessionContext) DeploymentCalls(stub func() string) {
	fake.deploymentMutex.Lock()
	defer fake.deploymentMutex.Unlock()
	fake.DeploymentStub = stub
}

func (fake *FakeSessionContext) DeploymentReturns(result1 string) {
	fake.deploymentMutex.Lock()
	defer fake.deploymentMutex.Unlock()
	fake.DeploymentStub = nil
	fake.deploymentReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeSessionContext) DeploymentReturnsOnCall(i int, result1 string) {
	fake.deploymentMutex.Lock()
	defer fake.deploymentMutex.Unlock()
	fake.DeploymentStub = nil
	if fake.deploymentReturnsOnCall == nil {
		fake.deploymentReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.deploymentReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeSessionContext) Environment() string {
	fake.environmentMutex.Lock()
	ret, specificReturn := fake.environmentReturnsOnCall[len(fake.environmentArgsForCall)]
	fake.environmentArgsForCall = append(fake.environmentArgsForCall, struct {
	}{})
	stub := fake.EnvironmentStub
	fakeReturns := fake.environmentReturns
	fake.recordInvocation("Environment", []interface{}{})
	fake.environmentMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeSessionContext) EnvironmentCallCount() int {
	fake.environmentMutex.RLock()
	defer fake.environmentMutex.RUnlock()
	return len(fake.environmentArgsForCall)
}

func (fake *FakeSessionContext) EnvironmentCalls(stub func() string) {
	fake.environmentMutex.Lock()
	defer fake.environmentMutex.Unlock()
	fake.EnvironmentStub = stub
}

func (fake *FakeSessionContext) EnvironmentReturns(result1 string) {
	fake.environmentMutex.Lock()
	defer fake.environmentMutex.Unlock()
	fake.EnvironmentStub = nil
	fake.environmentReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeSessionContext) EnvironmentReturnsOnCall(i int, result1 string) {
	fake.environmentMutex.Lock()
	defer fake.environmentMutex.Unlock()
	fake.EnvironmentStub = nil
	if fake.environmentReturnsOnCall == nil {
		fake.environmentReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.environmentReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeSessionContext) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.cACertMutex.RLock()
	defer fake.cACertMutex.RUnlock()
	fake.configMutex.RLock()
	defer fake.configMutex.RUnlock()
	fake.credentialsMutex.RLock()
	defer fake.credentialsMutex.RUnlock()
	fake.deploymentMutex.RLock()
	defer fake.deploymentMutex.RUnlock()
	fake.environmentMutex.RLock()
	defer fake.environmentMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSessionContext) recordInvocation(key string, args []interface{}) {
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

var _ cmd.SessionContext = new(FakeSessionContext)
