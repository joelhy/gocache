// Code generated by MockGen. DO NOT EDIT.
// Source: store/redis.go

// Package mocks is a generated GoMock package.
package mocks

import (
	v7 "github.com/go-redis/redis/v7"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockRedisClientInterface is a mock of RedisClientInterface interface
type MockRedisClientInterface struct {
	ctrl     *gomock.Controller
	recorder *MockRedisClientInterfaceMockRecorder
}

// MockRedisClientInterfaceMockRecorder is the mock recorder for MockRedisClientInterface
type MockRedisClientInterfaceMockRecorder struct {
	mock *MockRedisClientInterface
}

// NewMockRedisClientInterface creates a new mock instance
func NewMockRedisClientInterface(ctrl *gomock.Controller) *MockRedisClientInterface {
	mock := &MockRedisClientInterface{ctrl: ctrl}
	mock.recorder = &MockRedisClientInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRedisClientInterface) EXPECT() *MockRedisClientInterfaceMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockRedisClientInterface) Get(key string) *v7.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", key)
	ret0, _ := ret[0].(*v7.StringCmd)
	return ret0
}

// Get indicates an expected call of Get
func (mr *MockRedisClientInterfaceMockRecorder) Get(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRedisClientInterface)(nil).Get), key)
}

// Set mocks base method
func (m *MockRedisClientInterface) Set(key string, value interface{}, expiration time.Duration) *v7.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", key, value, expiration)
	ret0, _ := ret[0].(*v7.StatusCmd)
	return ret0
}

// Set indicates an expected call of Set
func (mr *MockRedisClientInterfaceMockRecorder) Set(key, value, expiration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockRedisClientInterface)(nil).Set), key, value, expiration)
}

// Del mocks base method
func (m *MockRedisClientInterface) Del(keys ...string) *v7.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Del", varargs...)
	ret0, _ := ret[0].(*v7.IntCmd)
	return ret0
}

// Del indicates an expected call of Del
func (mr *MockRedisClientInterfaceMockRecorder) Del(keys ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Del", reflect.TypeOf((*MockRedisClientInterface)(nil).Del), keys...)
}

// FlushAll mocks base method
func (m *MockRedisClientInterface) FlushAll() *v7.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FlushAll")
	ret0, _ := ret[0].(*v7.StatusCmd)
	return ret0
}

// FlushAll indicates an expected call of FlushAll
func (mr *MockRedisClientInterfaceMockRecorder) FlushAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FlushAll", reflect.TypeOf((*MockRedisClientInterface)(nil).FlushAll))
}