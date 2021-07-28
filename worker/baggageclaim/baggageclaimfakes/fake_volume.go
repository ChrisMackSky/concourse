// Code generated by counterfeiter. DO NOT EDIT.
package baggageclaimfakes

import (
	"context"
	"io"
	"sync"

	"github.com/concourse/concourse/worker/baggageclaim"
)

type FakeVolume struct {
	DestroyStub        func() error
	destroyMutex       sync.RWMutex
	destroyArgsForCall []struct {
	}
	destroyReturns struct {
		result1 error
	}
	destroyReturnsOnCall map[int]struct {
		result1 error
	}
	GetPrivilegedStub        func() (bool, error)
	getPrivilegedMutex       sync.RWMutex
	getPrivilegedArgsForCall []struct {
	}
	getPrivilegedReturns struct {
		result1 bool
		result2 error
	}
	getPrivilegedReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	GetStreamInP2pUrlStub        func(context.Context, string) (string, error)
	getStreamInP2pUrlMutex       sync.RWMutex
	getStreamInP2pUrlArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	getStreamInP2pUrlReturns struct {
		result1 string
		result2 error
	}
	getStreamInP2pUrlReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	HandleStub        func() string
	handleMutex       sync.RWMutex
	handleArgsForCall []struct {
	}
	handleReturns struct {
		result1 string
	}
	handleReturnsOnCall map[int]struct {
		result1 string
	}
	PathStub        func() string
	pathMutex       sync.RWMutex
	pathArgsForCall []struct {
	}
	pathReturns struct {
		result1 string
	}
	pathReturnsOnCall map[int]struct {
		result1 string
	}
	PropertiesStub        func() (baggageclaim.VolumeProperties, error)
	propertiesMutex       sync.RWMutex
	propertiesArgsForCall []struct {
	}
	propertiesReturns struct {
		result1 baggageclaim.VolumeProperties
		result2 error
	}
	propertiesReturnsOnCall map[int]struct {
		result1 baggageclaim.VolumeProperties
		result2 error
	}
	SetPrivilegedStub        func(bool) error
	setPrivilegedMutex       sync.RWMutex
	setPrivilegedArgsForCall []struct {
		arg1 bool
	}
	setPrivilegedReturns struct {
		result1 error
	}
	setPrivilegedReturnsOnCall map[int]struct {
		result1 error
	}
	SetPropertyStub        func(string, string) error
	setPropertyMutex       sync.RWMutex
	setPropertyArgsForCall []struct {
		arg1 string
		arg2 string
	}
	setPropertyReturns struct {
		result1 error
	}
	setPropertyReturnsOnCall map[int]struct {
		result1 error
	}
	StreamInStub        func(context.Context, string, baggageclaim.Encoding, io.Reader) error
	streamInMutex       sync.RWMutex
	streamInArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 baggageclaim.Encoding
		arg4 io.Reader
	}
	streamInReturns struct {
		result1 error
	}
	streamInReturnsOnCall map[int]struct {
		result1 error
	}
	StreamOutStub        func(context.Context, string, baggageclaim.Encoding) (io.ReadCloser, error)
	streamOutMutex       sync.RWMutex
	streamOutArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 baggageclaim.Encoding
	}
	streamOutReturns struct {
		result1 io.ReadCloser
		result2 error
	}
	streamOutReturnsOnCall map[int]struct {
		result1 io.ReadCloser
		result2 error
	}
	StreamP2pOutStub        func(context.Context, string, string, baggageclaim.Encoding) error
	streamP2pOutMutex       sync.RWMutex
	streamP2pOutArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 baggageclaim.Encoding
	}
	streamP2pOutReturns struct {
		result1 error
	}
	streamP2pOutReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeVolume) Destroy() error {
	fake.destroyMutex.Lock()
	ret, specificReturn := fake.destroyReturnsOnCall[len(fake.destroyArgsForCall)]
	fake.destroyArgsForCall = append(fake.destroyArgsForCall, struct {
	}{})
	fake.recordInvocation("Destroy", []interface{}{})
	fake.destroyMutex.Unlock()
	if fake.DestroyStub != nil {
		return fake.DestroyStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.destroyReturns
	return fakeReturns.result1
}

func (fake *FakeVolume) DestroyCallCount() int {
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	return len(fake.destroyArgsForCall)
}

func (fake *FakeVolume) DestroyCalls(stub func() error) {
	fake.destroyMutex.Lock()
	defer fake.destroyMutex.Unlock()
	fake.DestroyStub = stub
}

func (fake *FakeVolume) DestroyReturns(result1 error) {
	fake.destroyMutex.Lock()
	defer fake.destroyMutex.Unlock()
	fake.DestroyStub = nil
	fake.destroyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeVolume) DestroyReturnsOnCall(i int, result1 error) {
	fake.destroyMutex.Lock()
	defer fake.destroyMutex.Unlock()
	fake.DestroyStub = nil
	if fake.destroyReturnsOnCall == nil {
		fake.destroyReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.destroyReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeVolume) GetPrivileged() (bool, error) {
	fake.getPrivilegedMutex.Lock()
	ret, specificReturn := fake.getPrivilegedReturnsOnCall[len(fake.getPrivilegedArgsForCall)]
	fake.getPrivilegedArgsForCall = append(fake.getPrivilegedArgsForCall, struct {
	}{})
	fake.recordInvocation("GetPrivileged", []interface{}{})
	fake.getPrivilegedMutex.Unlock()
	if fake.GetPrivilegedStub != nil {
		return fake.GetPrivilegedStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getPrivilegedReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeVolume) GetPrivilegedCallCount() int {
	fake.getPrivilegedMutex.RLock()
	defer fake.getPrivilegedMutex.RUnlock()
	return len(fake.getPrivilegedArgsForCall)
}

func (fake *FakeVolume) GetPrivilegedCalls(stub func() (bool, error)) {
	fake.getPrivilegedMutex.Lock()
	defer fake.getPrivilegedMutex.Unlock()
	fake.GetPrivilegedStub = stub
}

func (fake *FakeVolume) GetPrivilegedReturns(result1 bool, result2 error) {
	fake.getPrivilegedMutex.Lock()
	defer fake.getPrivilegedMutex.Unlock()
	fake.GetPrivilegedStub = nil
	fake.getPrivilegedReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeVolume) GetPrivilegedReturnsOnCall(i int, result1 bool, result2 error) {
	fake.getPrivilegedMutex.Lock()
	defer fake.getPrivilegedMutex.Unlock()
	fake.GetPrivilegedStub = nil
	if fake.getPrivilegedReturnsOnCall == nil {
		fake.getPrivilegedReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.getPrivilegedReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeVolume) GetStreamInP2pUrl(arg1 context.Context, arg2 string) (string, error) {
	fake.getStreamInP2pUrlMutex.Lock()
	ret, specificReturn := fake.getStreamInP2pUrlReturnsOnCall[len(fake.getStreamInP2pUrlArgsForCall)]
	fake.getStreamInP2pUrlArgsForCall = append(fake.getStreamInP2pUrlArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetStreamInP2pUrl", []interface{}{arg1, arg2})
	fake.getStreamInP2pUrlMutex.Unlock()
	if fake.GetStreamInP2pUrlStub != nil {
		return fake.GetStreamInP2pUrlStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getStreamInP2pUrlReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeVolume) GetStreamInP2pUrlCallCount() int {
	fake.getStreamInP2pUrlMutex.RLock()
	defer fake.getStreamInP2pUrlMutex.RUnlock()
	return len(fake.getStreamInP2pUrlArgsForCall)
}

func (fake *FakeVolume) GetStreamInP2pUrlCalls(stub func(context.Context, string) (string, error)) {
	fake.getStreamInP2pUrlMutex.Lock()
	defer fake.getStreamInP2pUrlMutex.Unlock()
	fake.GetStreamInP2pUrlStub = stub
}

func (fake *FakeVolume) GetStreamInP2pUrlArgsForCall(i int) (context.Context, string) {
	fake.getStreamInP2pUrlMutex.RLock()
	defer fake.getStreamInP2pUrlMutex.RUnlock()
	argsForCall := fake.getStreamInP2pUrlArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeVolume) GetStreamInP2pUrlReturns(result1 string, result2 error) {
	fake.getStreamInP2pUrlMutex.Lock()
	defer fake.getStreamInP2pUrlMutex.Unlock()
	fake.GetStreamInP2pUrlStub = nil
	fake.getStreamInP2pUrlReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeVolume) GetStreamInP2pUrlReturnsOnCall(i int, result1 string, result2 error) {
	fake.getStreamInP2pUrlMutex.Lock()
	defer fake.getStreamInP2pUrlMutex.Unlock()
	fake.GetStreamInP2pUrlStub = nil
	if fake.getStreamInP2pUrlReturnsOnCall == nil {
		fake.getStreamInP2pUrlReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getStreamInP2pUrlReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeVolume) Handle() string {
	fake.handleMutex.Lock()
	ret, specificReturn := fake.handleReturnsOnCall[len(fake.handleArgsForCall)]
	fake.handleArgsForCall = append(fake.handleArgsForCall, struct {
	}{})
	fake.recordInvocation("Handle", []interface{}{})
	fake.handleMutex.Unlock()
	if fake.HandleStub != nil {
		return fake.HandleStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.handleReturns
	return fakeReturns.result1
}

func (fake *FakeVolume) HandleCallCount() int {
	fake.handleMutex.RLock()
	defer fake.handleMutex.RUnlock()
	return len(fake.handleArgsForCall)
}

func (fake *FakeVolume) HandleCalls(stub func() string) {
	fake.handleMutex.Lock()
	defer fake.handleMutex.Unlock()
	fake.HandleStub = stub
}

func (fake *FakeVolume) HandleReturns(result1 string) {
	fake.handleMutex.Lock()
	defer fake.handleMutex.Unlock()
	fake.HandleStub = nil
	fake.handleReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeVolume) HandleReturnsOnCall(i int, result1 string) {
	fake.handleMutex.Lock()
	defer fake.handleMutex.Unlock()
	fake.HandleStub = nil
	if fake.handleReturnsOnCall == nil {
		fake.handleReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.handleReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeVolume) Path() string {
	fake.pathMutex.Lock()
	ret, specificReturn := fake.pathReturnsOnCall[len(fake.pathArgsForCall)]
	fake.pathArgsForCall = append(fake.pathArgsForCall, struct {
	}{})
	fake.recordInvocation("Path", []interface{}{})
	fake.pathMutex.Unlock()
	if fake.PathStub != nil {
		return fake.PathStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.pathReturns
	return fakeReturns.result1
}

func (fake *FakeVolume) PathCallCount() int {
	fake.pathMutex.RLock()
	defer fake.pathMutex.RUnlock()
	return len(fake.pathArgsForCall)
}

func (fake *FakeVolume) PathCalls(stub func() string) {
	fake.pathMutex.Lock()
	defer fake.pathMutex.Unlock()
	fake.PathStub = stub
}

func (fake *FakeVolume) PathReturns(result1 string) {
	fake.pathMutex.Lock()
	defer fake.pathMutex.Unlock()
	fake.PathStub = nil
	fake.pathReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeVolume) PathReturnsOnCall(i int, result1 string) {
	fake.pathMutex.Lock()
	defer fake.pathMutex.Unlock()
	fake.PathStub = nil
	if fake.pathReturnsOnCall == nil {
		fake.pathReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.pathReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeVolume) Properties() (baggageclaim.VolumeProperties, error) {
	fake.propertiesMutex.Lock()
	ret, specificReturn := fake.propertiesReturnsOnCall[len(fake.propertiesArgsForCall)]
	fake.propertiesArgsForCall = append(fake.propertiesArgsForCall, struct {
	}{})
	fake.recordInvocation("Properties", []interface{}{})
	fake.propertiesMutex.Unlock()
	if fake.PropertiesStub != nil {
		return fake.PropertiesStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.propertiesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeVolume) PropertiesCallCount() int {
	fake.propertiesMutex.RLock()
	defer fake.propertiesMutex.RUnlock()
	return len(fake.propertiesArgsForCall)
}

func (fake *FakeVolume) PropertiesCalls(stub func() (baggageclaim.VolumeProperties, error)) {
	fake.propertiesMutex.Lock()
	defer fake.propertiesMutex.Unlock()
	fake.PropertiesStub = stub
}

func (fake *FakeVolume) PropertiesReturns(result1 baggageclaim.VolumeProperties, result2 error) {
	fake.propertiesMutex.Lock()
	defer fake.propertiesMutex.Unlock()
	fake.PropertiesStub = nil
	fake.propertiesReturns = struct {
		result1 baggageclaim.VolumeProperties
		result2 error
	}{result1, result2}
}

func (fake *FakeVolume) PropertiesReturnsOnCall(i int, result1 baggageclaim.VolumeProperties, result2 error) {
	fake.propertiesMutex.Lock()
	defer fake.propertiesMutex.Unlock()
	fake.PropertiesStub = nil
	if fake.propertiesReturnsOnCall == nil {
		fake.propertiesReturnsOnCall = make(map[int]struct {
			result1 baggageclaim.VolumeProperties
			result2 error
		})
	}
	fake.propertiesReturnsOnCall[i] = struct {
		result1 baggageclaim.VolumeProperties
		result2 error
	}{result1, result2}
}

func (fake *FakeVolume) SetPrivileged(arg1 bool) error {
	fake.setPrivilegedMutex.Lock()
	ret, specificReturn := fake.setPrivilegedReturnsOnCall[len(fake.setPrivilegedArgsForCall)]
	fake.setPrivilegedArgsForCall = append(fake.setPrivilegedArgsForCall, struct {
		arg1 bool
	}{arg1})
	fake.recordInvocation("SetPrivileged", []interface{}{arg1})
	fake.setPrivilegedMutex.Unlock()
	if fake.SetPrivilegedStub != nil {
		return fake.SetPrivilegedStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.setPrivilegedReturns
	return fakeReturns.result1
}

func (fake *FakeVolume) SetPrivilegedCallCount() int {
	fake.setPrivilegedMutex.RLock()
	defer fake.setPrivilegedMutex.RUnlock()
	return len(fake.setPrivilegedArgsForCall)
}

func (fake *FakeVolume) SetPrivilegedCalls(stub func(bool) error) {
	fake.setPrivilegedMutex.Lock()
	defer fake.setPrivilegedMutex.Unlock()
	fake.SetPrivilegedStub = stub
}

func (fake *FakeVolume) SetPrivilegedArgsForCall(i int) bool {
	fake.setPrivilegedMutex.RLock()
	defer fake.setPrivilegedMutex.RUnlock()
	argsForCall := fake.setPrivilegedArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeVolume) SetPrivilegedReturns(result1 error) {
	fake.setPrivilegedMutex.Lock()
	defer fake.setPrivilegedMutex.Unlock()
	fake.SetPrivilegedStub = nil
	fake.setPrivilegedReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeVolume) SetPrivilegedReturnsOnCall(i int, result1 error) {
	fake.setPrivilegedMutex.Lock()
	defer fake.setPrivilegedMutex.Unlock()
	fake.SetPrivilegedStub = nil
	if fake.setPrivilegedReturnsOnCall == nil {
		fake.setPrivilegedReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setPrivilegedReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeVolume) SetProperty(arg1 string, arg2 string) error {
	fake.setPropertyMutex.Lock()
	ret, specificReturn := fake.setPropertyReturnsOnCall[len(fake.setPropertyArgsForCall)]
	fake.setPropertyArgsForCall = append(fake.setPropertyArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("SetProperty", []interface{}{arg1, arg2})
	fake.setPropertyMutex.Unlock()
	if fake.SetPropertyStub != nil {
		return fake.SetPropertyStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.setPropertyReturns
	return fakeReturns.result1
}

func (fake *FakeVolume) SetPropertyCallCount() int {
	fake.setPropertyMutex.RLock()
	defer fake.setPropertyMutex.RUnlock()
	return len(fake.setPropertyArgsForCall)
}

func (fake *FakeVolume) SetPropertyCalls(stub func(string, string) error) {
	fake.setPropertyMutex.Lock()
	defer fake.setPropertyMutex.Unlock()
	fake.SetPropertyStub = stub
}

func (fake *FakeVolume) SetPropertyArgsForCall(i int) (string, string) {
	fake.setPropertyMutex.RLock()
	defer fake.setPropertyMutex.RUnlock()
	argsForCall := fake.setPropertyArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeVolume) SetPropertyReturns(result1 error) {
	fake.setPropertyMutex.Lock()
	defer fake.setPropertyMutex.Unlock()
	fake.SetPropertyStub = nil
	fake.setPropertyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeVolume) SetPropertyReturnsOnCall(i int, result1 error) {
	fake.setPropertyMutex.Lock()
	defer fake.setPropertyMutex.Unlock()
	fake.SetPropertyStub = nil
	if fake.setPropertyReturnsOnCall == nil {
		fake.setPropertyReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setPropertyReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeVolume) StreamIn(arg1 context.Context, arg2 string, arg3 baggageclaim.Encoding, arg4 io.Reader) error {
	fake.streamInMutex.Lock()
	ret, specificReturn := fake.streamInReturnsOnCall[len(fake.streamInArgsForCall)]
	fake.streamInArgsForCall = append(fake.streamInArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 baggageclaim.Encoding
		arg4 io.Reader
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("StreamIn", []interface{}{arg1, arg2, arg3, arg4})
	fake.streamInMutex.Unlock()
	if fake.StreamInStub != nil {
		return fake.StreamInStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.streamInReturns
	return fakeReturns.result1
}

func (fake *FakeVolume) StreamInCallCount() int {
	fake.streamInMutex.RLock()
	defer fake.streamInMutex.RUnlock()
	return len(fake.streamInArgsForCall)
}

func (fake *FakeVolume) StreamInCalls(stub func(context.Context, string, baggageclaim.Encoding, io.Reader) error) {
	fake.streamInMutex.Lock()
	defer fake.streamInMutex.Unlock()
	fake.StreamInStub = stub
}

func (fake *FakeVolume) StreamInArgsForCall(i int) (context.Context, string, baggageclaim.Encoding, io.Reader) {
	fake.streamInMutex.RLock()
	defer fake.streamInMutex.RUnlock()
	argsForCall := fake.streamInArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeVolume) StreamInReturns(result1 error) {
	fake.streamInMutex.Lock()
	defer fake.streamInMutex.Unlock()
	fake.StreamInStub = nil
	fake.streamInReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeVolume) StreamInReturnsOnCall(i int, result1 error) {
	fake.streamInMutex.Lock()
	defer fake.streamInMutex.Unlock()
	fake.StreamInStub = nil
	if fake.streamInReturnsOnCall == nil {
		fake.streamInReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.streamInReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeVolume) StreamOut(arg1 context.Context, arg2 string, arg3 baggageclaim.Encoding) (io.ReadCloser, error) {
	fake.streamOutMutex.Lock()
	ret, specificReturn := fake.streamOutReturnsOnCall[len(fake.streamOutArgsForCall)]
	fake.streamOutArgsForCall = append(fake.streamOutArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 baggageclaim.Encoding
	}{arg1, arg2, arg3})
	fake.recordInvocation("StreamOut", []interface{}{arg1, arg2, arg3})
	fake.streamOutMutex.Unlock()
	if fake.StreamOutStub != nil {
		return fake.StreamOutStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.streamOutReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeVolume) StreamOutCallCount() int {
	fake.streamOutMutex.RLock()
	defer fake.streamOutMutex.RUnlock()
	return len(fake.streamOutArgsForCall)
}

func (fake *FakeVolume) StreamOutCalls(stub func(context.Context, string, baggageclaim.Encoding) (io.ReadCloser, error)) {
	fake.streamOutMutex.Lock()
	defer fake.streamOutMutex.Unlock()
	fake.StreamOutStub = stub
}

func (fake *FakeVolume) StreamOutArgsForCall(i int) (context.Context, string, baggageclaim.Encoding) {
	fake.streamOutMutex.RLock()
	defer fake.streamOutMutex.RUnlock()
	argsForCall := fake.streamOutArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeVolume) StreamOutReturns(result1 io.ReadCloser, result2 error) {
	fake.streamOutMutex.Lock()
	defer fake.streamOutMutex.Unlock()
	fake.StreamOutStub = nil
	fake.streamOutReturns = struct {
		result1 io.ReadCloser
		result2 error
	}{result1, result2}
}

func (fake *FakeVolume) StreamOutReturnsOnCall(i int, result1 io.ReadCloser, result2 error) {
	fake.streamOutMutex.Lock()
	defer fake.streamOutMutex.Unlock()
	fake.StreamOutStub = nil
	if fake.streamOutReturnsOnCall == nil {
		fake.streamOutReturnsOnCall = make(map[int]struct {
			result1 io.ReadCloser
			result2 error
		})
	}
	fake.streamOutReturnsOnCall[i] = struct {
		result1 io.ReadCloser
		result2 error
	}{result1, result2}
}

func (fake *FakeVolume) StreamP2pOut(arg1 context.Context, arg2 string, arg3 string, arg4 baggageclaim.Encoding) error {
	fake.streamP2pOutMutex.Lock()
	ret, specificReturn := fake.streamP2pOutReturnsOnCall[len(fake.streamP2pOutArgsForCall)]
	fake.streamP2pOutArgsForCall = append(fake.streamP2pOutArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 baggageclaim.Encoding
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("StreamP2pOut", []interface{}{arg1, arg2, arg3, arg4})
	fake.streamP2pOutMutex.Unlock()
	if fake.StreamP2pOutStub != nil {
		return fake.StreamP2pOutStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.streamP2pOutReturns
	return fakeReturns.result1
}

func (fake *FakeVolume) StreamP2pOutCallCount() int {
	fake.streamP2pOutMutex.RLock()
	defer fake.streamP2pOutMutex.RUnlock()
	return len(fake.streamP2pOutArgsForCall)
}

func (fake *FakeVolume) StreamP2pOutCalls(stub func(context.Context, string, string, baggageclaim.Encoding) error) {
	fake.streamP2pOutMutex.Lock()
	defer fake.streamP2pOutMutex.Unlock()
	fake.StreamP2pOutStub = stub
}

func (fake *FakeVolume) StreamP2pOutArgsForCall(i int) (context.Context, string, string, baggageclaim.Encoding) {
	fake.streamP2pOutMutex.RLock()
	defer fake.streamP2pOutMutex.RUnlock()
	argsForCall := fake.streamP2pOutArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeVolume) StreamP2pOutReturns(result1 error) {
	fake.streamP2pOutMutex.Lock()
	defer fake.streamP2pOutMutex.Unlock()
	fake.StreamP2pOutStub = nil
	fake.streamP2pOutReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeVolume) StreamP2pOutReturnsOnCall(i int, result1 error) {
	fake.streamP2pOutMutex.Lock()
	defer fake.streamP2pOutMutex.Unlock()
	fake.StreamP2pOutStub = nil
	if fake.streamP2pOutReturnsOnCall == nil {
		fake.streamP2pOutReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.streamP2pOutReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeVolume) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	fake.getPrivilegedMutex.RLock()
	defer fake.getPrivilegedMutex.RUnlock()
	fake.getStreamInP2pUrlMutex.RLock()
	defer fake.getStreamInP2pUrlMutex.RUnlock()
	fake.handleMutex.RLock()
	defer fake.handleMutex.RUnlock()
	fake.pathMutex.RLock()
	defer fake.pathMutex.RUnlock()
	fake.propertiesMutex.RLock()
	defer fake.propertiesMutex.RUnlock()
	fake.setPrivilegedMutex.RLock()
	defer fake.setPrivilegedMutex.RUnlock()
	fake.setPropertyMutex.RLock()
	defer fake.setPropertyMutex.RUnlock()
	fake.streamInMutex.RLock()
	defer fake.streamInMutex.RUnlock()
	fake.streamOutMutex.RLock()
	defer fake.streamOutMutex.RUnlock()
	fake.streamP2pOutMutex.RLock()
	defer fake.streamP2pOutMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeVolume) recordInvocation(key string, args []interface{}) {
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

var _ baggageclaim.Volume = new(FakeVolume)
