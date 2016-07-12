// This file was generated by counterfeiter
package dnsfakes

import (
	"sync"

	"code.cloudfoundry.org/lager"
	"github.com/cloudfoundry-incubator/guardian/kawasaki/dns"
)

type FakeFileWriter struct {
	WriteFileStub        func(log lager.Logger, filePath string, contents []byte) error
	writeFileMutex       sync.RWMutex
	writeFileArgsForCall []struct {
		log      lager.Logger
		filePath string
		contents []byte
	}
	writeFileReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeFileWriter) WriteFile(log lager.Logger, filePath string, contents []byte) error {
	var contentsCopy []byte
	if contents != nil {
		contentsCopy = make([]byte, len(contents))
		copy(contentsCopy, contents)
	}
	fake.writeFileMutex.Lock()
	fake.writeFileArgsForCall = append(fake.writeFileArgsForCall, struct {
		log      lager.Logger
		filePath string
		contents []byte
	}{log, filePath, contentsCopy})
	fake.recordInvocation("WriteFile", []interface{}{log, filePath, contentsCopy})
	fake.writeFileMutex.Unlock()
	if fake.WriteFileStub != nil {
		return fake.WriteFileStub(log, filePath, contents)
	} else {
		return fake.writeFileReturns.result1
	}
}

func (fake *FakeFileWriter) WriteFileCallCount() int {
	fake.writeFileMutex.RLock()
	defer fake.writeFileMutex.RUnlock()
	return len(fake.writeFileArgsForCall)
}

func (fake *FakeFileWriter) WriteFileArgsForCall(i int) (lager.Logger, string, []byte) {
	fake.writeFileMutex.RLock()
	defer fake.writeFileMutex.RUnlock()
	return fake.writeFileArgsForCall[i].log, fake.writeFileArgsForCall[i].filePath, fake.writeFileArgsForCall[i].contents
}

func (fake *FakeFileWriter) WriteFileReturns(result1 error) {
	fake.WriteFileStub = nil
	fake.writeFileReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeFileWriter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.writeFileMutex.RLock()
	defer fake.writeFileMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeFileWriter) recordInvocation(key string, args []interface{}) {
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

var _ dns.FileWriter = new(FakeFileWriter)
