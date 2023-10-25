// Code generated by MockGen. DO NOT EDIT.
// Source: internal/core/ports/repository.go

// Package mockrepository is a generated GoMock package.
package mockrepository

import (
	domain "hugeman/internal/core/domain"
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

// CreateTodo mocks base method.
func (m *MockRepository) CreateTodo(request domain.TodoRequest) (*domain.TodoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTodo", request)
	ret0, _ := ret[0].(*domain.TodoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTodo indicates an expected call of CreateTodo.
func (mr *MockRepositoryMockRecorder) CreateTodo(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTodo", reflect.TypeOf((*MockRepository)(nil).CreateTodo), request)
}

// DeleteTodo mocks base method.
func (m *MockRepository) DeleteTodo(request domain.TodoRequest) (*domain.TodoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTodo", request)
	ret0, _ := ret[0].(*domain.TodoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteTodo indicates an expected call of DeleteTodo.
func (mr *MockRepositoryMockRecorder) DeleteTodo(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTodo", reflect.TypeOf((*MockRepository)(nil).DeleteTodo), request)
}

// GetTodo mocks base method.
func (m *MockRepository) GetTodo(condition domain.QueryTodoRequest) (*domain.TodoListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTodo", condition)
	ret0, _ := ret[0].(*domain.TodoListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTodo indicates an expected call of GetTodo.
func (mr *MockRepositoryMockRecorder) GetTodo(condition interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTodo", reflect.TypeOf((*MockRepository)(nil).GetTodo), condition)
}

// UpdateTodo mocks base method.
func (m *MockRepository) UpdateTodo(request domain.TodoRequest) (*domain.TodoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTodo", request)
	ret0, _ := ret[0].(*domain.TodoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTodo indicates an expected call of UpdateTodo.
func (mr *MockRepositoryMockRecorder) UpdateTodo(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTodo", reflect.TypeOf((*MockRepository)(nil).UpdateTodo), request)
}
