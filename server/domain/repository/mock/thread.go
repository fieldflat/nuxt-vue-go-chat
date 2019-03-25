// Code generated by MockGen. DO NOT EDIT.
// Source: domain/repository/thread.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	model "github.com/sekky0905/nuxt-vue-go-chat/server/domain/model"
	repository "github.com/sekky0905/nuxt-vue-go-chat/server/domain/repository"
	reflect "reflect"
)

// MockThreadRepository is a mock of ThreadRepository interface
type MockThreadRepository struct {
	ctrl     *gomock.Controller
	recorder *MockThreadRepositoryMockRecorder
}

// MockThreadRepositoryMockRecorder is the mock recorder for MockThreadRepository
type MockThreadRepositoryMockRecorder struct {
	mock *MockThreadRepository
}

// NewMockThreadRepository creates a new mock instance
func NewMockThreadRepository(ctrl *gomock.Controller) *MockThreadRepository {
	mock := &MockThreadRepository{ctrl: ctrl}
	mock.recorder = &MockThreadRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockThreadRepository) EXPECT() *MockThreadRepositoryMockRecorder {
	return m.recorder
}

// ListThreads mocks base method
func (m_2 *MockThreadRepository) ListThreads(ctx context.Context, m repository.SQLManager, cursor uint32, limit int) (*model.ThreadList, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "ListThreads", ctx, m, cursor, limit)
	ret0, _ := ret[0].(*model.ThreadList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListThreads indicates an expected call of ListThreads
func (mr *MockThreadRepositoryMockRecorder) ListThreads(ctx, m, cursor, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListThreads", reflect.TypeOf((*MockThreadRepository)(nil).ListThreads), ctx, m, cursor, limit)
}

// GetThreadByID mocks base method
func (m_2 *MockThreadRepository) GetThreadByID(ctx context.Context, m repository.SQLManager, id uint32) (*model.Thread, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "GetThreadByID", ctx, m, id)
	ret0, _ := ret[0].(*model.Thread)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetThreadByID indicates an expected call of GetThreadByID
func (mr *MockThreadRepositoryMockRecorder) GetThreadByID(ctx, m, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetThreadByID", reflect.TypeOf((*MockThreadRepository)(nil).GetThreadByID), ctx, m, id)
}

// GetThreadByTitle mocks base method
func (m_2 *MockThreadRepository) GetThreadByTitle(ctx context.Context, m repository.SQLManager, name string) (*model.Thread, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "GetThreadByTitle", ctx, m, name)
	ret0, _ := ret[0].(*model.Thread)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetThreadByTitle indicates an expected call of GetThreadByTitle
func (mr *MockThreadRepositoryMockRecorder) GetThreadByTitle(ctx, m, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetThreadByTitle", reflect.TypeOf((*MockThreadRepository)(nil).GetThreadByTitle), ctx, m, name)
}

// InsertThread mocks base method
func (m_2 *MockThreadRepository) InsertThread(ctx context.Context, m repository.SQLManager, thead *model.Thread) (uint32, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "InsertThread", ctx, m, thead)
	ret0, _ := ret[0].(uint32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertThread indicates an expected call of InsertThread
func (mr *MockThreadRepositoryMockRecorder) InsertThread(ctx, m, thead interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertThread", reflect.TypeOf((*MockThreadRepository)(nil).InsertThread), ctx, m, thead)
}

// UpdateThread mocks base method
func (m_2 *MockThreadRepository) UpdateThread(ctx context.Context, m repository.SQLManager, id uint32, thead *model.Thread) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "UpdateThread", ctx, m, id, thead)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateThread indicates an expected call of UpdateThread
func (mr *MockThreadRepositoryMockRecorder) UpdateThread(ctx, m, id, thead interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateThread", reflect.TypeOf((*MockThreadRepository)(nil).UpdateThread), ctx, m, id, thead)
}

// DeleteThread mocks base method
func (m_2 *MockThreadRepository) DeleteThread(ctx context.Context, m repository.SQLManager, id uint32) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "DeleteThread", ctx, m, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteThread indicates an expected call of DeleteThread
func (mr *MockThreadRepositoryMockRecorder) DeleteThread(ctx, m, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteThread", reflect.TypeOf((*MockThreadRepository)(nil).DeleteThread), ctx, m, id)
}
