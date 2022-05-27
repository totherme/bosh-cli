// Code generated by counterfeiter. DO NOT EDIT.
package releasefakes

import (
	"sync"

	"github.com/cloudfoundry/bosh-cli/v6/release"
	"github.com/cloudfoundry/bosh-cli/v6/release/job"
	"github.com/cloudfoundry/bosh-cli/v6/release/license"
	"github.com/cloudfoundry/bosh-cli/v6/release/manifest"
	"github.com/cloudfoundry/bosh-cli/v6/release/pkg"
)

type FakeRelease struct {
	BuildStub        func(release.ArchiveIndicies, release.ArchiveIndicies, int) error
	buildMutex       sync.RWMutex
	buildArgsForCall []struct {
		arg1 release.ArchiveIndicies
		arg2 release.ArchiveIndicies
		arg3 int
	}
	buildReturns struct {
		result1 error
	}
	buildReturnsOnCall map[int]struct {
		result1 error
	}
	CleanUpStub        func() error
	cleanUpMutex       sync.RWMutex
	cleanUpArgsForCall []struct {
	}
	cleanUpReturns struct {
		result1 error
	}
	cleanUpReturnsOnCall map[int]struct {
		result1 error
	}
	CommitHashWithMarkStub        func(string) string
	commitHashWithMarkMutex       sync.RWMutex
	commitHashWithMarkArgsForCall []struct {
		arg1 string
	}
	commitHashWithMarkReturns struct {
		result1 string
	}
	commitHashWithMarkReturnsOnCall map[int]struct {
		result1 string
	}
	CompiledPackagesStub        func() []*pkg.CompiledPackage
	compiledPackagesMutex       sync.RWMutex
	compiledPackagesArgsForCall []struct {
	}
	compiledPackagesReturns struct {
		result1 []*pkg.CompiledPackage
	}
	compiledPackagesReturnsOnCall map[int]struct {
		result1 []*pkg.CompiledPackage
	}
	CopyWithStub        func([]*job.Job, []*pkg.Package, *license.License, []*pkg.CompiledPackage) release.Release
	copyWithMutex       sync.RWMutex
	copyWithArgsForCall []struct {
		arg1 []*job.Job
		arg2 []*pkg.Package
		arg3 *license.License
		arg4 []*pkg.CompiledPackage
	}
	copyWithReturns struct {
		result1 release.Release
	}
	copyWithReturnsOnCall map[int]struct {
		result1 release.Release
	}
	FinalizeStub        func(release.ArchiveIndicies, int) error
	finalizeMutex       sync.RWMutex
	finalizeArgsForCall []struct {
		arg1 release.ArchiveIndicies
		arg2 int
	}
	finalizeReturns struct {
		result1 error
	}
	finalizeReturnsOnCall map[int]struct {
		result1 error
	}
	FindJobByNameStub        func(string) (job.Job, bool)
	findJobByNameMutex       sync.RWMutex
	findJobByNameArgsForCall []struct {
		arg1 string
	}
	findJobByNameReturns struct {
		result1 job.Job
		result2 bool
	}
	findJobByNameReturnsOnCall map[int]struct {
		result1 job.Job
		result2 bool
	}
	IsCompiledStub        func() bool
	isCompiledMutex       sync.RWMutex
	isCompiledArgsForCall []struct {
	}
	isCompiledReturns struct {
		result1 bool
	}
	isCompiledReturnsOnCall map[int]struct {
		result1 bool
	}
	JobsStub        func() []*job.Job
	jobsMutex       sync.RWMutex
	jobsArgsForCall []struct {
	}
	jobsReturns struct {
		result1 []*job.Job
	}
	jobsReturnsOnCall map[int]struct {
		result1 []*job.Job
	}
	LicenseStub        func() *license.License
	licenseMutex       sync.RWMutex
	licenseArgsForCall []struct {
	}
	licenseReturns struct {
		result1 *license.License
	}
	licenseReturnsOnCall map[int]struct {
		result1 *license.License
	}
	ManifestStub        func() manifest.Manifest
	manifestMutex       sync.RWMutex
	manifestArgsForCall []struct {
	}
	manifestReturns struct {
		result1 manifest.Manifest
	}
	manifestReturnsOnCall map[int]struct {
		result1 manifest.Manifest
	}
	NameStub        func() string
	nameMutex       sync.RWMutex
	nameArgsForCall []struct {
	}
	nameReturns struct {
		result1 string
	}
	nameReturnsOnCall map[int]struct {
		result1 string
	}
	PackagesStub        func() []*pkg.Package
	packagesMutex       sync.RWMutex
	packagesArgsForCall []struct {
	}
	packagesReturns struct {
		result1 []*pkg.Package
	}
	packagesReturnsOnCall map[int]struct {
		result1 []*pkg.Package
	}
	SetCommitHashStub        func(string)
	setCommitHashMutex       sync.RWMutex
	setCommitHashArgsForCall []struct {
		arg1 string
	}
	SetNameStub        func(string)
	setNameMutex       sync.RWMutex
	setNameArgsForCall []struct {
		arg1 string
	}
	SetUncommittedChangesStub        func(bool)
	setUncommittedChangesMutex       sync.RWMutex
	setUncommittedChangesArgsForCall []struct {
		arg1 bool
	}
	SetVersionStub        func(string)
	setVersionMutex       sync.RWMutex
	setVersionArgsForCall []struct {
		arg1 string
	}
	VersionStub        func() string
	versionMutex       sync.RWMutex
	versionArgsForCall []struct {
	}
	versionReturns struct {
		result1 string
	}
	versionReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRelease) Build(arg1 release.ArchiveIndicies, arg2 release.ArchiveIndicies, arg3 int) error {
	fake.buildMutex.Lock()
	ret, specificReturn := fake.buildReturnsOnCall[len(fake.buildArgsForCall)]
	fake.buildArgsForCall = append(fake.buildArgsForCall, struct {
		arg1 release.ArchiveIndicies
		arg2 release.ArchiveIndicies
		arg3 int
	}{arg1, arg2, arg3})
	fake.recordInvocation("Build", []interface{}{arg1, arg2, arg3})
	fake.buildMutex.Unlock()
	if fake.BuildStub != nil {
		return fake.BuildStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.buildReturns
	return fakeReturns.result1
}

func (fake *FakeRelease) BuildCallCount() int {
	fake.buildMutex.RLock()
	defer fake.buildMutex.RUnlock()
	return len(fake.buildArgsForCall)
}

func (fake *FakeRelease) BuildCalls(stub func(release.ArchiveIndicies, release.ArchiveIndicies, int) error) {
	fake.buildMutex.Lock()
	defer fake.buildMutex.Unlock()
	fake.BuildStub = stub
}

func (fake *FakeRelease) BuildArgsForCall(i int) (release.ArchiveIndicies, release.ArchiveIndicies, int) {
	fake.buildMutex.RLock()
	defer fake.buildMutex.RUnlock()
	argsForCall := fake.buildArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeRelease) BuildReturns(result1 error) {
	fake.buildMutex.Lock()
	defer fake.buildMutex.Unlock()
	fake.BuildStub = nil
	fake.buildReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRelease) BuildReturnsOnCall(i int, result1 error) {
	fake.buildMutex.Lock()
	defer fake.buildMutex.Unlock()
	fake.BuildStub = nil
	if fake.buildReturnsOnCall == nil {
		fake.buildReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.buildReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRelease) CleanUp() error {
	fake.cleanUpMutex.Lock()
	ret, specificReturn := fake.cleanUpReturnsOnCall[len(fake.cleanUpArgsForCall)]
	fake.cleanUpArgsForCall = append(fake.cleanUpArgsForCall, struct {
	}{})
	fake.recordInvocation("CleanUp", []interface{}{})
	fake.cleanUpMutex.Unlock()
	if fake.CleanUpStub != nil {
		return fake.CleanUpStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.cleanUpReturns
	return fakeReturns.result1
}

func (fake *FakeRelease) CleanUpCallCount() int {
	fake.cleanUpMutex.RLock()
	defer fake.cleanUpMutex.RUnlock()
	return len(fake.cleanUpArgsForCall)
}

func (fake *FakeRelease) CleanUpCalls(stub func() error) {
	fake.cleanUpMutex.Lock()
	defer fake.cleanUpMutex.Unlock()
	fake.CleanUpStub = stub
}

func (fake *FakeRelease) CleanUpReturns(result1 error) {
	fake.cleanUpMutex.Lock()
	defer fake.cleanUpMutex.Unlock()
	fake.CleanUpStub = nil
	fake.cleanUpReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRelease) CleanUpReturnsOnCall(i int, result1 error) {
	fake.cleanUpMutex.Lock()
	defer fake.cleanUpMutex.Unlock()
	fake.CleanUpStub = nil
	if fake.cleanUpReturnsOnCall == nil {
		fake.cleanUpReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.cleanUpReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRelease) CommitHashWithMark(arg1 string) string {
	fake.commitHashWithMarkMutex.Lock()
	ret, specificReturn := fake.commitHashWithMarkReturnsOnCall[len(fake.commitHashWithMarkArgsForCall)]
	fake.commitHashWithMarkArgsForCall = append(fake.commitHashWithMarkArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("CommitHashWithMark", []interface{}{arg1})
	fake.commitHashWithMarkMutex.Unlock()
	if fake.CommitHashWithMarkStub != nil {
		return fake.CommitHashWithMarkStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.commitHashWithMarkReturns
	return fakeReturns.result1
}

func (fake *FakeRelease) CommitHashWithMarkCallCount() int {
	fake.commitHashWithMarkMutex.RLock()
	defer fake.commitHashWithMarkMutex.RUnlock()
	return len(fake.commitHashWithMarkArgsForCall)
}

func (fake *FakeRelease) CommitHashWithMarkCalls(stub func(string) string) {
	fake.commitHashWithMarkMutex.Lock()
	defer fake.commitHashWithMarkMutex.Unlock()
	fake.CommitHashWithMarkStub = stub
}

func (fake *FakeRelease) CommitHashWithMarkArgsForCall(i int) string {
	fake.commitHashWithMarkMutex.RLock()
	defer fake.commitHashWithMarkMutex.RUnlock()
	argsForCall := fake.commitHashWithMarkArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRelease) CommitHashWithMarkReturns(result1 string) {
	fake.commitHashWithMarkMutex.Lock()
	defer fake.commitHashWithMarkMutex.Unlock()
	fake.CommitHashWithMarkStub = nil
	fake.commitHashWithMarkReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeRelease) CommitHashWithMarkReturnsOnCall(i int, result1 string) {
	fake.commitHashWithMarkMutex.Lock()
	defer fake.commitHashWithMarkMutex.Unlock()
	fake.CommitHashWithMarkStub = nil
	if fake.commitHashWithMarkReturnsOnCall == nil {
		fake.commitHashWithMarkReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.commitHashWithMarkReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeRelease) CompiledPackages() []*pkg.CompiledPackage {
	fake.compiledPackagesMutex.Lock()
	ret, specificReturn := fake.compiledPackagesReturnsOnCall[len(fake.compiledPackagesArgsForCall)]
	fake.compiledPackagesArgsForCall = append(fake.compiledPackagesArgsForCall, struct {
	}{})
	fake.recordInvocation("CompiledPackages", []interface{}{})
	fake.compiledPackagesMutex.Unlock()
	if fake.CompiledPackagesStub != nil {
		return fake.CompiledPackagesStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.compiledPackagesReturns
	return fakeReturns.result1
}

func (fake *FakeRelease) CompiledPackagesCallCount() int {
	fake.compiledPackagesMutex.RLock()
	defer fake.compiledPackagesMutex.RUnlock()
	return len(fake.compiledPackagesArgsForCall)
}

func (fake *FakeRelease) CompiledPackagesCalls(stub func() []*pkg.CompiledPackage) {
	fake.compiledPackagesMutex.Lock()
	defer fake.compiledPackagesMutex.Unlock()
	fake.CompiledPackagesStub = stub
}

func (fake *FakeRelease) CompiledPackagesReturns(result1 []*pkg.CompiledPackage) {
	fake.compiledPackagesMutex.Lock()
	defer fake.compiledPackagesMutex.Unlock()
	fake.CompiledPackagesStub = nil
	fake.compiledPackagesReturns = struct {
		result1 []*pkg.CompiledPackage
	}{result1}
}

func (fake *FakeRelease) CompiledPackagesReturnsOnCall(i int, result1 []*pkg.CompiledPackage) {
	fake.compiledPackagesMutex.Lock()
	defer fake.compiledPackagesMutex.Unlock()
	fake.CompiledPackagesStub = nil
	if fake.compiledPackagesReturnsOnCall == nil {
		fake.compiledPackagesReturnsOnCall = make(map[int]struct {
			result1 []*pkg.CompiledPackage
		})
	}
	fake.compiledPackagesReturnsOnCall[i] = struct {
		result1 []*pkg.CompiledPackage
	}{result1}
}

func (fake *FakeRelease) CopyWith(arg1 []*job.Job, arg2 []*pkg.Package, arg3 *license.License, arg4 []*pkg.CompiledPackage) release.Release {
	var arg1Copy []*job.Job
	if arg1 != nil {
		arg1Copy = make([]*job.Job, len(arg1))
		copy(arg1Copy, arg1)
	}
	var arg2Copy []*pkg.Package
	if arg2 != nil {
		arg2Copy = make([]*pkg.Package, len(arg2))
		copy(arg2Copy, arg2)
	}
	var arg4Copy []*pkg.CompiledPackage
	if arg4 != nil {
		arg4Copy = make([]*pkg.CompiledPackage, len(arg4))
		copy(arg4Copy, arg4)
	}
	fake.copyWithMutex.Lock()
	ret, specificReturn := fake.copyWithReturnsOnCall[len(fake.copyWithArgsForCall)]
	fake.copyWithArgsForCall = append(fake.copyWithArgsForCall, struct {
		arg1 []*job.Job
		arg2 []*pkg.Package
		arg3 *license.License
		arg4 []*pkg.CompiledPackage
	}{arg1Copy, arg2Copy, arg3, arg4Copy})
	fake.recordInvocation("CopyWith", []interface{}{arg1Copy, arg2Copy, arg3, arg4Copy})
	fake.copyWithMutex.Unlock()
	if fake.CopyWithStub != nil {
		return fake.CopyWithStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.copyWithReturns
	return fakeReturns.result1
}

func (fake *FakeRelease) CopyWithCallCount() int {
	fake.copyWithMutex.RLock()
	defer fake.copyWithMutex.RUnlock()
	return len(fake.copyWithArgsForCall)
}

func (fake *FakeRelease) CopyWithCalls(stub func([]*job.Job, []*pkg.Package, *license.License, []*pkg.CompiledPackage) release.Release) {
	fake.copyWithMutex.Lock()
	defer fake.copyWithMutex.Unlock()
	fake.CopyWithStub = stub
}

func (fake *FakeRelease) CopyWithArgsForCall(i int) ([]*job.Job, []*pkg.Package, *license.License, []*pkg.CompiledPackage) {
	fake.copyWithMutex.RLock()
	defer fake.copyWithMutex.RUnlock()
	argsForCall := fake.copyWithArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeRelease) CopyWithReturns(result1 release.Release) {
	fake.copyWithMutex.Lock()
	defer fake.copyWithMutex.Unlock()
	fake.CopyWithStub = nil
	fake.copyWithReturns = struct {
		result1 release.Release
	}{result1}
}

func (fake *FakeRelease) CopyWithReturnsOnCall(i int, result1 release.Release) {
	fake.copyWithMutex.Lock()
	defer fake.copyWithMutex.Unlock()
	fake.CopyWithStub = nil
	if fake.copyWithReturnsOnCall == nil {
		fake.copyWithReturnsOnCall = make(map[int]struct {
			result1 release.Release
		})
	}
	fake.copyWithReturnsOnCall[i] = struct {
		result1 release.Release
	}{result1}
}

func (fake *FakeRelease) Finalize(arg1 release.ArchiveIndicies, arg2 int) error {
	fake.finalizeMutex.Lock()
	ret, specificReturn := fake.finalizeReturnsOnCall[len(fake.finalizeArgsForCall)]
	fake.finalizeArgsForCall = append(fake.finalizeArgsForCall, struct {
		arg1 release.ArchiveIndicies
		arg2 int
	}{arg1, arg2})
	fake.recordInvocation("Finalize", []interface{}{arg1, arg2})
	fake.finalizeMutex.Unlock()
	if fake.FinalizeStub != nil {
		return fake.FinalizeStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.finalizeReturns
	return fakeReturns.result1
}

func (fake *FakeRelease) FinalizeCallCount() int {
	fake.finalizeMutex.RLock()
	defer fake.finalizeMutex.RUnlock()
	return len(fake.finalizeArgsForCall)
}

func (fake *FakeRelease) FinalizeCalls(stub func(release.ArchiveIndicies, int) error) {
	fake.finalizeMutex.Lock()
	defer fake.finalizeMutex.Unlock()
	fake.FinalizeStub = stub
}

func (fake *FakeRelease) FinalizeArgsForCall(i int) (release.ArchiveIndicies, int) {
	fake.finalizeMutex.RLock()
	defer fake.finalizeMutex.RUnlock()
	argsForCall := fake.finalizeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRelease) FinalizeReturns(result1 error) {
	fake.finalizeMutex.Lock()
	defer fake.finalizeMutex.Unlock()
	fake.FinalizeStub = nil
	fake.finalizeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRelease) FinalizeReturnsOnCall(i int, result1 error) {
	fake.finalizeMutex.Lock()
	defer fake.finalizeMutex.Unlock()
	fake.FinalizeStub = nil
	if fake.finalizeReturnsOnCall == nil {
		fake.finalizeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.finalizeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRelease) FindJobByName(arg1 string) (job.Job, bool) {
	fake.findJobByNameMutex.Lock()
	ret, specificReturn := fake.findJobByNameReturnsOnCall[len(fake.findJobByNameArgsForCall)]
	fake.findJobByNameArgsForCall = append(fake.findJobByNameArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("FindJobByName", []interface{}{arg1})
	fake.findJobByNameMutex.Unlock()
	if fake.FindJobByNameStub != nil {
		return fake.FindJobByNameStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.findJobByNameReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRelease) FindJobByNameCallCount() int {
	fake.findJobByNameMutex.RLock()
	defer fake.findJobByNameMutex.RUnlock()
	return len(fake.findJobByNameArgsForCall)
}

func (fake *FakeRelease) FindJobByNameCalls(stub func(string) (job.Job, bool)) {
	fake.findJobByNameMutex.Lock()
	defer fake.findJobByNameMutex.Unlock()
	fake.FindJobByNameStub = stub
}

func (fake *FakeRelease) FindJobByNameArgsForCall(i int) string {
	fake.findJobByNameMutex.RLock()
	defer fake.findJobByNameMutex.RUnlock()
	argsForCall := fake.findJobByNameArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRelease) FindJobByNameReturns(result1 job.Job, result2 bool) {
	fake.findJobByNameMutex.Lock()
	defer fake.findJobByNameMutex.Unlock()
	fake.FindJobByNameStub = nil
	fake.findJobByNameReturns = struct {
		result1 job.Job
		result2 bool
	}{result1, result2}
}

func (fake *FakeRelease) FindJobByNameReturnsOnCall(i int, result1 job.Job, result2 bool) {
	fake.findJobByNameMutex.Lock()
	defer fake.findJobByNameMutex.Unlock()
	fake.FindJobByNameStub = nil
	if fake.findJobByNameReturnsOnCall == nil {
		fake.findJobByNameReturnsOnCall = make(map[int]struct {
			result1 job.Job
			result2 bool
		})
	}
	fake.findJobByNameReturnsOnCall[i] = struct {
		result1 job.Job
		result2 bool
	}{result1, result2}
}

func (fake *FakeRelease) IsCompiled() bool {
	fake.isCompiledMutex.Lock()
	ret, specificReturn := fake.isCompiledReturnsOnCall[len(fake.isCompiledArgsForCall)]
	fake.isCompiledArgsForCall = append(fake.isCompiledArgsForCall, struct {
	}{})
	fake.recordInvocation("IsCompiled", []interface{}{})
	fake.isCompiledMutex.Unlock()
	if fake.IsCompiledStub != nil {
		return fake.IsCompiledStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.isCompiledReturns
	return fakeReturns.result1
}

func (fake *FakeRelease) IsCompiledCallCount() int {
	fake.isCompiledMutex.RLock()
	defer fake.isCompiledMutex.RUnlock()
	return len(fake.isCompiledArgsForCall)
}

func (fake *FakeRelease) IsCompiledCalls(stub func() bool) {
	fake.isCompiledMutex.Lock()
	defer fake.isCompiledMutex.Unlock()
	fake.IsCompiledStub = stub
}

func (fake *FakeRelease) IsCompiledReturns(result1 bool) {
	fake.isCompiledMutex.Lock()
	defer fake.isCompiledMutex.Unlock()
	fake.IsCompiledStub = nil
	fake.isCompiledReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeRelease) IsCompiledReturnsOnCall(i int, result1 bool) {
	fake.isCompiledMutex.Lock()
	defer fake.isCompiledMutex.Unlock()
	fake.IsCompiledStub = nil
	if fake.isCompiledReturnsOnCall == nil {
		fake.isCompiledReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.isCompiledReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeRelease) Jobs() []*job.Job {
	fake.jobsMutex.Lock()
	ret, specificReturn := fake.jobsReturnsOnCall[len(fake.jobsArgsForCall)]
	fake.jobsArgsForCall = append(fake.jobsArgsForCall, struct {
	}{})
	fake.recordInvocation("Jobs", []interface{}{})
	fake.jobsMutex.Unlock()
	if fake.JobsStub != nil {
		return fake.JobsStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.jobsReturns
	return fakeReturns.result1
}

func (fake *FakeRelease) JobsCallCount() int {
	fake.jobsMutex.RLock()
	defer fake.jobsMutex.RUnlock()
	return len(fake.jobsArgsForCall)
}

func (fake *FakeRelease) JobsCalls(stub func() []*job.Job) {
	fake.jobsMutex.Lock()
	defer fake.jobsMutex.Unlock()
	fake.JobsStub = stub
}

func (fake *FakeRelease) JobsReturns(result1 []*job.Job) {
	fake.jobsMutex.Lock()
	defer fake.jobsMutex.Unlock()
	fake.JobsStub = nil
	fake.jobsReturns = struct {
		result1 []*job.Job
	}{result1}
}

func (fake *FakeRelease) JobsReturnsOnCall(i int, result1 []*job.Job) {
	fake.jobsMutex.Lock()
	defer fake.jobsMutex.Unlock()
	fake.JobsStub = nil
	if fake.jobsReturnsOnCall == nil {
		fake.jobsReturnsOnCall = make(map[int]struct {
			result1 []*job.Job
		})
	}
	fake.jobsReturnsOnCall[i] = struct {
		result1 []*job.Job
	}{result1}
}

func (fake *FakeRelease) License() *license.License {
	fake.licenseMutex.Lock()
	ret, specificReturn := fake.licenseReturnsOnCall[len(fake.licenseArgsForCall)]
	fake.licenseArgsForCall = append(fake.licenseArgsForCall, struct {
	}{})
	fake.recordInvocation("License", []interface{}{})
	fake.licenseMutex.Unlock()
	if fake.LicenseStub != nil {
		return fake.LicenseStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.licenseReturns
	return fakeReturns.result1
}

func (fake *FakeRelease) LicenseCallCount() int {
	fake.licenseMutex.RLock()
	defer fake.licenseMutex.RUnlock()
	return len(fake.licenseArgsForCall)
}

func (fake *FakeRelease) LicenseCalls(stub func() *license.License) {
	fake.licenseMutex.Lock()
	defer fake.licenseMutex.Unlock()
	fake.LicenseStub = stub
}

func (fake *FakeRelease) LicenseReturns(result1 *license.License) {
	fake.licenseMutex.Lock()
	defer fake.licenseMutex.Unlock()
	fake.LicenseStub = nil
	fake.licenseReturns = struct {
		result1 *license.License
	}{result1}
}

func (fake *FakeRelease) LicenseReturnsOnCall(i int, result1 *license.License) {
	fake.licenseMutex.Lock()
	defer fake.licenseMutex.Unlock()
	fake.LicenseStub = nil
	if fake.licenseReturnsOnCall == nil {
		fake.licenseReturnsOnCall = make(map[int]struct {
			result1 *license.License
		})
	}
	fake.licenseReturnsOnCall[i] = struct {
		result1 *license.License
	}{result1}
}

func (fake *FakeRelease) Manifest() manifest.Manifest {
	fake.manifestMutex.Lock()
	ret, specificReturn := fake.manifestReturnsOnCall[len(fake.manifestArgsForCall)]
	fake.manifestArgsForCall = append(fake.manifestArgsForCall, struct {
	}{})
	fake.recordInvocation("Manifest", []interface{}{})
	fake.manifestMutex.Unlock()
	if fake.ManifestStub != nil {
		return fake.ManifestStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.manifestReturns
	return fakeReturns.result1
}

func (fake *FakeRelease) ManifestCallCount() int {
	fake.manifestMutex.RLock()
	defer fake.manifestMutex.RUnlock()
	return len(fake.manifestArgsForCall)
}

func (fake *FakeRelease) ManifestCalls(stub func() manifest.Manifest) {
	fake.manifestMutex.Lock()
	defer fake.manifestMutex.Unlock()
	fake.ManifestStub = stub
}

func (fake *FakeRelease) ManifestReturns(result1 manifest.Manifest) {
	fake.manifestMutex.Lock()
	defer fake.manifestMutex.Unlock()
	fake.ManifestStub = nil
	fake.manifestReturns = struct {
		result1 manifest.Manifest
	}{result1}
}

func (fake *FakeRelease) ManifestReturnsOnCall(i int, result1 manifest.Manifest) {
	fake.manifestMutex.Lock()
	defer fake.manifestMutex.Unlock()
	fake.ManifestStub = nil
	if fake.manifestReturnsOnCall == nil {
		fake.manifestReturnsOnCall = make(map[int]struct {
			result1 manifest.Manifest
		})
	}
	fake.manifestReturnsOnCall[i] = struct {
		result1 manifest.Manifest
	}{result1}
}

func (fake *FakeRelease) Name() string {
	fake.nameMutex.Lock()
	ret, specificReturn := fake.nameReturnsOnCall[len(fake.nameArgsForCall)]
	fake.nameArgsForCall = append(fake.nameArgsForCall, struct {
	}{})
	fake.recordInvocation("Name", []interface{}{})
	fake.nameMutex.Unlock()
	if fake.NameStub != nil {
		return fake.NameStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.nameReturns
	return fakeReturns.result1
}

func (fake *FakeRelease) NameCallCount() int {
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	return len(fake.nameArgsForCall)
}

func (fake *FakeRelease) NameCalls(stub func() string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = stub
}

func (fake *FakeRelease) NameReturns(result1 string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = nil
	fake.nameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeRelease) NameReturnsOnCall(i int, result1 string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = nil
	if fake.nameReturnsOnCall == nil {
		fake.nameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.nameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeRelease) Packages() []*pkg.Package {
	fake.packagesMutex.Lock()
	ret, specificReturn := fake.packagesReturnsOnCall[len(fake.packagesArgsForCall)]
	fake.packagesArgsForCall = append(fake.packagesArgsForCall, struct {
	}{})
	fake.recordInvocation("Packages", []interface{}{})
	fake.packagesMutex.Unlock()
	if fake.PackagesStub != nil {
		return fake.PackagesStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.packagesReturns
	return fakeReturns.result1
}

func (fake *FakeRelease) PackagesCallCount() int {
	fake.packagesMutex.RLock()
	defer fake.packagesMutex.RUnlock()
	return len(fake.packagesArgsForCall)
}

func (fake *FakeRelease) PackagesCalls(stub func() []*pkg.Package) {
	fake.packagesMutex.Lock()
	defer fake.packagesMutex.Unlock()
	fake.PackagesStub = stub
}

func (fake *FakeRelease) PackagesReturns(result1 []*pkg.Package) {
	fake.packagesMutex.Lock()
	defer fake.packagesMutex.Unlock()
	fake.PackagesStub = nil
	fake.packagesReturns = struct {
		result1 []*pkg.Package
	}{result1}
}

func (fake *FakeRelease) PackagesReturnsOnCall(i int, result1 []*pkg.Package) {
	fake.packagesMutex.Lock()
	defer fake.packagesMutex.Unlock()
	fake.PackagesStub = nil
	if fake.packagesReturnsOnCall == nil {
		fake.packagesReturnsOnCall = make(map[int]struct {
			result1 []*pkg.Package
		})
	}
	fake.packagesReturnsOnCall[i] = struct {
		result1 []*pkg.Package
	}{result1}
}

func (fake *FakeRelease) SetCommitHash(arg1 string) {
	fake.setCommitHashMutex.Lock()
	fake.setCommitHashArgsForCall = append(fake.setCommitHashArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("SetCommitHash", []interface{}{arg1})
	fake.setCommitHashMutex.Unlock()
	if fake.SetCommitHashStub != nil {
		fake.SetCommitHashStub(arg1)
	}
}

func (fake *FakeRelease) SetCommitHashCallCount() int {
	fake.setCommitHashMutex.RLock()
	defer fake.setCommitHashMutex.RUnlock()
	return len(fake.setCommitHashArgsForCall)
}

func (fake *FakeRelease) SetCommitHashCalls(stub func(string)) {
	fake.setCommitHashMutex.Lock()
	defer fake.setCommitHashMutex.Unlock()
	fake.SetCommitHashStub = stub
}

func (fake *FakeRelease) SetCommitHashArgsForCall(i int) string {
	fake.setCommitHashMutex.RLock()
	defer fake.setCommitHashMutex.RUnlock()
	argsForCall := fake.setCommitHashArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRelease) SetName(arg1 string) {
	fake.setNameMutex.Lock()
	fake.setNameArgsForCall = append(fake.setNameArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("SetName", []interface{}{arg1})
	fake.setNameMutex.Unlock()
	if fake.SetNameStub != nil {
		fake.SetNameStub(arg1)
	}
}

func (fake *FakeRelease) SetNameCallCount() int {
	fake.setNameMutex.RLock()
	defer fake.setNameMutex.RUnlock()
	return len(fake.setNameArgsForCall)
}

func (fake *FakeRelease) SetNameCalls(stub func(string)) {
	fake.setNameMutex.Lock()
	defer fake.setNameMutex.Unlock()
	fake.SetNameStub = stub
}

func (fake *FakeRelease) SetNameArgsForCall(i int) string {
	fake.setNameMutex.RLock()
	defer fake.setNameMutex.RUnlock()
	argsForCall := fake.setNameArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRelease) SetUncommittedChanges(arg1 bool) {
	fake.setUncommittedChangesMutex.Lock()
	fake.setUncommittedChangesArgsForCall = append(fake.setUncommittedChangesArgsForCall, struct {
		arg1 bool
	}{arg1})
	fake.recordInvocation("SetUncommittedChanges", []interface{}{arg1})
	fake.setUncommittedChangesMutex.Unlock()
	if fake.SetUncommittedChangesStub != nil {
		fake.SetUncommittedChangesStub(arg1)
	}
}

func (fake *FakeRelease) SetUncommittedChangesCallCount() int {
	fake.setUncommittedChangesMutex.RLock()
	defer fake.setUncommittedChangesMutex.RUnlock()
	return len(fake.setUncommittedChangesArgsForCall)
}

func (fake *FakeRelease) SetUncommittedChangesCalls(stub func(bool)) {
	fake.setUncommittedChangesMutex.Lock()
	defer fake.setUncommittedChangesMutex.Unlock()
	fake.SetUncommittedChangesStub = stub
}

func (fake *FakeRelease) SetUncommittedChangesArgsForCall(i int) bool {
	fake.setUncommittedChangesMutex.RLock()
	defer fake.setUncommittedChangesMutex.RUnlock()
	argsForCall := fake.setUncommittedChangesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRelease) SetVersion(arg1 string) {
	fake.setVersionMutex.Lock()
	fake.setVersionArgsForCall = append(fake.setVersionArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("SetVersion", []interface{}{arg1})
	fake.setVersionMutex.Unlock()
	if fake.SetVersionStub != nil {
		fake.SetVersionStub(arg1)
	}
}

func (fake *FakeRelease) SetVersionCallCount() int {
	fake.setVersionMutex.RLock()
	defer fake.setVersionMutex.RUnlock()
	return len(fake.setVersionArgsForCall)
}

func (fake *FakeRelease) SetVersionCalls(stub func(string)) {
	fake.setVersionMutex.Lock()
	defer fake.setVersionMutex.Unlock()
	fake.SetVersionStub = stub
}

func (fake *FakeRelease) SetVersionArgsForCall(i int) string {
	fake.setVersionMutex.RLock()
	defer fake.setVersionMutex.RUnlock()
	argsForCall := fake.setVersionArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRelease) Version() string {
	fake.versionMutex.Lock()
	ret, specificReturn := fake.versionReturnsOnCall[len(fake.versionArgsForCall)]
	fake.versionArgsForCall = append(fake.versionArgsForCall, struct {
	}{})
	fake.recordInvocation("Version", []interface{}{})
	fake.versionMutex.Unlock()
	if fake.VersionStub != nil {
		return fake.VersionStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.versionReturns
	return fakeReturns.result1
}

func (fake *FakeRelease) VersionCallCount() int {
	fake.versionMutex.RLock()
	defer fake.versionMutex.RUnlock()
	return len(fake.versionArgsForCall)
}

func (fake *FakeRelease) VersionCalls(stub func() string) {
	fake.versionMutex.Lock()
	defer fake.versionMutex.Unlock()
	fake.VersionStub = stub
}

func (fake *FakeRelease) VersionReturns(result1 string) {
	fake.versionMutex.Lock()
	defer fake.versionMutex.Unlock()
	fake.VersionStub = nil
	fake.versionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeRelease) VersionReturnsOnCall(i int, result1 string) {
	fake.versionMutex.Lock()
	defer fake.versionMutex.Unlock()
	fake.VersionStub = nil
	if fake.versionReturnsOnCall == nil {
		fake.versionReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.versionReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeRelease) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.buildMutex.RLock()
	defer fake.buildMutex.RUnlock()
	fake.cleanUpMutex.RLock()
	defer fake.cleanUpMutex.RUnlock()
	fake.commitHashWithMarkMutex.RLock()
	defer fake.commitHashWithMarkMutex.RUnlock()
	fake.compiledPackagesMutex.RLock()
	defer fake.compiledPackagesMutex.RUnlock()
	fake.copyWithMutex.RLock()
	defer fake.copyWithMutex.RUnlock()
	fake.finalizeMutex.RLock()
	defer fake.finalizeMutex.RUnlock()
	fake.findJobByNameMutex.RLock()
	defer fake.findJobByNameMutex.RUnlock()
	fake.isCompiledMutex.RLock()
	defer fake.isCompiledMutex.RUnlock()
	fake.jobsMutex.RLock()
	defer fake.jobsMutex.RUnlock()
	fake.licenseMutex.RLock()
	defer fake.licenseMutex.RUnlock()
	fake.manifestMutex.RLock()
	defer fake.manifestMutex.RUnlock()
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	fake.packagesMutex.RLock()
	defer fake.packagesMutex.RUnlock()
	fake.setCommitHashMutex.RLock()
	defer fake.setCommitHashMutex.RUnlock()
	fake.setNameMutex.RLock()
	defer fake.setNameMutex.RUnlock()
	fake.setUncommittedChangesMutex.RLock()
	defer fake.setUncommittedChangesMutex.RUnlock()
	fake.setVersionMutex.RLock()
	defer fake.setVersionMutex.RUnlock()
	fake.versionMutex.RLock()
	defer fake.versionMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRelease) recordInvocation(key string, args []interface{}) {
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

var _ release.Release = new(FakeRelease)
