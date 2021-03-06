// Code generated by MockGen. DO NOT EDIT.
// Source: server/application/authentication.go

// Package mock_application is a generated GoMock package.
package mock_application

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	model "github.com/sekky0905/nuxt-vue-go-chat/server/domain/model"
	reflect "reflect"
)

// MockAuthenticationService is a mock of AuthenticationService interface
type MockAuthenticationService struct {
	ctrl     *gomock.Controller
	recorder *MockAuthenticationServiceMockRecorder
}

// MockAuthenticationServiceMockRecorder is the mock recorder for MockAuthenticationService
type MockAuthenticationServiceMockRecorder struct {
	mock *MockAuthenticationService
}

// NewMockAuthenticationService creates a new mock instance
func NewMockAuthenticationService(ctrl *gomock.Controller) *MockAuthenticationService {
	mock := &MockAuthenticationService{ctrl: ctrl}
	mock.recorder = &MockAuthenticationServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuthenticationService) EXPECT() *MockAuthenticationServiceMockRecorder {
	return m.recorder
}

// SignUp mocks base method
func (m *MockAuthenticationService) SignUp(ctx context.Context, param *model.User) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignUp", ctx, param)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignUp indicates an expected call of SignUp
func (mr *MockAuthenticationServiceMockRecorder) SignUp(ctx, param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignUp", reflect.TypeOf((*MockAuthenticationService)(nil).SignUp), ctx, param)
}

// Login mocks base method
func (m *MockAuthenticationService) Login(ctx context.Context, param *model.User) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", ctx, param)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login
func (mr *MockAuthenticationServiceMockRecorder) Login(ctx, param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockAuthenticationService)(nil).Login), ctx, param)
}

// Logout mocks base method
func (m *MockAuthenticationService) Logout(ctx context.Context, sessionID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logout", ctx, sessionID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Logout indicates an expected call of Logout
func (mr *MockAuthenticationServiceMockRecorder) Logout(ctx, sessionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logout", reflect.TypeOf((*MockAuthenticationService)(nil).Logout), ctx, sessionID)
}
