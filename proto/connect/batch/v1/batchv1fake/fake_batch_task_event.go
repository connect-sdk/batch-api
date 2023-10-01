// Code generated by counterfeiter. DO NOT EDIT.
package batchv1fake

import (
	"sync"

	batchv1 "github.com/connect-sdk/batch-api/proto/connect/batch/v1"
)

type FakeBatchTaskEvent struct {
	GetTaskNameStub        func() string
	getTaskNameMutex       sync.RWMutex
	getTaskNameArgsForCall []struct {
	}
	getTaskNameReturns struct {
		result1 string
	}
	getTaskNameReturnsOnCall map[int]struct {
		result1 string
	}
	GetTaskUidStub        func() string
	getTaskUidMutex       sync.RWMutex
	getTaskUidArgsForCall []struct {
	}
	getTaskUidReturns struct {
		result1 string
	}
	getTaskUidReturnsOnCall map[int]struct {
		result1 string
	}
	ParseTaskNameStub        func() (batchv1.ParsedTaskName, error)
	parseTaskNameMutex       sync.RWMutex
	parseTaskNameArgsForCall []struct {
	}
	parseTaskNameReturns struct {
		result1 batchv1.ParsedTaskName
		result2 error
	}
	parseTaskNameReturnsOnCall map[int]struct {
		result1 batchv1.ParsedTaskName
		result2 error
	}
	SetTaskNameStub        func(string)
	setTaskNameMutex       sync.RWMutex
	setTaskNameArgsForCall []struct {
		arg1 string
	}
	SetTaskUidStub        func(string)
	setTaskUidMutex       sync.RWMutex
	setTaskUidArgsForCall []struct {
		arg1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBatchTaskEvent) GetTaskName() string {
	fake.getTaskNameMutex.Lock()
	ret, specificReturn := fake.getTaskNameReturnsOnCall[len(fake.getTaskNameArgsForCall)]
	fake.getTaskNameArgsForCall = append(fake.getTaskNameArgsForCall, struct {
	}{})
	stub := fake.GetTaskNameStub
	fakeReturns := fake.getTaskNameReturns
	fake.recordInvocation("GetTaskName", []interface{}{})
	fake.getTaskNameMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeBatchTaskEvent) GetTaskNameCallCount() int {
	fake.getTaskNameMutex.RLock()
	defer fake.getTaskNameMutex.RUnlock()
	return len(fake.getTaskNameArgsForCall)
}

func (fake *FakeBatchTaskEvent) GetTaskNameCalls(stub func() string) {
	fake.getTaskNameMutex.Lock()
	defer fake.getTaskNameMutex.Unlock()
	fake.GetTaskNameStub = stub
}

func (fake *FakeBatchTaskEvent) GetTaskNameReturns(result1 string) {
	fake.getTaskNameMutex.Lock()
	defer fake.getTaskNameMutex.Unlock()
	fake.GetTaskNameStub = nil
	fake.getTaskNameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeBatchTaskEvent) GetTaskNameReturnsOnCall(i int, result1 string) {
	fake.getTaskNameMutex.Lock()
	defer fake.getTaskNameMutex.Unlock()
	fake.GetTaskNameStub = nil
	if fake.getTaskNameReturnsOnCall == nil {
		fake.getTaskNameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.getTaskNameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeBatchTaskEvent) GetTaskUid() string {
	fake.getTaskUidMutex.Lock()
	ret, specificReturn := fake.getTaskUidReturnsOnCall[len(fake.getTaskUidArgsForCall)]
	fake.getTaskUidArgsForCall = append(fake.getTaskUidArgsForCall, struct {
	}{})
	stub := fake.GetTaskUidStub
	fakeReturns := fake.getTaskUidReturns
	fake.recordInvocation("GetTaskUid", []interface{}{})
	fake.getTaskUidMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeBatchTaskEvent) GetTaskUidCallCount() int {
	fake.getTaskUidMutex.RLock()
	defer fake.getTaskUidMutex.RUnlock()
	return len(fake.getTaskUidArgsForCall)
}

func (fake *FakeBatchTaskEvent) GetTaskUidCalls(stub func() string) {
	fake.getTaskUidMutex.Lock()
	defer fake.getTaskUidMutex.Unlock()
	fake.GetTaskUidStub = stub
}

func (fake *FakeBatchTaskEvent) GetTaskUidReturns(result1 string) {
	fake.getTaskUidMutex.Lock()
	defer fake.getTaskUidMutex.Unlock()
	fake.GetTaskUidStub = nil
	fake.getTaskUidReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeBatchTaskEvent) GetTaskUidReturnsOnCall(i int, result1 string) {
	fake.getTaskUidMutex.Lock()
	defer fake.getTaskUidMutex.Unlock()
	fake.GetTaskUidStub = nil
	if fake.getTaskUidReturnsOnCall == nil {
		fake.getTaskUidReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.getTaskUidReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeBatchTaskEvent) ParseTaskName() (batchv1.ParsedTaskName, error) {
	fake.parseTaskNameMutex.Lock()
	ret, specificReturn := fake.parseTaskNameReturnsOnCall[len(fake.parseTaskNameArgsForCall)]
	fake.parseTaskNameArgsForCall = append(fake.parseTaskNameArgsForCall, struct {
	}{})
	stub := fake.ParseTaskNameStub
	fakeReturns := fake.parseTaskNameReturns
	fake.recordInvocation("ParseTaskName", []interface{}{})
	fake.parseTaskNameMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeBatchTaskEvent) ParseTaskNameCallCount() int {
	fake.parseTaskNameMutex.RLock()
	defer fake.parseTaskNameMutex.RUnlock()
	return len(fake.parseTaskNameArgsForCall)
}

func (fake *FakeBatchTaskEvent) ParseTaskNameCalls(stub func() (batchv1.ParsedTaskName, error)) {
	fake.parseTaskNameMutex.Lock()
	defer fake.parseTaskNameMutex.Unlock()
	fake.ParseTaskNameStub = stub
}

func (fake *FakeBatchTaskEvent) ParseTaskNameReturns(result1 batchv1.ParsedTaskName, result2 error) {
	fake.parseTaskNameMutex.Lock()
	defer fake.parseTaskNameMutex.Unlock()
	fake.ParseTaskNameStub = nil
	fake.parseTaskNameReturns = struct {
		result1 batchv1.ParsedTaskName
		result2 error
	}{result1, result2}
}

func (fake *FakeBatchTaskEvent) ParseTaskNameReturnsOnCall(i int, result1 batchv1.ParsedTaskName, result2 error) {
	fake.parseTaskNameMutex.Lock()
	defer fake.parseTaskNameMutex.Unlock()
	fake.ParseTaskNameStub = nil
	if fake.parseTaskNameReturnsOnCall == nil {
		fake.parseTaskNameReturnsOnCall = make(map[int]struct {
			result1 batchv1.ParsedTaskName
			result2 error
		})
	}
	fake.parseTaskNameReturnsOnCall[i] = struct {
		result1 batchv1.ParsedTaskName
		result2 error
	}{result1, result2}
}

func (fake *FakeBatchTaskEvent) SetTaskName(arg1 string) {
	fake.setTaskNameMutex.Lock()
	fake.setTaskNameArgsForCall = append(fake.setTaskNameArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.SetTaskNameStub
	fake.recordInvocation("SetTaskName", []interface{}{arg1})
	fake.setTaskNameMutex.Unlock()
	if stub != nil {
		fake.SetTaskNameStub(arg1)
	}
}

func (fake *FakeBatchTaskEvent) SetTaskNameCallCount() int {
	fake.setTaskNameMutex.RLock()
	defer fake.setTaskNameMutex.RUnlock()
	return len(fake.setTaskNameArgsForCall)
}

func (fake *FakeBatchTaskEvent) SetTaskNameCalls(stub func(string)) {
	fake.setTaskNameMutex.Lock()
	defer fake.setTaskNameMutex.Unlock()
	fake.SetTaskNameStub = stub
}

func (fake *FakeBatchTaskEvent) SetTaskNameArgsForCall(i int) string {
	fake.setTaskNameMutex.RLock()
	defer fake.setTaskNameMutex.RUnlock()
	argsForCall := fake.setTaskNameArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeBatchTaskEvent) SetTaskUid(arg1 string) {
	fake.setTaskUidMutex.Lock()
	fake.setTaskUidArgsForCall = append(fake.setTaskUidArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.SetTaskUidStub
	fake.recordInvocation("SetTaskUid", []interface{}{arg1})
	fake.setTaskUidMutex.Unlock()
	if stub != nil {
		fake.SetTaskUidStub(arg1)
	}
}

func (fake *FakeBatchTaskEvent) SetTaskUidCallCount() int {
	fake.setTaskUidMutex.RLock()
	defer fake.setTaskUidMutex.RUnlock()
	return len(fake.setTaskUidArgsForCall)
}

func (fake *FakeBatchTaskEvent) SetTaskUidCalls(stub func(string)) {
	fake.setTaskUidMutex.Lock()
	defer fake.setTaskUidMutex.Unlock()
	fake.SetTaskUidStub = stub
}

func (fake *FakeBatchTaskEvent) SetTaskUidArgsForCall(i int) string {
	fake.setTaskUidMutex.RLock()
	defer fake.setTaskUidMutex.RUnlock()
	argsForCall := fake.setTaskUidArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeBatchTaskEvent) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getTaskNameMutex.RLock()
	defer fake.getTaskNameMutex.RUnlock()
	fake.getTaskUidMutex.RLock()
	defer fake.getTaskUidMutex.RUnlock()
	fake.parseTaskNameMutex.RLock()
	defer fake.parseTaskNameMutex.RUnlock()
	fake.setTaskNameMutex.RLock()
	defer fake.setTaskNameMutex.RUnlock()
	fake.setTaskUidMutex.RLock()
	defer fake.setTaskUidMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeBatchTaskEvent) recordInvocation(key string, args []interface{}) {
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

var _ batchv1.BatchTaskEvent = new(FakeBatchTaskEvent)
