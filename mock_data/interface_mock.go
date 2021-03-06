// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package mock_data is a generated GoMock package.
package mock_data

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockInterface is a mock of Interface interface
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// Method mocks base method
func (m *MockInterface) Method() ReturnStruct {
	ret := m.ctrl.Call(m, "Method")
	ret0, _ := ret[0].(ReturnStruct)
	return ret0
}

// Method indicates an expected call of Method
func (mr *MockInterfaceMockRecorder) Method() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Method", reflect.TypeOf((*MockInterface)(nil).Method))
}
