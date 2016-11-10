// This file was generated by counterfeiter
package workerfakes

import (
	"sync"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/atc/dbng"
	"github.com/concourse/atc/worker"
)

type FakeVolumeClient struct {
	FindOrCreateVolumeForResourceCacheStub        func(lager.Logger, worker.VolumeSpec, *dbng.UsedResourceCache) (worker.Volume, error)
	findOrCreateVolumeForResourceCacheMutex       sync.RWMutex
	findOrCreateVolumeForResourceCacheArgsForCall []struct {
		arg1 lager.Logger
		arg2 worker.VolumeSpec
		arg3 *dbng.UsedResourceCache
	}
	findOrCreateVolumeForResourceCacheReturns struct {
		result1 worker.Volume
		result2 error
	}
	FindOrCreateVolumeForContainerStub        func(lager.Logger, worker.VolumeSpec, *dbng.CreatingContainer, *dbng.Team, string) (worker.Volume, error)
	findOrCreateVolumeForContainerMutex       sync.RWMutex
	findOrCreateVolumeForContainerArgsForCall []struct {
		arg1 lager.Logger
		arg2 worker.VolumeSpec
		arg3 *dbng.CreatingContainer
		arg4 *dbng.Team
		arg5 string
	}
	findOrCreateVolumeForContainerReturns struct {
		result1 worker.Volume
		result2 error
	}
	FindOrCreateVolumeForBaseResourceTypeStub        func(lager.Logger, worker.VolumeSpec, *dbng.Team, string) (worker.Volume, error)
	findOrCreateVolumeForBaseResourceTypeMutex       sync.RWMutex
	findOrCreateVolumeForBaseResourceTypeArgsForCall []struct {
		arg1 lager.Logger
		arg2 worker.VolumeSpec
		arg3 *dbng.Team
		arg4 string
	}
	findOrCreateVolumeForBaseResourceTypeReturns struct {
		result1 worker.Volume
		result2 error
	}
	LookupVolumeStub        func(lager.Logger, string) (worker.Volume, bool, error)
	lookupVolumeMutex       sync.RWMutex
	lookupVolumeArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
	}
	lookupVolumeReturns struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeVolumeClient) FindOrCreateVolumeForResourceCache(arg1 lager.Logger, arg2 worker.VolumeSpec, arg3 *dbng.UsedResourceCache) (worker.Volume, error) {
	fake.findOrCreateVolumeForResourceCacheMutex.Lock()
	fake.findOrCreateVolumeForResourceCacheArgsForCall = append(fake.findOrCreateVolumeForResourceCacheArgsForCall, struct {
		arg1 lager.Logger
		arg2 worker.VolumeSpec
		arg3 *dbng.UsedResourceCache
	}{arg1, arg2, arg3})
	fake.recordInvocation("FindOrCreateVolumeForResourceCache", []interface{}{arg1, arg2, arg3})
	fake.findOrCreateVolumeForResourceCacheMutex.Unlock()
	if fake.FindOrCreateVolumeForResourceCacheStub != nil {
		return fake.FindOrCreateVolumeForResourceCacheStub(arg1, arg2, arg3)
	} else {
		return fake.findOrCreateVolumeForResourceCacheReturns.result1, fake.findOrCreateVolumeForResourceCacheReturns.result2
	}
}

func (fake *FakeVolumeClient) FindOrCreateVolumeForResourceCacheCallCount() int {
	fake.findOrCreateVolumeForResourceCacheMutex.RLock()
	defer fake.findOrCreateVolumeForResourceCacheMutex.RUnlock()
	return len(fake.findOrCreateVolumeForResourceCacheArgsForCall)
}

func (fake *FakeVolumeClient) FindOrCreateVolumeForResourceCacheArgsForCall(i int) (lager.Logger, worker.VolumeSpec, *dbng.UsedResourceCache) {
	fake.findOrCreateVolumeForResourceCacheMutex.RLock()
	defer fake.findOrCreateVolumeForResourceCacheMutex.RUnlock()
	return fake.findOrCreateVolumeForResourceCacheArgsForCall[i].arg1, fake.findOrCreateVolumeForResourceCacheArgsForCall[i].arg2, fake.findOrCreateVolumeForResourceCacheArgsForCall[i].arg3
}

func (fake *FakeVolumeClient) FindOrCreateVolumeForResourceCacheReturns(result1 worker.Volume, result2 error) {
	fake.FindOrCreateVolumeForResourceCacheStub = nil
	fake.findOrCreateVolumeForResourceCacheReturns = struct {
		result1 worker.Volume
		result2 error
	}{result1, result2}
}

func (fake *FakeVolumeClient) FindOrCreateVolumeForContainer(arg1 lager.Logger, arg2 worker.VolumeSpec, arg3 *dbng.CreatingContainer, arg4 *dbng.Team, arg5 string) (worker.Volume, error) {
	fake.findOrCreateVolumeForContainerMutex.Lock()
	fake.findOrCreateVolumeForContainerArgsForCall = append(fake.findOrCreateVolumeForContainerArgsForCall, struct {
		arg1 lager.Logger
		arg2 worker.VolumeSpec
		arg3 *dbng.CreatingContainer
		arg4 *dbng.Team
		arg5 string
	}{arg1, arg2, arg3, arg4, arg5})
	fake.recordInvocation("FindOrCreateVolumeForContainer", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.findOrCreateVolumeForContainerMutex.Unlock()
	if fake.FindOrCreateVolumeForContainerStub != nil {
		return fake.FindOrCreateVolumeForContainerStub(arg1, arg2, arg3, arg4, arg5)
	} else {
		return fake.findOrCreateVolumeForContainerReturns.result1, fake.findOrCreateVolumeForContainerReturns.result2
	}
}

func (fake *FakeVolumeClient) FindOrCreateVolumeForContainerCallCount() int {
	fake.findOrCreateVolumeForContainerMutex.RLock()
	defer fake.findOrCreateVolumeForContainerMutex.RUnlock()
	return len(fake.findOrCreateVolumeForContainerArgsForCall)
}

func (fake *FakeVolumeClient) FindOrCreateVolumeForContainerArgsForCall(i int) (lager.Logger, worker.VolumeSpec, *dbng.CreatingContainer, *dbng.Team, string) {
	fake.findOrCreateVolumeForContainerMutex.RLock()
	defer fake.findOrCreateVolumeForContainerMutex.RUnlock()
	return fake.findOrCreateVolumeForContainerArgsForCall[i].arg1, fake.findOrCreateVolumeForContainerArgsForCall[i].arg2, fake.findOrCreateVolumeForContainerArgsForCall[i].arg3, fake.findOrCreateVolumeForContainerArgsForCall[i].arg4, fake.findOrCreateVolumeForContainerArgsForCall[i].arg5
}

func (fake *FakeVolumeClient) FindOrCreateVolumeForContainerReturns(result1 worker.Volume, result2 error) {
	fake.FindOrCreateVolumeForContainerStub = nil
	fake.findOrCreateVolumeForContainerReturns = struct {
		result1 worker.Volume
		result2 error
	}{result1, result2}
}

func (fake *FakeVolumeClient) FindOrCreateVolumeForBaseResourceType(arg1 lager.Logger, arg2 worker.VolumeSpec, arg3 *dbng.Team, arg4 string) (worker.Volume, error) {
	fake.findOrCreateVolumeForBaseResourceTypeMutex.Lock()
	fake.findOrCreateVolumeForBaseResourceTypeArgsForCall = append(fake.findOrCreateVolumeForBaseResourceTypeArgsForCall, struct {
		arg1 lager.Logger
		arg2 worker.VolumeSpec
		arg3 *dbng.Team
		arg4 string
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("FindOrCreateVolumeForBaseResourceType", []interface{}{arg1, arg2, arg3, arg4})
	fake.findOrCreateVolumeForBaseResourceTypeMutex.Unlock()
	if fake.FindOrCreateVolumeForBaseResourceTypeStub != nil {
		return fake.FindOrCreateVolumeForBaseResourceTypeStub(arg1, arg2, arg3, arg4)
	} else {
		return fake.findOrCreateVolumeForBaseResourceTypeReturns.result1, fake.findOrCreateVolumeForBaseResourceTypeReturns.result2
	}
}

func (fake *FakeVolumeClient) FindOrCreateVolumeForBaseResourceTypeCallCount() int {
	fake.findOrCreateVolumeForBaseResourceTypeMutex.RLock()
	defer fake.findOrCreateVolumeForBaseResourceTypeMutex.RUnlock()
	return len(fake.findOrCreateVolumeForBaseResourceTypeArgsForCall)
}

func (fake *FakeVolumeClient) FindOrCreateVolumeForBaseResourceTypeArgsForCall(i int) (lager.Logger, worker.VolumeSpec, *dbng.Team, string) {
	fake.findOrCreateVolumeForBaseResourceTypeMutex.RLock()
	defer fake.findOrCreateVolumeForBaseResourceTypeMutex.RUnlock()
	return fake.findOrCreateVolumeForBaseResourceTypeArgsForCall[i].arg1, fake.findOrCreateVolumeForBaseResourceTypeArgsForCall[i].arg2, fake.findOrCreateVolumeForBaseResourceTypeArgsForCall[i].arg3, fake.findOrCreateVolumeForBaseResourceTypeArgsForCall[i].arg4
}

func (fake *FakeVolumeClient) FindOrCreateVolumeForBaseResourceTypeReturns(result1 worker.Volume, result2 error) {
	fake.FindOrCreateVolumeForBaseResourceTypeStub = nil
	fake.findOrCreateVolumeForBaseResourceTypeReturns = struct {
		result1 worker.Volume
		result2 error
	}{result1, result2}
}

func (fake *FakeVolumeClient) LookupVolume(arg1 lager.Logger, arg2 string) (worker.Volume, bool, error) {
	fake.lookupVolumeMutex.Lock()
	fake.lookupVolumeArgsForCall = append(fake.lookupVolumeArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("LookupVolume", []interface{}{arg1, arg2})
	fake.lookupVolumeMutex.Unlock()
	if fake.LookupVolumeStub != nil {
		return fake.LookupVolumeStub(arg1, arg2)
	} else {
		return fake.lookupVolumeReturns.result1, fake.lookupVolumeReturns.result2, fake.lookupVolumeReturns.result3
	}
}

func (fake *FakeVolumeClient) LookupVolumeCallCount() int {
	fake.lookupVolumeMutex.RLock()
	defer fake.lookupVolumeMutex.RUnlock()
	return len(fake.lookupVolumeArgsForCall)
}

func (fake *FakeVolumeClient) LookupVolumeArgsForCall(i int) (lager.Logger, string) {
	fake.lookupVolumeMutex.RLock()
	defer fake.lookupVolumeMutex.RUnlock()
	return fake.lookupVolumeArgsForCall[i].arg1, fake.lookupVolumeArgsForCall[i].arg2
}

func (fake *FakeVolumeClient) LookupVolumeReturns(result1 worker.Volume, result2 bool, result3 error) {
	fake.LookupVolumeStub = nil
	fake.lookupVolumeReturns = struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeVolumeClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.findOrCreateVolumeForResourceCacheMutex.RLock()
	defer fake.findOrCreateVolumeForResourceCacheMutex.RUnlock()
	fake.findOrCreateVolumeForContainerMutex.RLock()
	defer fake.findOrCreateVolumeForContainerMutex.RUnlock()
	fake.findOrCreateVolumeForBaseResourceTypeMutex.RLock()
	defer fake.findOrCreateVolumeForBaseResourceTypeMutex.RUnlock()
	fake.lookupVolumeMutex.RLock()
	defer fake.lookupVolumeMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeVolumeClient) recordInvocation(key string, args []interface{}) {
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

var _ worker.VolumeClient = new(FakeVolumeClient)
