// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/google/trillian/extension (interfaces: Registry)

package extension

import (
	gomock "github.com/golang/mock/gomock"
	storage "github.com/google/trillian/storage"
)

// Mock of Registry interface
type MockRegistry struct {
	ctrl     *gomock.Controller
	recorder *_MockRegistryRecorder
}

// Recorder for MockRegistry (not exported)
type _MockRegistryRecorder struct {
	mock *MockRegistry
}

func NewMockRegistry(ctrl *gomock.Controller) *MockRegistry {
	mock := &MockRegistry{ctrl: ctrl}
	mock.recorder = &_MockRegistryRecorder{mock}
	return mock
}

func (_m *MockRegistry) EXPECT() *_MockRegistryRecorder {
	return _m.recorder
}

func (_m *MockRegistry) GetLogStorage() (storage.LogStorage, error) {
	ret := _m.ctrl.Call(_m, "GetLogStorage")
	ret0, _ := ret[0].(storage.LogStorage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockRegistryRecorder) GetLogStorage() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetLogStorage")
}

func (_m *MockRegistry) GetMapStorage() (storage.MapStorage, error) {
	ret := _m.ctrl.Call(_m, "GetMapStorage")
	ret0, _ := ret[0].(storage.MapStorage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockRegistryRecorder) GetMapStorage() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetMapStorage")
}