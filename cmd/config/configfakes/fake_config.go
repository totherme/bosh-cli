// Code generated by counterfeiter. DO NOT EDIT.
package configfakes

import (
	"sync"

	"github.com/cloudfoundry/bosh-cli/v6/cmd/config"
	"github.com/cloudfoundry/bosh-cli/v6/uaa"
)

type FakeConfig struct {
	AliasEnvironmentStub        func(string, string, string) (config.Config, error)
	aliasEnvironmentMutex       sync.RWMutex
	aliasEnvironmentArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
	}
	aliasEnvironmentReturns struct {
		result1 config.Config
		result2 error
	}
	aliasEnvironmentReturnsOnCall map[int]struct {
		result1 config.Config
		result2 error
	}
	CACertStub        func(string) string
	cACertMutex       sync.RWMutex
	cACertArgsForCall []struct {
		arg1 string
	}
	cACertReturns struct {
		result1 string
	}
	cACertReturnsOnCall map[int]struct {
		result1 string
	}
	CredentialsStub        func(string) config.Creds
	credentialsMutex       sync.RWMutex
	credentialsArgsForCall []struct {
		arg1 string
	}
	credentialsReturns struct {
		result1 config.Creds
	}
	credentialsReturnsOnCall map[int]struct {
		result1 config.Creds
	}
	EnvironmentsStub        func() []config.Environment
	environmentsMutex       sync.RWMutex
	environmentsArgsForCall []struct {
	}
	environmentsReturns struct {
		result1 []config.Environment
	}
	environmentsReturnsOnCall map[int]struct {
		result1 []config.Environment
	}
	ResolveEnvironmentStub        func(string) string
	resolveEnvironmentMutex       sync.RWMutex
	resolveEnvironmentArgsForCall []struct {
		arg1 string
	}
	resolveEnvironmentReturns struct {
		result1 string
	}
	resolveEnvironmentReturnsOnCall map[int]struct {
		result1 string
	}
	SaveStub        func() error
	saveMutex       sync.RWMutex
	saveArgsForCall []struct {
	}
	saveReturns struct {
		result1 error
	}
	saveReturnsOnCall map[int]struct {
		result1 error
	}
	SetCredentialsStub        func(string, config.Creds) config.Config
	setCredentialsMutex       sync.RWMutex
	setCredentialsArgsForCall []struct {
		arg1 string
		arg2 config.Creds
	}
	setCredentialsReturns struct {
		result1 config.Config
	}
	setCredentialsReturnsOnCall map[int]struct {
		result1 config.Config
	}
	UnaliasEnvironmentStub        func(string) (config.Config, error)
	unaliasEnvironmentMutex       sync.RWMutex
	unaliasEnvironmentArgsForCall []struct {
		arg1 string
	}
	unaliasEnvironmentReturns struct {
		result1 config.Config
		result2 error
	}
	unaliasEnvironmentReturnsOnCall map[int]struct {
		result1 config.Config
		result2 error
	}
	UnsetCredentialsStub        func(string) config.Config
	unsetCredentialsMutex       sync.RWMutex
	unsetCredentialsArgsForCall []struct {
		arg1 string
	}
	unsetCredentialsReturns struct {
		result1 config.Config
	}
	unsetCredentialsReturnsOnCall map[int]struct {
		result1 config.Config
	}
	UpdateConfigWithTokenStub        func(string, uaa.AccessToken) error
	updateConfigWithTokenMutex       sync.RWMutex
	updateConfigWithTokenArgsForCall []struct {
		arg1 string
		arg2 uaa.AccessToken
	}
	updateConfigWithTokenReturns struct {
		result1 error
	}
	updateConfigWithTokenReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeConfig) AliasEnvironment(arg1 string, arg2 string, arg3 string) (config.Config, error) {
	fake.aliasEnvironmentMutex.Lock()
	ret, specificReturn := fake.aliasEnvironmentReturnsOnCall[len(fake.aliasEnvironmentArgsForCall)]
	fake.aliasEnvironmentArgsForCall = append(fake.aliasEnvironmentArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.AliasEnvironmentStub
	fakeReturns := fake.aliasEnvironmentReturns
	fake.recordInvocation("AliasEnvironment", []interface{}{arg1, arg2, arg3})
	fake.aliasEnvironmentMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeConfig) AliasEnvironmentCallCount() int {
	fake.aliasEnvironmentMutex.RLock()
	defer fake.aliasEnvironmentMutex.RUnlock()
	return len(fake.aliasEnvironmentArgsForCall)
}

func (fake *FakeConfig) AliasEnvironmentCalls(stub func(string, string, string) (config.Config, error)) {
	fake.aliasEnvironmentMutex.Lock()
	defer fake.aliasEnvironmentMutex.Unlock()
	fake.AliasEnvironmentStub = stub
}

func (fake *FakeConfig) AliasEnvironmentArgsForCall(i int) (string, string, string) {
	fake.aliasEnvironmentMutex.RLock()
	defer fake.aliasEnvironmentMutex.RUnlock()
	argsForCall := fake.aliasEnvironmentArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeConfig) AliasEnvironmentReturns(result1 config.Config, result2 error) {
	fake.aliasEnvironmentMutex.Lock()
	defer fake.aliasEnvironmentMutex.Unlock()
	fake.AliasEnvironmentStub = nil
	fake.aliasEnvironmentReturns = struct {
		result1 config.Config
		result2 error
	}{result1, result2}
}

func (fake *FakeConfig) AliasEnvironmentReturnsOnCall(i int, result1 config.Config, result2 error) {
	fake.aliasEnvironmentMutex.Lock()
	defer fake.aliasEnvironmentMutex.Unlock()
	fake.AliasEnvironmentStub = nil
	if fake.aliasEnvironmentReturnsOnCall == nil {
		fake.aliasEnvironmentReturnsOnCall = make(map[int]struct {
			result1 config.Config
			result2 error
		})
	}
	fake.aliasEnvironmentReturnsOnCall[i] = struct {
		result1 config.Config
		result2 error
	}{result1, result2}
}

func (fake *FakeConfig) CACert(arg1 string) string {
	fake.cACertMutex.Lock()
	ret, specificReturn := fake.cACertReturnsOnCall[len(fake.cACertArgsForCall)]
	fake.cACertArgsForCall = append(fake.cACertArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.CACertStub
	fakeReturns := fake.cACertReturns
	fake.recordInvocation("CACert", []interface{}{arg1})
	fake.cACertMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeConfig) CACertCallCount() int {
	fake.cACertMutex.RLock()
	defer fake.cACertMutex.RUnlock()
	return len(fake.cACertArgsForCall)
}

func (fake *FakeConfig) CACertCalls(stub func(string) string) {
	fake.cACertMutex.Lock()
	defer fake.cACertMutex.Unlock()
	fake.CACertStub = stub
}

func (fake *FakeConfig) CACertArgsForCall(i int) string {
	fake.cACertMutex.RLock()
	defer fake.cACertMutex.RUnlock()
	argsForCall := fake.cACertArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeConfig) CACertReturns(result1 string) {
	fake.cACertMutex.Lock()
	defer fake.cACertMutex.Unlock()
	fake.CACertStub = nil
	fake.cACertReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeConfig) CACertReturnsOnCall(i int, result1 string) {
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

func (fake *FakeConfig) Credentials(arg1 string) config.Creds {
	fake.credentialsMutex.Lock()
	ret, specificReturn := fake.credentialsReturnsOnCall[len(fake.credentialsArgsForCall)]
	fake.credentialsArgsForCall = append(fake.credentialsArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.CredentialsStub
	fakeReturns := fake.credentialsReturns
	fake.recordInvocation("Credentials", []interface{}{arg1})
	fake.credentialsMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeConfig) CredentialsCallCount() int {
	fake.credentialsMutex.RLock()
	defer fake.credentialsMutex.RUnlock()
	return len(fake.credentialsArgsForCall)
}

func (fake *FakeConfig) CredentialsCalls(stub func(string) config.Creds) {
	fake.credentialsMutex.Lock()
	defer fake.credentialsMutex.Unlock()
	fake.CredentialsStub = stub
}

func (fake *FakeConfig) CredentialsArgsForCall(i int) string {
	fake.credentialsMutex.RLock()
	defer fake.credentialsMutex.RUnlock()
	argsForCall := fake.credentialsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeConfig) CredentialsReturns(result1 config.Creds) {
	fake.credentialsMutex.Lock()
	defer fake.credentialsMutex.Unlock()
	fake.CredentialsStub = nil
	fake.credentialsReturns = struct {
		result1 config.Creds
	}{result1}
}

func (fake *FakeConfig) CredentialsReturnsOnCall(i int, result1 config.Creds) {
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

func (fake *FakeConfig) Environments() []config.Environment {
	fake.environmentsMutex.Lock()
	ret, specificReturn := fake.environmentsReturnsOnCall[len(fake.environmentsArgsForCall)]
	fake.environmentsArgsForCall = append(fake.environmentsArgsForCall, struct {
	}{})
	stub := fake.EnvironmentsStub
	fakeReturns := fake.environmentsReturns
	fake.recordInvocation("Environments", []interface{}{})
	fake.environmentsMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeConfig) EnvironmentsCallCount() int {
	fake.environmentsMutex.RLock()
	defer fake.environmentsMutex.RUnlock()
	return len(fake.environmentsArgsForCall)
}

func (fake *FakeConfig) EnvironmentsCalls(stub func() []config.Environment) {
	fake.environmentsMutex.Lock()
	defer fake.environmentsMutex.Unlock()
	fake.EnvironmentsStub = stub
}

func (fake *FakeConfig) EnvironmentsReturns(result1 []config.Environment) {
	fake.environmentsMutex.Lock()
	defer fake.environmentsMutex.Unlock()
	fake.EnvironmentsStub = nil
	fake.environmentsReturns = struct {
		result1 []config.Environment
	}{result1}
}

func (fake *FakeConfig) EnvironmentsReturnsOnCall(i int, result1 []config.Environment) {
	fake.environmentsMutex.Lock()
	defer fake.environmentsMutex.Unlock()
	fake.EnvironmentsStub = nil
	if fake.environmentsReturnsOnCall == nil {
		fake.environmentsReturnsOnCall = make(map[int]struct {
			result1 []config.Environment
		})
	}
	fake.environmentsReturnsOnCall[i] = struct {
		result1 []config.Environment
	}{result1}
}

func (fake *FakeConfig) ResolveEnvironment(arg1 string) string {
	fake.resolveEnvironmentMutex.Lock()
	ret, specificReturn := fake.resolveEnvironmentReturnsOnCall[len(fake.resolveEnvironmentArgsForCall)]
	fake.resolveEnvironmentArgsForCall = append(fake.resolveEnvironmentArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.ResolveEnvironmentStub
	fakeReturns := fake.resolveEnvironmentReturns
	fake.recordInvocation("ResolveEnvironment", []interface{}{arg1})
	fake.resolveEnvironmentMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeConfig) ResolveEnvironmentCallCount() int {
	fake.resolveEnvironmentMutex.RLock()
	defer fake.resolveEnvironmentMutex.RUnlock()
	return len(fake.resolveEnvironmentArgsForCall)
}

func (fake *FakeConfig) ResolveEnvironmentCalls(stub func(string) string) {
	fake.resolveEnvironmentMutex.Lock()
	defer fake.resolveEnvironmentMutex.Unlock()
	fake.ResolveEnvironmentStub = stub
}

func (fake *FakeConfig) ResolveEnvironmentArgsForCall(i int) string {
	fake.resolveEnvironmentMutex.RLock()
	defer fake.resolveEnvironmentMutex.RUnlock()
	argsForCall := fake.resolveEnvironmentArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeConfig) ResolveEnvironmentReturns(result1 string) {
	fake.resolveEnvironmentMutex.Lock()
	defer fake.resolveEnvironmentMutex.Unlock()
	fake.ResolveEnvironmentStub = nil
	fake.resolveEnvironmentReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeConfig) ResolveEnvironmentReturnsOnCall(i int, result1 string) {
	fake.resolveEnvironmentMutex.Lock()
	defer fake.resolveEnvironmentMutex.Unlock()
	fake.ResolveEnvironmentStub = nil
	if fake.resolveEnvironmentReturnsOnCall == nil {
		fake.resolveEnvironmentReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.resolveEnvironmentReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeConfig) Save() error {
	fake.saveMutex.Lock()
	ret, specificReturn := fake.saveReturnsOnCall[len(fake.saveArgsForCall)]
	fake.saveArgsForCall = append(fake.saveArgsForCall, struct {
	}{})
	stub := fake.SaveStub
	fakeReturns := fake.saveReturns
	fake.recordInvocation("Save", []interface{}{})
	fake.saveMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeConfig) SaveCallCount() int {
	fake.saveMutex.RLock()
	defer fake.saveMutex.RUnlock()
	return len(fake.saveArgsForCall)
}

func (fake *FakeConfig) SaveCalls(stub func() error) {
	fake.saveMutex.Lock()
	defer fake.saveMutex.Unlock()
	fake.SaveStub = stub
}

func (fake *FakeConfig) SaveReturns(result1 error) {
	fake.saveMutex.Lock()
	defer fake.saveMutex.Unlock()
	fake.SaveStub = nil
	fake.saveReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeConfig) SaveReturnsOnCall(i int, result1 error) {
	fake.saveMutex.Lock()
	defer fake.saveMutex.Unlock()
	fake.SaveStub = nil
	if fake.saveReturnsOnCall == nil {
		fake.saveReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.saveReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeConfig) SetCredentials(arg1 string, arg2 config.Creds) config.Config {
	fake.setCredentialsMutex.Lock()
	ret, specificReturn := fake.setCredentialsReturnsOnCall[len(fake.setCredentialsArgsForCall)]
	fake.setCredentialsArgsForCall = append(fake.setCredentialsArgsForCall, struct {
		arg1 string
		arg2 config.Creds
	}{arg1, arg2})
	stub := fake.SetCredentialsStub
	fakeReturns := fake.setCredentialsReturns
	fake.recordInvocation("SetCredentials", []interface{}{arg1, arg2})
	fake.setCredentialsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeConfig) SetCredentialsCallCount() int {
	fake.setCredentialsMutex.RLock()
	defer fake.setCredentialsMutex.RUnlock()
	return len(fake.setCredentialsArgsForCall)
}

func (fake *FakeConfig) SetCredentialsCalls(stub func(string, config.Creds) config.Config) {
	fake.setCredentialsMutex.Lock()
	defer fake.setCredentialsMutex.Unlock()
	fake.SetCredentialsStub = stub
}

func (fake *FakeConfig) SetCredentialsArgsForCall(i int) (string, config.Creds) {
	fake.setCredentialsMutex.RLock()
	defer fake.setCredentialsMutex.RUnlock()
	argsForCall := fake.setCredentialsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeConfig) SetCredentialsReturns(result1 config.Config) {
	fake.setCredentialsMutex.Lock()
	defer fake.setCredentialsMutex.Unlock()
	fake.SetCredentialsStub = nil
	fake.setCredentialsReturns = struct {
		result1 config.Config
	}{result1}
}

func (fake *FakeConfig) SetCredentialsReturnsOnCall(i int, result1 config.Config) {
	fake.setCredentialsMutex.Lock()
	defer fake.setCredentialsMutex.Unlock()
	fake.SetCredentialsStub = nil
	if fake.setCredentialsReturnsOnCall == nil {
		fake.setCredentialsReturnsOnCall = make(map[int]struct {
			result1 config.Config
		})
	}
	fake.setCredentialsReturnsOnCall[i] = struct {
		result1 config.Config
	}{result1}
}

func (fake *FakeConfig) UnaliasEnvironment(arg1 string) (config.Config, error) {
	fake.unaliasEnvironmentMutex.Lock()
	ret, specificReturn := fake.unaliasEnvironmentReturnsOnCall[len(fake.unaliasEnvironmentArgsForCall)]
	fake.unaliasEnvironmentArgsForCall = append(fake.unaliasEnvironmentArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.UnaliasEnvironmentStub
	fakeReturns := fake.unaliasEnvironmentReturns
	fake.recordInvocation("UnaliasEnvironment", []interface{}{arg1})
	fake.unaliasEnvironmentMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeConfig) UnaliasEnvironmentCallCount() int {
	fake.unaliasEnvironmentMutex.RLock()
	defer fake.unaliasEnvironmentMutex.RUnlock()
	return len(fake.unaliasEnvironmentArgsForCall)
}

func (fake *FakeConfig) UnaliasEnvironmentCalls(stub func(string) (config.Config, error)) {
	fake.unaliasEnvironmentMutex.Lock()
	defer fake.unaliasEnvironmentMutex.Unlock()
	fake.UnaliasEnvironmentStub = stub
}

func (fake *FakeConfig) UnaliasEnvironmentArgsForCall(i int) string {
	fake.unaliasEnvironmentMutex.RLock()
	defer fake.unaliasEnvironmentMutex.RUnlock()
	argsForCall := fake.unaliasEnvironmentArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeConfig) UnaliasEnvironmentReturns(result1 config.Config, result2 error) {
	fake.unaliasEnvironmentMutex.Lock()
	defer fake.unaliasEnvironmentMutex.Unlock()
	fake.UnaliasEnvironmentStub = nil
	fake.unaliasEnvironmentReturns = struct {
		result1 config.Config
		result2 error
	}{result1, result2}
}

func (fake *FakeConfig) UnaliasEnvironmentReturnsOnCall(i int, result1 config.Config, result2 error) {
	fake.unaliasEnvironmentMutex.Lock()
	defer fake.unaliasEnvironmentMutex.Unlock()
	fake.UnaliasEnvironmentStub = nil
	if fake.unaliasEnvironmentReturnsOnCall == nil {
		fake.unaliasEnvironmentReturnsOnCall = make(map[int]struct {
			result1 config.Config
			result2 error
		})
	}
	fake.unaliasEnvironmentReturnsOnCall[i] = struct {
		result1 config.Config
		result2 error
	}{result1, result2}
}

func (fake *FakeConfig) UnsetCredentials(arg1 string) config.Config {
	fake.unsetCredentialsMutex.Lock()
	ret, specificReturn := fake.unsetCredentialsReturnsOnCall[len(fake.unsetCredentialsArgsForCall)]
	fake.unsetCredentialsArgsForCall = append(fake.unsetCredentialsArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.UnsetCredentialsStub
	fakeReturns := fake.unsetCredentialsReturns
	fake.recordInvocation("UnsetCredentials", []interface{}{arg1})
	fake.unsetCredentialsMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeConfig) UnsetCredentialsCallCount() int {
	fake.unsetCredentialsMutex.RLock()
	defer fake.unsetCredentialsMutex.RUnlock()
	return len(fake.unsetCredentialsArgsForCall)
}

func (fake *FakeConfig) UnsetCredentialsCalls(stub func(string) config.Config) {
	fake.unsetCredentialsMutex.Lock()
	defer fake.unsetCredentialsMutex.Unlock()
	fake.UnsetCredentialsStub = stub
}

func (fake *FakeConfig) UnsetCredentialsArgsForCall(i int) string {
	fake.unsetCredentialsMutex.RLock()
	defer fake.unsetCredentialsMutex.RUnlock()
	argsForCall := fake.unsetCredentialsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeConfig) UnsetCredentialsReturns(result1 config.Config) {
	fake.unsetCredentialsMutex.Lock()
	defer fake.unsetCredentialsMutex.Unlock()
	fake.UnsetCredentialsStub = nil
	fake.unsetCredentialsReturns = struct {
		result1 config.Config
	}{result1}
}

func (fake *FakeConfig) UnsetCredentialsReturnsOnCall(i int, result1 config.Config) {
	fake.unsetCredentialsMutex.Lock()
	defer fake.unsetCredentialsMutex.Unlock()
	fake.UnsetCredentialsStub = nil
	if fake.unsetCredentialsReturnsOnCall == nil {
		fake.unsetCredentialsReturnsOnCall = make(map[int]struct {
			result1 config.Config
		})
	}
	fake.unsetCredentialsReturnsOnCall[i] = struct {
		result1 config.Config
	}{result1}
}

func (fake *FakeConfig) UpdateConfigWithToken(arg1 string, arg2 uaa.AccessToken) error {
	fake.updateConfigWithTokenMutex.Lock()
	ret, specificReturn := fake.updateConfigWithTokenReturnsOnCall[len(fake.updateConfigWithTokenArgsForCall)]
	fake.updateConfigWithTokenArgsForCall = append(fake.updateConfigWithTokenArgsForCall, struct {
		arg1 string
		arg2 uaa.AccessToken
	}{arg1, arg2})
	stub := fake.UpdateConfigWithTokenStub
	fakeReturns := fake.updateConfigWithTokenReturns
	fake.recordInvocation("UpdateConfigWithToken", []interface{}{arg1, arg2})
	fake.updateConfigWithTokenMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeConfig) UpdateConfigWithTokenCallCount() int {
	fake.updateConfigWithTokenMutex.RLock()
	defer fake.updateConfigWithTokenMutex.RUnlock()
	return len(fake.updateConfigWithTokenArgsForCall)
}

func (fake *FakeConfig) UpdateConfigWithTokenCalls(stub func(string, uaa.AccessToken) error) {
	fake.updateConfigWithTokenMutex.Lock()
	defer fake.updateConfigWithTokenMutex.Unlock()
	fake.UpdateConfigWithTokenStub = stub
}

func (fake *FakeConfig) UpdateConfigWithTokenArgsForCall(i int) (string, uaa.AccessToken) {
	fake.updateConfigWithTokenMutex.RLock()
	defer fake.updateConfigWithTokenMutex.RUnlock()
	argsForCall := fake.updateConfigWithTokenArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeConfig) UpdateConfigWithTokenReturns(result1 error) {
	fake.updateConfigWithTokenMutex.Lock()
	defer fake.updateConfigWithTokenMutex.Unlock()
	fake.UpdateConfigWithTokenStub = nil
	fake.updateConfigWithTokenReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeConfig) UpdateConfigWithTokenReturnsOnCall(i int, result1 error) {
	fake.updateConfigWithTokenMutex.Lock()
	defer fake.updateConfigWithTokenMutex.Unlock()
	fake.UpdateConfigWithTokenStub = nil
	if fake.updateConfigWithTokenReturnsOnCall == nil {
		fake.updateConfigWithTokenReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateConfigWithTokenReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeConfig) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.aliasEnvironmentMutex.RLock()
	defer fake.aliasEnvironmentMutex.RUnlock()
	fake.cACertMutex.RLock()
	defer fake.cACertMutex.RUnlock()
	fake.credentialsMutex.RLock()
	defer fake.credentialsMutex.RUnlock()
	fake.environmentsMutex.RLock()
	defer fake.environmentsMutex.RUnlock()
	fake.resolveEnvironmentMutex.RLock()
	defer fake.resolveEnvironmentMutex.RUnlock()
	fake.saveMutex.RLock()
	defer fake.saveMutex.RUnlock()
	fake.setCredentialsMutex.RLock()
	defer fake.setCredentialsMutex.RUnlock()
	fake.unaliasEnvironmentMutex.RLock()
	defer fake.unaliasEnvironmentMutex.RUnlock()
	fake.unsetCredentialsMutex.RLock()
	defer fake.unsetCredentialsMutex.RUnlock()
	fake.updateConfigWithTokenMutex.RLock()
	defer fake.updateConfigWithTokenMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeConfig) recordInvocation(key string, args []interface{}) {
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

var _ config.Config = new(FakeConfig)
