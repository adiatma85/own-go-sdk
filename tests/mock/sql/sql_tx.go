// Code generated by MockGen. DO NOT EDIT.
// Source: ./sql/sql_tx.go
//
// Generated by this command:
//
//	mockgen -source ./sql/sql_tx.go -destination ./tests/mock/sql/sql_tx.go
//

// Package mock_sql is a generated GoMock package.
package mock_sql

import (
	sql "database/sql"
	reflect "reflect"

	sql0 "github.com/adiatma85/own-go-sdk/sql"
	sqlx "github.com/jmoiron/sqlx"
	gomock "go.uber.org/mock/gomock"
)

// MockCommandTx is a mock of CommandTx interface.
type MockCommandTx struct {
	ctrl     *gomock.Controller
	recorder *MockCommandTxMockRecorder
}

// MockCommandTxMockRecorder is the mock recorder for MockCommandTx.
type MockCommandTxMockRecorder struct {
	mock *MockCommandTx
}

// NewMockCommandTx creates a new mock instance.
func NewMockCommandTx(ctrl *gomock.Controller) *MockCommandTx {
	mock := &MockCommandTx{ctrl: ctrl}
	mock.recorder = &MockCommandTxMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCommandTx) EXPECT() *MockCommandTxMockRecorder {
	return m.recorder
}

// Commit mocks base method.
func (m *MockCommandTx) Commit() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit")
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit.
func (mr *MockCommandTxMockRecorder) Commit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockCommandTx)(nil).Commit))
}

// Exec mocks base method.
func (m *MockCommandTx) Exec(name, query string, args ...any) (sql.Result, error) {
	m.ctrl.T.Helper()
	varargs := []any{name, query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Exec", varargs...)
	ret0, _ := ret[0].(sql.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exec indicates an expected call of Exec.
func (mr *MockCommandTxMockRecorder) Exec(name, query any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{name, query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exec", reflect.TypeOf((*MockCommandTx)(nil).Exec), varargs...)
}

// Get mocks base method.
func (m *MockCommandTx) Get(name, query string, dest any, args ...any) error {
	m.ctrl.T.Helper()
	varargs := []any{name, query, dest}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockCommandTxMockRecorder) Get(name, query, dest any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{name, query, dest}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCommandTx)(nil).Get), varargs...)
}

// NamedExec mocks base method.
func (m *MockCommandTx) NamedExec(name, query string, args any) (sql.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NamedExec", name, query, args)
	ret0, _ := ret[0].(sql.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NamedExec indicates an expected call of NamedExec.
func (mr *MockCommandTxMockRecorder) NamedExec(name, query, args any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NamedExec", reflect.TypeOf((*MockCommandTx)(nil).NamedExec), name, query, args)
}

// Prepare mocks base method.
func (m *MockCommandTx) Prepare(name, query string) (sql0.CommandStmt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Prepare", name, query)
	ret0, _ := ret[0].(sql0.CommandStmt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Prepare indicates an expected call of Prepare.
func (mr *MockCommandTxMockRecorder) Prepare(name, query any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Prepare", reflect.TypeOf((*MockCommandTx)(nil).Prepare), name, query)
}

// Query mocks base method.
func (m *MockCommandTx) Query(name, query string, args ...any) (*sqlx.Rows, error) {
	m.ctrl.T.Helper()
	varargs := []any{name, query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Query", varargs...)
	ret0, _ := ret[0].(*sqlx.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Query indicates an expected call of Query.
func (mr *MockCommandTxMockRecorder) Query(name, query any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{name, query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockCommandTx)(nil).Query), varargs...)
}

// QueryRow mocks base method.
func (m *MockCommandTx) QueryRow(name, query string, args ...any) (*sqlx.Row, error) {
	m.ctrl.T.Helper()
	varargs := []any{name, query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryRow", varargs...)
	ret0, _ := ret[0].(*sqlx.Row)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryRow indicates an expected call of QueryRow.
func (mr *MockCommandTxMockRecorder) QueryRow(name, query any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{name, query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryRow", reflect.TypeOf((*MockCommandTx)(nil).QueryRow), varargs...)
}

// Rebind mocks base method.
func (m *MockCommandTx) Rebind(query string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rebind", query)
	ret0, _ := ret[0].(string)
	return ret0
}

// Rebind indicates an expected call of Rebind.
func (mr *MockCommandTxMockRecorder) Rebind(query any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rebind", reflect.TypeOf((*MockCommandTx)(nil).Rebind), query)
}

// Rollback mocks base method.
func (m *MockCommandTx) Rollback() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Rollback")
}

// Rollback indicates an expected call of Rollback.
func (mr *MockCommandTxMockRecorder) Rollback() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rollback", reflect.TypeOf((*MockCommandTx)(nil).Rollback))
}

// Select mocks base method.
func (m *MockCommandTx) Select(name, query string, dest any, args ...any) error {
	m.ctrl.T.Helper()
	varargs := []any{name, query, dest}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Select", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Select indicates an expected call of Select.
func (mr *MockCommandTxMockRecorder) Select(name, query, dest any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{name, query, dest}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Select", reflect.TypeOf((*MockCommandTx)(nil).Select), varargs...)
}

// Stmt mocks base method.
func (m *MockCommandTx) Stmt(name string, stmt *sqlx.Stmt) sql0.CommandStmt {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stmt", name, stmt)
	ret0, _ := ret[0].(sql0.CommandStmt)
	return ret0
}

// Stmt indicates an expected call of Stmt.
func (mr *MockCommandTxMockRecorder) Stmt(name, stmt any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stmt", reflect.TypeOf((*MockCommandTx)(nil).Stmt), name, stmt)
}
