// Code generated by MockGen. DO NOT EDIT.
// Source: ../services/service.go

// Package controllers is a generated GoMock package.
package controllers

import (
	model "api/model"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAuthorService is a mock of AuthorService interface.
type MockAuthorService struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorServiceMockRecorder
}

// MockAuthorServiceMockRecorder is the mock recorder for MockAuthorService.
type MockAuthorServiceMockRecorder struct {
	mock *MockAuthorService
}

// NewMockAuthorService creates a new mock instance.
func NewMockAuthorService(ctrl *gomock.Controller) *MockAuthorService {
	mock := &MockAuthorService{ctrl: ctrl}
	mock.recorder = &MockAuthorServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthorService) EXPECT() *MockAuthorServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockAuthorService) Create(ctx context.Context, author model.AuthorBase) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, author)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockAuthorServiceMockRecorder) Create(ctx, author interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAuthorService)(nil).Create), ctx, author)
}

// Delete mocks base method.
func (m *MockAuthorService) Delete(ctx context.Context, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockAuthorServiceMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAuthorService)(nil).Delete), ctx, id)
}

// GetAll mocks base method.
func (m *MockAuthorService) GetAll(ctx context.Context, nameFilter string) ([]model.Author, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx, nameFilter)
	ret0, _ := ret[0].([]model.Author)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockAuthorServiceMockRecorder) GetAll(ctx, nameFilter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockAuthorService)(nil).GetAll), ctx, nameFilter)
}

// GetById mocks base method.
func (m *MockAuthorService) GetById(ctx context.Context, id int64) (model.Author, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", ctx, id)
	ret0, _ := ret[0].(model.Author)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockAuthorServiceMockRecorder) GetById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockAuthorService)(nil).GetById), ctx, id)
}

// MockLogger is a mock of Logger interface.
type MockLogger struct {
	ctrl     *gomock.Controller
	recorder *MockLoggerMockRecorder
}

// MockLoggerMockRecorder is the mock recorder for MockLogger.
type MockLoggerMockRecorder struct {
	mock *MockLogger
}

// NewMockLogger creates a new mock instance.
func NewMockLogger(ctrl *gomock.Controller) *MockLogger {
	mock := &MockLogger{ctrl: ctrl}
	mock.recorder = &MockLoggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogger) EXPECT() *MockLoggerMockRecorder {
	return m.recorder
}

// Printf mocks base method.
func (m *MockLogger) Printf(format string, args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Printf", varargs...)
}

// Printf indicates an expected call of Printf.
func (mr *MockLoggerMockRecorder) Printf(format interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Printf", reflect.TypeOf((*MockLogger)(nil).Printf), varargs...)
}

// MockValidator is a mock of Validator interface.
type MockValidator struct {
	ctrl     *gomock.Controller
	recorder *MockValidatorMockRecorder
}

// MockValidatorMockRecorder is the mock recorder for MockValidator.
type MockValidatorMockRecorder struct {
	mock *MockValidator
}

// NewMockValidator creates a new mock instance.
func NewMockValidator(ctrl *gomock.Controller) *MockValidator {
	mock := &MockValidator{ctrl: ctrl}
	mock.recorder = &MockValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockValidator) EXPECT() *MockValidatorMockRecorder {
	return m.recorder
}

// Struct mocks base method.
func (m *MockValidator) Struct(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Struct", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Struct indicates an expected call of Struct.
func (mr *MockValidatorMockRecorder) Struct(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Struct", reflect.TypeOf((*MockValidator)(nil).Struct), arg0)
}

// StructCtx mocks base method.
func (m *MockValidator) StructCtx(ctx context.Context, s interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StructCtx", ctx, s)
	ret0, _ := ret[0].(error)
	return ret0
}

// StructCtx indicates an expected call of StructCtx.
func (mr *MockValidatorMockRecorder) StructCtx(ctx, s interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StructCtx", reflect.TypeOf((*MockValidator)(nil).StructCtx), ctx, s)
}
