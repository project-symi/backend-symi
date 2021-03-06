// Code generated by MockGen. DO NOT EDIT.
// Source: ./app/usecase/repository/user_auth_repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockUserAuthRepository is a mock of UserAuthRepository interface
type MockUserAuthRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserAuthRepositoryMockRecorder
}

// MockUserAuthRepositoryMockRecorder is the mock recorder for MockUserAuthRepository
type MockUserAuthRepositoryMockRecorder struct {
	mock *MockUserAuthRepository
}

// NewMockUserAuthRepository creates a new mock instance
func NewMockUserAuthRepository(ctrl *gomock.Controller) *MockUserAuthRepository {
	mock := &MockUserAuthRepository{ctrl: ctrl}
	mock.recorder = &MockUserAuthRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserAuthRepository) EXPECT() *MockUserAuthRepositoryMockRecorder {
	return m.recorder
}

// IssueToken mocks base method
func (m *MockUserAuthRepository) IssueToken(arg0, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IssueToken", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IssueToken indicates an expected call of IssueToken
func (mr *MockUserAuthRepositoryMockRecorder) IssueToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IssueToken", reflect.TypeOf((*MockUserAuthRepository)(nil).IssueToken), arg0, arg1)
}

// RegisterToken mocks base method
func (m *MockUserAuthRepository) RegisterToken(arg0, arg1 string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterToken", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterToken indicates an expected call of RegisterToken
func (mr *MockUserAuthRepositoryMockRecorder) RegisterToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterToken", reflect.TypeOf((*MockUserAuthRepository)(nil).RegisterToken), arg0, arg1)
}

// GetPermissionName mocks base method
func (m *MockUserAuthRepository) GetPermissionName(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPermissionName", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPermissionName indicates an expected call of GetPermissionName
func (mr *MockUserAuthRepositoryMockRecorder) GetPermissionName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPermissionName", reflect.TypeOf((*MockUserAuthRepository)(nil).GetPermissionName), arg0)
}

// ValidateToken mocks base method
func (m *MockUserAuthRepository) ValidateToken(arg0 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateToken", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateToken indicates an expected call of ValidateToken
func (mr *MockUserAuthRepositoryMockRecorder) ValidateToken(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateToken", reflect.TypeOf((*MockUserAuthRepository)(nil).ValidateToken), arg0)
}

// RevokeToken mocks base method
func (m *MockUserAuthRepository) RevokeToken(arg0 string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RevokeToken", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RevokeToken indicates an expected call of RevokeToken
func (mr *MockUserAuthRepositoryMockRecorder) RevokeToken(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokeToken", reflect.TypeOf((*MockUserAuthRepository)(nil).RevokeToken), arg0)
}
