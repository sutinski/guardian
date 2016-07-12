// This file was generated by counterfeiter
package gardenerfakes

import (
	"sync"

	"code.cloudfoundry.org/lager"
	"github.com/cloudfoundry-incubator/garden"
	"github.com/cloudfoundry-incubator/garden-shed/rootfs_provider"
	"github.com/cloudfoundry-incubator/guardian/gardener"
)

type FakeVolumeCreator struct {
	CreateStub        func(log lager.Logger, handle string, spec rootfs_provider.Spec) (string, []string, error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		log    lager.Logger
		handle string
		spec   rootfs_provider.Spec
	}
	createReturns struct {
		result1 string
		result2 []string
		result3 error
	}
	DestroyStub        func(log lager.Logger, handle string) error
	destroyMutex       sync.RWMutex
	destroyArgsForCall []struct {
		log    lager.Logger
		handle string
	}
	destroyReturns struct {
		result1 error
	}
	MetricsStub        func(log lager.Logger, handle string) (garden.ContainerDiskStat, error)
	metricsMutex       sync.RWMutex
	metricsArgsForCall []struct {
		log    lager.Logger
		handle string
	}
	metricsReturns struct {
		result1 garden.ContainerDiskStat
		result2 error
	}
	GCStub        func(log lager.Logger) error
	gCMutex       sync.RWMutex
	gCArgsForCall []struct {
		log lager.Logger
	}
	gCReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeVolumeCreator) Create(log lager.Logger, handle string, spec rootfs_provider.Spec) (string, []string, error) {
	fake.createMutex.Lock()
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		log    lager.Logger
		handle string
		spec   rootfs_provider.Spec
	}{log, handle, spec})
	fake.recordInvocation("Create", []interface{}{log, handle, spec})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(log, handle, spec)
	} else {
		return fake.createReturns.result1, fake.createReturns.result2, fake.createReturns.result3
	}
}

func (fake *FakeVolumeCreator) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeVolumeCreator) CreateArgsForCall(i int) (lager.Logger, string, rootfs_provider.Spec) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].log, fake.createArgsForCall[i].handle, fake.createArgsForCall[i].spec
}

func (fake *FakeVolumeCreator) CreateReturns(result1 string, result2 []string, result3 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 string
		result2 []string
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeVolumeCreator) Destroy(log lager.Logger, handle string) error {
	fake.destroyMutex.Lock()
	fake.destroyArgsForCall = append(fake.destroyArgsForCall, struct {
		log    lager.Logger
		handle string
	}{log, handle})
	fake.recordInvocation("Destroy", []interface{}{log, handle})
	fake.destroyMutex.Unlock()
	if fake.DestroyStub != nil {
		return fake.DestroyStub(log, handle)
	} else {
		return fake.destroyReturns.result1
	}
}

func (fake *FakeVolumeCreator) DestroyCallCount() int {
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	return len(fake.destroyArgsForCall)
}

func (fake *FakeVolumeCreator) DestroyArgsForCall(i int) (lager.Logger, string) {
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	return fake.destroyArgsForCall[i].log, fake.destroyArgsForCall[i].handle
}

func (fake *FakeVolumeCreator) DestroyReturns(result1 error) {
	fake.DestroyStub = nil
	fake.destroyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeVolumeCreator) Metrics(log lager.Logger, handle string) (garden.ContainerDiskStat, error) {
	fake.metricsMutex.Lock()
	fake.metricsArgsForCall = append(fake.metricsArgsForCall, struct {
		log    lager.Logger
		handle string
	}{log, handle})
	fake.recordInvocation("Metrics", []interface{}{log, handle})
	fake.metricsMutex.Unlock()
	if fake.MetricsStub != nil {
		return fake.MetricsStub(log, handle)
	} else {
		return fake.metricsReturns.result1, fake.metricsReturns.result2
	}
}

func (fake *FakeVolumeCreator) MetricsCallCount() int {
	fake.metricsMutex.RLock()
	defer fake.metricsMutex.RUnlock()
	return len(fake.metricsArgsForCall)
}

func (fake *FakeVolumeCreator) MetricsArgsForCall(i int) (lager.Logger, string) {
	fake.metricsMutex.RLock()
	defer fake.metricsMutex.RUnlock()
	return fake.metricsArgsForCall[i].log, fake.metricsArgsForCall[i].handle
}

func (fake *FakeVolumeCreator) MetricsReturns(result1 garden.ContainerDiskStat, result2 error) {
	fake.MetricsStub = nil
	fake.metricsReturns = struct {
		result1 garden.ContainerDiskStat
		result2 error
	}{result1, result2}
}

func (fake *FakeVolumeCreator) GC(log lager.Logger) error {
	fake.gCMutex.Lock()
	fake.gCArgsForCall = append(fake.gCArgsForCall, struct {
		log lager.Logger
	}{log})
	fake.recordInvocation("GC", []interface{}{log})
	fake.gCMutex.Unlock()
	if fake.GCStub != nil {
		return fake.GCStub(log)
	} else {
		return fake.gCReturns.result1
	}
}

func (fake *FakeVolumeCreator) GCCallCount() int {
	fake.gCMutex.RLock()
	defer fake.gCMutex.RUnlock()
	return len(fake.gCArgsForCall)
}

func (fake *FakeVolumeCreator) GCArgsForCall(i int) lager.Logger {
	fake.gCMutex.RLock()
	defer fake.gCMutex.RUnlock()
	return fake.gCArgsForCall[i].log
}

func (fake *FakeVolumeCreator) GCReturns(result1 error) {
	fake.GCStub = nil
	fake.gCReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeVolumeCreator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	fake.metricsMutex.RLock()
	defer fake.metricsMutex.RUnlock()
	fake.gCMutex.RLock()
	defer fake.gCMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeVolumeCreator) recordInvocation(key string, args []interface{}) {
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

var _ gardener.VolumeCreator = new(FakeVolumeCreator)
