// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/guardian/rundmc"
	"github.com/pivotal-golang/lager"
)

type FakeStopper struct {
	StopAllStub        func(log lager.Logger, cgroupName string, save []int, kill bool) error
	stopAllMutex       sync.RWMutex
	stopAllArgsForCall []struct {
		log        lager.Logger
		cgroupName string
		save       []int
		kill       bool
	}
	stopAllReturns struct {
		result1 error
	}
}

func (fake *FakeStopper) StopAll(log lager.Logger, cgroupName string, save []int, kill bool) error {
	fake.stopAllMutex.Lock()
	fake.stopAllArgsForCall = append(fake.stopAllArgsForCall, struct {
		log        lager.Logger
		cgroupName string
		save       []int
		kill       bool
	}{log, cgroupName, save, kill})
	fake.stopAllMutex.Unlock()
	if fake.StopAllStub != nil {
		return fake.StopAllStub(log, cgroupName, save, kill)
	} else {
		return fake.stopAllReturns.result1
	}
}

func (fake *FakeStopper) StopAllCallCount() int {
	fake.stopAllMutex.RLock()
	defer fake.stopAllMutex.RUnlock()
	return len(fake.stopAllArgsForCall)
}

func (fake *FakeStopper) StopAllArgsForCall(i int) (lager.Logger, string, []int, bool) {
	fake.stopAllMutex.RLock()
	defer fake.stopAllMutex.RUnlock()
	return fake.stopAllArgsForCall[i].log, fake.stopAllArgsForCall[i].cgroupName, fake.stopAllArgsForCall[i].save, fake.stopAllArgsForCall[i].kill
}

func (fake *FakeStopper) StopAllReturns(result1 error) {
	fake.StopAllStub = nil
	fake.stopAllReturns = struct {
		result1 error
	}{result1}
}

var _ rundmc.Stopper = new(FakeStopper)