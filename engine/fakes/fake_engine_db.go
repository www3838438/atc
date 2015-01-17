// This file was generated by counterfeiter
package fakes

import (
	"sync"
	"time"

	"github.com/concourse/atc"
	"github.com/concourse/atc/db"
	"github.com/concourse/atc/engine"
)

type FakeEngineDB struct {
	SaveBuildEventStub        func(buildID int, event atc.Event) error
	saveBuildEventMutex       sync.RWMutex
	saveBuildEventArgsForCall []struct {
		buildID int
		event   atc.Event
	}
	saveBuildEventReturns struct {
		result1 error
	}
	CompleteBuildStub        func(buildID int) error
	completeBuildMutex       sync.RWMutex
	completeBuildArgsForCall []struct {
		buildID int
	}
	completeBuildReturns struct {
		result1 error
	}
	SaveBuildEngineMetadataStub        func(buildID int, metadata string) error
	saveBuildEngineMetadataMutex       sync.RWMutex
	saveBuildEngineMetadataArgsForCall []struct {
		buildID  int
		metadata string
	}
	saveBuildEngineMetadataReturns struct {
		result1 error
	}
	SaveBuildStartTimeStub        func(buildID int, startTime time.Time) error
	saveBuildStartTimeMutex       sync.RWMutex
	saveBuildStartTimeArgsForCall []struct {
		buildID   int
		startTime time.Time
	}
	saveBuildStartTimeReturns struct {
		result1 error
	}
	SaveBuildEndTimeStub        func(buildID int, startTime time.Time) error
	saveBuildEndTimeMutex       sync.RWMutex
	saveBuildEndTimeArgsForCall []struct {
		buildID   int
		startTime time.Time
	}
	saveBuildEndTimeReturns struct {
		result1 error
	}
	SaveBuildInputStub        func(buildID int, input db.BuildInput) (db.SavedVersionedResource, error)
	saveBuildInputMutex       sync.RWMutex
	saveBuildInputArgsForCall []struct {
		buildID int
		input   db.BuildInput
	}
	saveBuildInputReturns struct {
		result1 db.SavedVersionedResource
		result2 error
	}
	SaveBuildOutputStub        func(buildID int, vr db.VersionedResource) (db.SavedVersionedResource, error)
	saveBuildOutputMutex       sync.RWMutex
	saveBuildOutputArgsForCall []struct {
		buildID int
		vr      db.VersionedResource
	}
	saveBuildOutputReturns struct {
		result1 db.SavedVersionedResource
		result2 error
	}
	SaveBuildStatusStub        func(buildID int, status db.Status) error
	saveBuildStatusMutex       sync.RWMutex
	saveBuildStatusArgsForCall []struct {
		buildID int
		status  db.Status
	}
	saveBuildStatusReturns struct {
		result1 error
	}
}

func (fake *FakeEngineDB) SaveBuildEvent(buildID int, event atc.Event) error {
	fake.saveBuildEventMutex.Lock()
	fake.saveBuildEventArgsForCall = append(fake.saveBuildEventArgsForCall, struct {
		buildID int
		event   atc.Event
	}{buildID, event})
	fake.saveBuildEventMutex.Unlock()
	if fake.SaveBuildEventStub != nil {
		return fake.SaveBuildEventStub(buildID, event)
	} else {
		return fake.saveBuildEventReturns.result1
	}
}

func (fake *FakeEngineDB) SaveBuildEventCallCount() int {
	fake.saveBuildEventMutex.RLock()
	defer fake.saveBuildEventMutex.RUnlock()
	return len(fake.saveBuildEventArgsForCall)
}

func (fake *FakeEngineDB) SaveBuildEventArgsForCall(i int) (int, atc.Event) {
	fake.saveBuildEventMutex.RLock()
	defer fake.saveBuildEventMutex.RUnlock()
	return fake.saveBuildEventArgsForCall[i].buildID, fake.saveBuildEventArgsForCall[i].event
}

func (fake *FakeEngineDB) SaveBuildEventReturns(result1 error) {
	fake.SaveBuildEventStub = nil
	fake.saveBuildEventReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeEngineDB) CompleteBuild(buildID int) error {
	fake.completeBuildMutex.Lock()
	fake.completeBuildArgsForCall = append(fake.completeBuildArgsForCall, struct {
		buildID int
	}{buildID})
	fake.completeBuildMutex.Unlock()
	if fake.CompleteBuildStub != nil {
		return fake.CompleteBuildStub(buildID)
	} else {
		return fake.completeBuildReturns.result1
	}
}

func (fake *FakeEngineDB) CompleteBuildCallCount() int {
	fake.completeBuildMutex.RLock()
	defer fake.completeBuildMutex.RUnlock()
	return len(fake.completeBuildArgsForCall)
}

func (fake *FakeEngineDB) CompleteBuildArgsForCall(i int) int {
	fake.completeBuildMutex.RLock()
	defer fake.completeBuildMutex.RUnlock()
	return fake.completeBuildArgsForCall[i].buildID
}

func (fake *FakeEngineDB) CompleteBuildReturns(result1 error) {
	fake.CompleteBuildStub = nil
	fake.completeBuildReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeEngineDB) SaveBuildEngineMetadata(buildID int, metadata string) error {
	fake.saveBuildEngineMetadataMutex.Lock()
	fake.saveBuildEngineMetadataArgsForCall = append(fake.saveBuildEngineMetadataArgsForCall, struct {
		buildID  int
		metadata string
	}{buildID, metadata})
	fake.saveBuildEngineMetadataMutex.Unlock()
	if fake.SaveBuildEngineMetadataStub != nil {
		return fake.SaveBuildEngineMetadataStub(buildID, metadata)
	} else {
		return fake.saveBuildEngineMetadataReturns.result1
	}
}

func (fake *FakeEngineDB) SaveBuildEngineMetadataCallCount() int {
	fake.saveBuildEngineMetadataMutex.RLock()
	defer fake.saveBuildEngineMetadataMutex.RUnlock()
	return len(fake.saveBuildEngineMetadataArgsForCall)
}

func (fake *FakeEngineDB) SaveBuildEngineMetadataArgsForCall(i int) (int, string) {
	fake.saveBuildEngineMetadataMutex.RLock()
	defer fake.saveBuildEngineMetadataMutex.RUnlock()
	return fake.saveBuildEngineMetadataArgsForCall[i].buildID, fake.saveBuildEngineMetadataArgsForCall[i].metadata
}

func (fake *FakeEngineDB) SaveBuildEngineMetadataReturns(result1 error) {
	fake.SaveBuildEngineMetadataStub = nil
	fake.saveBuildEngineMetadataReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeEngineDB) SaveBuildStartTime(buildID int, startTime time.Time) error {
	fake.saveBuildStartTimeMutex.Lock()
	fake.saveBuildStartTimeArgsForCall = append(fake.saveBuildStartTimeArgsForCall, struct {
		buildID   int
		startTime time.Time
	}{buildID, startTime})
	fake.saveBuildStartTimeMutex.Unlock()
	if fake.SaveBuildStartTimeStub != nil {
		return fake.SaveBuildStartTimeStub(buildID, startTime)
	} else {
		return fake.saveBuildStartTimeReturns.result1
	}
}

func (fake *FakeEngineDB) SaveBuildStartTimeCallCount() int {
	fake.saveBuildStartTimeMutex.RLock()
	defer fake.saveBuildStartTimeMutex.RUnlock()
	return len(fake.saveBuildStartTimeArgsForCall)
}

func (fake *FakeEngineDB) SaveBuildStartTimeArgsForCall(i int) (int, time.Time) {
	fake.saveBuildStartTimeMutex.RLock()
	defer fake.saveBuildStartTimeMutex.RUnlock()
	return fake.saveBuildStartTimeArgsForCall[i].buildID, fake.saveBuildStartTimeArgsForCall[i].startTime
}

func (fake *FakeEngineDB) SaveBuildStartTimeReturns(result1 error) {
	fake.SaveBuildStartTimeStub = nil
	fake.saveBuildStartTimeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeEngineDB) SaveBuildEndTime(buildID int, startTime time.Time) error {
	fake.saveBuildEndTimeMutex.Lock()
	fake.saveBuildEndTimeArgsForCall = append(fake.saveBuildEndTimeArgsForCall, struct {
		buildID   int
		startTime time.Time
	}{buildID, startTime})
	fake.saveBuildEndTimeMutex.Unlock()
	if fake.SaveBuildEndTimeStub != nil {
		return fake.SaveBuildEndTimeStub(buildID, startTime)
	} else {
		return fake.saveBuildEndTimeReturns.result1
	}
}

func (fake *FakeEngineDB) SaveBuildEndTimeCallCount() int {
	fake.saveBuildEndTimeMutex.RLock()
	defer fake.saveBuildEndTimeMutex.RUnlock()
	return len(fake.saveBuildEndTimeArgsForCall)
}

func (fake *FakeEngineDB) SaveBuildEndTimeArgsForCall(i int) (int, time.Time) {
	fake.saveBuildEndTimeMutex.RLock()
	defer fake.saveBuildEndTimeMutex.RUnlock()
	return fake.saveBuildEndTimeArgsForCall[i].buildID, fake.saveBuildEndTimeArgsForCall[i].startTime
}

func (fake *FakeEngineDB) SaveBuildEndTimeReturns(result1 error) {
	fake.SaveBuildEndTimeStub = nil
	fake.saveBuildEndTimeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeEngineDB) SaveBuildInput(buildID int, input db.BuildInput) (db.SavedVersionedResource, error) {
	fake.saveBuildInputMutex.Lock()
	fake.saveBuildInputArgsForCall = append(fake.saveBuildInputArgsForCall, struct {
		buildID int
		input   db.BuildInput
	}{buildID, input})
	fake.saveBuildInputMutex.Unlock()
	if fake.SaveBuildInputStub != nil {
		return fake.SaveBuildInputStub(buildID, input)
	} else {
		return fake.saveBuildInputReturns.result1, fake.saveBuildInputReturns.result2
	}
}

func (fake *FakeEngineDB) SaveBuildInputCallCount() int {
	fake.saveBuildInputMutex.RLock()
	defer fake.saveBuildInputMutex.RUnlock()
	return len(fake.saveBuildInputArgsForCall)
}

func (fake *FakeEngineDB) SaveBuildInputArgsForCall(i int) (int, db.BuildInput) {
	fake.saveBuildInputMutex.RLock()
	defer fake.saveBuildInputMutex.RUnlock()
	return fake.saveBuildInputArgsForCall[i].buildID, fake.saveBuildInputArgsForCall[i].input
}

func (fake *FakeEngineDB) SaveBuildInputReturns(result1 db.SavedVersionedResource, result2 error) {
	fake.SaveBuildInputStub = nil
	fake.saveBuildInputReturns = struct {
		result1 db.SavedVersionedResource
		result2 error
	}{result1, result2}
}

func (fake *FakeEngineDB) SaveBuildOutput(buildID int, vr db.VersionedResource) (db.SavedVersionedResource, error) {
	fake.saveBuildOutputMutex.Lock()
	fake.saveBuildOutputArgsForCall = append(fake.saveBuildOutputArgsForCall, struct {
		buildID int
		vr      db.VersionedResource
	}{buildID, vr})
	fake.saveBuildOutputMutex.Unlock()
	if fake.SaveBuildOutputStub != nil {
		return fake.SaveBuildOutputStub(buildID, vr)
	} else {
		return fake.saveBuildOutputReturns.result1, fake.saveBuildOutputReturns.result2
	}
}

func (fake *FakeEngineDB) SaveBuildOutputCallCount() int {
	fake.saveBuildOutputMutex.RLock()
	defer fake.saveBuildOutputMutex.RUnlock()
	return len(fake.saveBuildOutputArgsForCall)
}

func (fake *FakeEngineDB) SaveBuildOutputArgsForCall(i int) (int, db.VersionedResource) {
	fake.saveBuildOutputMutex.RLock()
	defer fake.saveBuildOutputMutex.RUnlock()
	return fake.saveBuildOutputArgsForCall[i].buildID, fake.saveBuildOutputArgsForCall[i].vr
}

func (fake *FakeEngineDB) SaveBuildOutputReturns(result1 db.SavedVersionedResource, result2 error) {
	fake.SaveBuildOutputStub = nil
	fake.saveBuildOutputReturns = struct {
		result1 db.SavedVersionedResource
		result2 error
	}{result1, result2}
}

func (fake *FakeEngineDB) SaveBuildStatus(buildID int, status db.Status) error {
	fake.saveBuildStatusMutex.Lock()
	fake.saveBuildStatusArgsForCall = append(fake.saveBuildStatusArgsForCall, struct {
		buildID int
		status  db.Status
	}{buildID, status})
	fake.saveBuildStatusMutex.Unlock()
	if fake.SaveBuildStatusStub != nil {
		return fake.SaveBuildStatusStub(buildID, status)
	} else {
		return fake.saveBuildStatusReturns.result1
	}
}

func (fake *FakeEngineDB) SaveBuildStatusCallCount() int {
	fake.saveBuildStatusMutex.RLock()
	defer fake.saveBuildStatusMutex.RUnlock()
	return len(fake.saveBuildStatusArgsForCall)
}

func (fake *FakeEngineDB) SaveBuildStatusArgsForCall(i int) (int, db.Status) {
	fake.saveBuildStatusMutex.RLock()
	defer fake.saveBuildStatusMutex.RUnlock()
	return fake.saveBuildStatusArgsForCall[i].buildID, fake.saveBuildStatusArgsForCall[i].status
}

func (fake *FakeEngineDB) SaveBuildStatusReturns(result1 error) {
	fake.SaveBuildStatusStub = nil
	fake.saveBuildStatusReturns = struct {
		result1 error
	}{result1}
}

var _ engine.EngineDB = new(FakeEngineDB)
