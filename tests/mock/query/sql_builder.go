// Code generated by MockGen. DO NOT EDIT.
// Source: ./query/sql_builder.go
//
// Generated by this command:
//
//	mockgen -source ./query/sql_builder.go -destination ./tests/mock/query/sql_builder.go
//

// Package mock_query is a generated GoMock package.
package mock_query

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockCursor is a mock of Cursor interface.
type MockCursor struct {
	ctrl     *gomock.Controller
	recorder *MockCursorMockRecorder
}

// MockCursorMockRecorder is the mock recorder for MockCursor.
type MockCursorMockRecorder struct {
	mock *MockCursor
}

// NewMockCursor creates a new mock instance.
func NewMockCursor(ctrl *gomock.Controller) *MockCursor {
	mock := &MockCursor{ctrl: ctrl}
	mock.recorder = &MockCursorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCursor) EXPECT() *MockCursorMockRecorder {
	return m.recorder
}

// DecodeCursor mocks base method.
func (m *MockCursor) DecodeCursor(v string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DecodeCursor", v)
	ret0, _ := ret[0].(error)
	return ret0
}

// DecodeCursor indicates an expected call of DecodeCursor.
func (mr *MockCursorMockRecorder) DecodeCursor(v any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecodeCursor", reflect.TypeOf((*MockCursor)(nil).DecodeCursor), v)
}

// EncodeCursor mocks base method.
func (m *MockCursor) EncodeCursor() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EncodeCursor")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EncodeCursor indicates an expected call of EncodeCursor.
func (mr *MockCursorMockRecorder) EncodeCursor() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EncodeCursor", reflect.TypeOf((*MockCursor)(nil).EncodeCursor))
}
