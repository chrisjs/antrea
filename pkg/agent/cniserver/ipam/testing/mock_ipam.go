// /*
// Copyright 2019 Antrea Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// */
//

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/vmware-tanzu/antrea/pkg/agent/cniserver/ipam (interfaces: IPAMDriver)

// Package testing is a generated GoMock package.
package testing

import (
	invoke "github.com/containernetworking/cni/pkg/invoke"
	current "github.com/containernetworking/cni/pkg/types/current"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIPAMDriver is a mock of IPAMDriver interface
type MockIPAMDriver struct {
	ctrl     *gomock.Controller
	recorder *MockIPAMDriverMockRecorder
}

// MockIPAMDriverMockRecorder is the mock recorder for MockIPAMDriver
type MockIPAMDriverMockRecorder struct {
	mock *MockIPAMDriver
}

// NewMockIPAMDriver creates a new mock instance
func NewMockIPAMDriver(ctrl *gomock.Controller) *MockIPAMDriver {
	mock := &MockIPAMDriver{ctrl: ctrl}
	mock.recorder = &MockIPAMDriverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIPAMDriver) EXPECT() *MockIPAMDriverMockRecorder {
	return m.recorder
}

// Add mocks base method
func (m *MockIPAMDriver) Add(arg0 *invoke.Args, arg1 []byte) (*current.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0, arg1)
	ret0, _ := ret[0].(*current.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add
func (mr *MockIPAMDriverMockRecorder) Add(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockIPAMDriver)(nil).Add), arg0, arg1)
}

// Check mocks base method
func (m *MockIPAMDriver) Check(arg0 *invoke.Args, arg1 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Check", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Check indicates an expected call of Check
func (mr *MockIPAMDriverMockRecorder) Check(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Check", reflect.TypeOf((*MockIPAMDriver)(nil).Check), arg0, arg1)
}

// Del mocks base method
func (m *MockIPAMDriver) Del(arg0 *invoke.Args, arg1 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Del", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Del indicates an expected call of Del
func (mr *MockIPAMDriverMockRecorder) Del(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Del", reflect.TypeOf((*MockIPAMDriver)(nil).Del), arg0, arg1)
}
