// Code generated by MockGen. DO NOT EDIT.
// Source: ./redis/redis.go
//
// Generated by this command:
//
//	mockgen -source ./redis/redis.go -destination ./tests/mock/redis/redis.go
//

// Package mock_redis is a generated GoMock package.
package mock_redis

import (
	context "context"
	reflect "reflect"
	time "time"

	redislock "github.com/bsm/redislock"
	redis "github.com/go-redis/redis/v8"
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

// Decrement mocks base method.
func (m *MockInterface) Decrement(ctx context.Context, key string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decrement", ctx, key)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Decrement indicates an expected call of Decrement.
func (mr *MockInterfaceMockRecorder) Decrement(ctx, key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decrement", reflect.TypeOf((*MockInterface)(nil).Decrement), ctx, key)
}

// DecrementBy mocks base method.
func (m *MockInterface) DecrementBy(ctx context.Context, key string, decreasingFactor int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DecrementBy", ctx, key, decreasingFactor)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DecrementBy indicates an expected call of DecrementBy.
func (mr *MockInterfaceMockRecorder) DecrementBy(ctx, key, decreasingFactor any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecrementBy", reflect.TypeOf((*MockInterface)(nil).DecrementBy), ctx, key, decreasingFactor)
}

// Del mocks base method.
func (m *MockInterface) Del(ctx context.Context, key string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Del", ctx, key)
	ret0, _ := ret[0].(error)
	return ret0
}

// Del indicates an expected call of Del.
func (mr *MockInterfaceMockRecorder) Del(ctx, key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Del", reflect.TypeOf((*MockInterface)(nil).Del), ctx, key)
}

// FlushAll mocks base method.
func (m *MockInterface) FlushAll(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FlushAll", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// FlushAll indicates an expected call of FlushAll.
func (mr *MockInterfaceMockRecorder) FlushAll(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FlushAll", reflect.TypeOf((*MockInterface)(nil).FlushAll), ctx)
}

// FlushAllAsync mocks base method.
func (m *MockInterface) FlushAllAsync(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FlushAllAsync", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// FlushAllAsync indicates an expected call of FlushAllAsync.
func (mr *MockInterfaceMockRecorder) FlushAllAsync(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FlushAllAsync", reflect.TypeOf((*MockInterface)(nil).FlushAllAsync), ctx)
}

// FlushDB mocks base method.
func (m *MockInterface) FlushDB(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FlushDB", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// FlushDB indicates an expected call of FlushDB.
func (mr *MockInterfaceMockRecorder) FlushDB(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FlushDB", reflect.TypeOf((*MockInterface)(nil).FlushDB), ctx)
}

// FlushDBAsync mocks base method.
func (m *MockInterface) FlushDBAsync(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FlushDBAsync", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// FlushDBAsync indicates an expected call of FlushDBAsync.
func (mr *MockInterfaceMockRecorder) FlushDBAsync(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FlushDBAsync", reflect.TypeOf((*MockInterface)(nil).FlushDBAsync), ctx)
}

// Get mocks base method.
func (m *MockInterface) Get(ctx context.Context, key string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, key)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockInterfaceMockRecorder) Get(ctx, key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockInterface)(nil).Get), ctx, key)
}

// GetClient mocks base method.
func (m *MockInterface) GetClient() *redis.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClient")
	ret0, _ := ret[0].(*redis.Client)
	return ret0
}

// GetClient indicates an expected call of GetClient.
func (mr *MockInterfaceMockRecorder) GetClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClient", reflect.TypeOf((*MockInterface)(nil).GetClient))
}

// Increment mocks base method.
func (m *MockInterface) Increment(ctx context.Context, key string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Increment", ctx, key)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Increment indicates an expected call of Increment.
func (mr *MockInterfaceMockRecorder) Increment(ctx, key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Increment", reflect.TypeOf((*MockInterface)(nil).Increment), ctx, key)
}

// IncrementBy mocks base method.
func (m *MockInterface) IncrementBy(ctx context.Context, key string, increasingFactor int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncrementBy", ctx, key, increasingFactor)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IncrementBy indicates an expected call of IncrementBy.
func (mr *MockInterfaceMockRecorder) IncrementBy(ctx, key, increasingFactor any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncrementBy", reflect.TypeOf((*MockInterface)(nil).IncrementBy), ctx, key, increasingFactor)
}

// Lock mocks base method.
func (m *MockInterface) Lock(ctx context.Context, key string, expTime time.Duration) (*redislock.Lock, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Lock", ctx, key, expTime)
	ret0, _ := ret[0].(*redislock.Lock)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Lock indicates an expected call of Lock.
func (mr *MockInterfaceMockRecorder) Lock(ctx, key, expTime any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Lock", reflect.TypeOf((*MockInterface)(nil).Lock), ctx, key, expTime)
}

// LockRelease mocks base method.
func (m *MockInterface) LockRelease(ctx context.Context, lock *redislock.Lock) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LockRelease", ctx, lock)
	ret0, _ := ret[0].(error)
	return ret0
}

// LockRelease indicates an expected call of LockRelease.
func (mr *MockInterfaceMockRecorder) LockRelease(ctx, lock any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LockRelease", reflect.TypeOf((*MockInterface)(nil).LockRelease), ctx, lock)
}

// Scan mocks base method.
func (m *MockInterface) Scan(ctx context.Context, key string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Scan", ctx, key)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Scan indicates an expected call of Scan.
func (mr *MockInterfaceMockRecorder) Scan(ctx, key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scan", reflect.TypeOf((*MockInterface)(nil).Scan), ctx, key)
}

// SetEX mocks base method.
func (m *MockInterface) SetEX(ctx context.Context, key, val string, expTime time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetEX", ctx, key, val, expTime)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetEX indicates an expected call of SetEX.
func (mr *MockInterfaceMockRecorder) SetEX(ctx, key, val, expTime any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetEX", reflect.TypeOf((*MockInterface)(nil).SetEX), ctx, key, val, expTime)
}
