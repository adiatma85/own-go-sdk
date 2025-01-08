// Code generated by MockGen. DO NOT EDIT.
// Source: ./jwtAuth/jwt.go
//
// Generated by this command:
//
//	mockgen -source ./jwtAuth/jwt.go -destination ./tests/mock/jwtAuth/jwt.go
//

// Package mock_jwtAuth is a generated GoMock package.
package mock_jwtAuth

import (
	context "context"
	reflect "reflect"

	jwtAuth "github.com/adiatma85/own-go-sdk/jwtAuth"
	gomock "go.uber.org/mock/gomock"
)

// MockInterface is a mock of Interface interface.
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface.
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance.
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// CreateAccessToken mocks base method.
func (m *MockInterface) CreateAccessToken(user jwtAuth.User) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccessToken", user)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccessToken indicates an expected call of CreateAccessToken.
func (mr *MockInterfaceMockRecorder) CreateAccessToken(user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccessToken", reflect.TypeOf((*MockInterface)(nil).CreateAccessToken), user)
}

// CreateRefreshToken mocks base method.
func (m *MockInterface) CreateRefreshToken(user jwtAuth.User) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRefreshToken", user)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRefreshToken indicates an expected call of CreateRefreshToken.
func (mr *MockInterfaceMockRecorder) CreateRefreshToken(user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRefreshToken", reflect.TypeOf((*MockInterface)(nil).CreateRefreshToken), user)
}

// GetUserAuthInfo mocks base method.
func (m *MockInterface) GetUserAuthInfo(ctx context.Context) (jwtAuth.UserAuthInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserAuthInfo", ctx)
	ret0, _ := ret[0].(jwtAuth.UserAuthInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserAuthInfo indicates an expected call of GetUserAuthInfo.
func (mr *MockInterfaceMockRecorder) GetUserAuthInfo(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserAuthInfo", reflect.TypeOf((*MockInterface)(nil).GetUserAuthInfo), ctx)
}

// SetUserAuthInfo mocks base method.
func (m *MockInterface) SetUserAuthInfo(ctx context.Context, param jwtAuth.UserAuthParam) context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetUserAuthInfo", ctx, param)
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// SetUserAuthInfo indicates an expected call of SetUserAuthInfo.
func (mr *MockInterfaceMockRecorder) SetUserAuthInfo(ctx, param any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUserAuthInfo", reflect.TypeOf((*MockInterface)(nil).SetUserAuthInfo), ctx, param)
}

// ValidateAccessToken mocks base method.
func (m *MockInterface) ValidateAccessToken(token string) (jwtAuth.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateAccessToken", token)
	ret0, _ := ret[0].(jwtAuth.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateAccessToken indicates an expected call of ValidateAccessToken.
func (mr *MockInterfaceMockRecorder) ValidateAccessToken(token any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateAccessToken", reflect.TypeOf((*MockInterface)(nil).ValidateAccessToken), token)
}

// ValidateRefreshToken mocks base method.
func (m *MockInterface) ValidateRefreshToken(token string) (jwtAuth.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateRefreshToken", token)
	ret0, _ := ret[0].(jwtAuth.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateRefreshToken indicates an expected call of ValidateRefreshToken.
func (mr *MockInterfaceMockRecorder) ValidateRefreshToken(token any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateRefreshToken", reflect.TypeOf((*MockInterface)(nil).ValidateRefreshToken), token)
}