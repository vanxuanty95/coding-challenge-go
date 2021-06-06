// Code generated by MockGen. DO NOT EDIT.
// Source: coding-challenge-go/pkg/api/product/v2 (interfaces: Repository)

// Package v2 is a generated GoMock package.
package v2

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// findByUUID mocks base method.
func (m *MockRepository) findByUUID(arg0 string) (*product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "findByUUID", arg0)
	ret0, _ := ret[0].(*product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// findByUUID indicates an expected call of findByUUID.
func (mr *MockRepositoryMockRecorder) findByUUID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "findByUUID", reflect.TypeOf((*MockRepository)(nil).findByUUID), arg0)
}

// list mocks base method.
func (m *MockRepository) list(arg0, arg1 int) ([]*product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "list", arg0, arg1)
	ret0, _ := ret[0].([]*product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// list indicates an expected call of list.
func (mr *MockRepositoryMockRecorder) list(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "list", reflect.TypeOf((*MockRepository)(nil).list), arg0, arg1)
}
