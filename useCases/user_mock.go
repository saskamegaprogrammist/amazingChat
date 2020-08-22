// Code generated by MockGen. DO NOT EDIT.
// Source: useCases/user_interface.go

// Package mock_useCases is a generated GoMock package.
package useCases

import (
	gomock "github.com/golang/mock/gomock"
	models "github.com/saskamegaprogrammist/amazingChat/models"
	reflect "reflect"
)

// MockUsersUCInterface is a mock of UsersUCInterface interface
type MockUsersUCInterface struct {
	ctrl     *gomock.Controller
	recorder *MockUsersUCInterfaceMockRecorder
}

// MockUsersUCInterfaceMockRecorder is the mock recorder for MockUsersUCInterface
type MockUsersUCInterfaceMockRecorder struct {
	mock *MockUsersUCInterface
}

// NewMockUsersUCInterface creates a new mock instance
func NewMockUsersUCInterface(ctrl *gomock.Controller) *MockUsersUCInterface {
	mock := &MockUsersUCInterface{ctrl: ctrl}
	mock.recorder = &MockUsersUCInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUsersUCInterface) EXPECT() *MockUsersUCInterfaceMockRecorder {
	return m.recorder
}

// Add mocks base method
func (m *MockUsersUCInterface) Add(user *models.User) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", user)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add
func (mr *MockUsersUCInterfaceMockRecorder) Add(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockUsersUCInterface)(nil).Add), user)
}