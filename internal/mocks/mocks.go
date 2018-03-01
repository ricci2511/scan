// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	sql "database/sql"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockRowsScanner is a mock of RowsScanner interface
type MockRowsScanner struct {
	ctrl     *gomock.Controller
	recorder *MockRowsScannerMockRecorder
}

// MockRowsScannerMockRecorder is the mock recorder for MockRowsScanner
type MockRowsScannerMockRecorder struct {
	mock *MockRowsScanner
}

// NewMockRowsScanner creates a new mock instance
func NewMockRowsScanner(ctrl *gomock.Controller) *MockRowsScanner {
	mock := &MockRowsScanner{ctrl: ctrl}
	mock.recorder = &MockRowsScannerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRowsScanner) EXPECT() *MockRowsScannerMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockRowsScanner) Close() error {
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockRowsScannerMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockRowsScanner)(nil).Close))
}

// Scan mocks base method
func (m *MockRowsScanner) Scan(dest ...interface{}) error {
	varargs := []interface{}{}
	for _, a := range dest {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Scan", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Scan indicates an expected call of Scan
func (mr *MockRowsScannerMockRecorder) Scan(dest ...interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scan", reflect.TypeOf((*MockRowsScanner)(nil).Scan), dest...)
}

// Columns mocks base method
func (m *MockRowsScanner) Columns() ([]string, error) {
	ret := m.ctrl.Call(m, "Columns")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Columns indicates an expected call of Columns
func (mr *MockRowsScannerMockRecorder) Columns() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Columns", reflect.TypeOf((*MockRowsScanner)(nil).Columns))
}

// ColumnTypes mocks base method
func (m *MockRowsScanner) ColumnTypes() ([]*sql.ColumnType, error) {
	ret := m.ctrl.Call(m, "ColumnTypes")
	ret0, _ := ret[0].([]*sql.ColumnType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ColumnTypes indicates an expected call of ColumnTypes
func (mr *MockRowsScannerMockRecorder) ColumnTypes() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ColumnTypes", reflect.TypeOf((*MockRowsScanner)(nil).ColumnTypes))
}

// Err mocks base method
func (m *MockRowsScanner) Err() error {
	ret := m.ctrl.Call(m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

// Err indicates an expected call of Err
func (mr *MockRowsScannerMockRecorder) Err() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Err", reflect.TypeOf((*MockRowsScanner)(nil).Err))
}

// Next mocks base method
func (m *MockRowsScanner) Next() bool {
	ret := m.ctrl.Call(m, "Next")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Next indicates an expected call of Next
func (mr *MockRowsScannerMockRecorder) Next() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockRowsScanner)(nil).Next))
}

// Mockcache is a mock of cache interface
type Mockcache struct {
	ctrl     *gomock.Controller
	recorder *MockcacheMockRecorder
}

// MockcacheMockRecorder is the mock recorder for Mockcache
type MockcacheMockRecorder struct {
	mock *Mockcache
}

// NewMockcache creates a new mock instance
func NewMockcache(ctrl *gomock.Controller) *Mockcache {
	mock := &Mockcache{ctrl: ctrl}
	mock.recorder = &MockcacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *Mockcache) EXPECT() *MockcacheMockRecorder {
	return m.recorder
}

// Delete mocks base method
func (m *Mockcache) Delete(key interface{}) {
	m.ctrl.Call(m, "Delete", key)
}

// Delete indicates an expected call of Delete
func (mr *MockcacheMockRecorder) Delete(key interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*Mockcache)(nil).Delete), key)
}

// Load mocks base method
func (m *Mockcache) Load(key interface{}) (interface{}, bool) {
	ret := m.ctrl.Call(m, "Load", key)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// Load indicates an expected call of Load
func (mr *MockcacheMockRecorder) Load(key interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Load", reflect.TypeOf((*Mockcache)(nil).Load), key)
}

// LoadOrStore mocks base method
func (m *Mockcache) LoadOrStore(key, value interface{}) (interface{}, bool) {
	ret := m.ctrl.Call(m, "LoadOrStore", key, value)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// LoadOrStore indicates an expected call of LoadOrStore
func (mr *MockcacheMockRecorder) LoadOrStore(key, value interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadOrStore", reflect.TypeOf((*Mockcache)(nil).LoadOrStore), key, value)
}

// Range mocks base method
func (m *Mockcache) Range(f func(interface{}, interface{}) bool) {
	m.ctrl.Call(m, "Range", f)
}

// Range indicates an expected call of Range
func (mr *MockcacheMockRecorder) Range(f interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Range", reflect.TypeOf((*Mockcache)(nil).Range), f)
}

// Store mocks base method
func (m *Mockcache) Store(key, value interface{}) {
	m.ctrl.Call(m, "Store", key, value)
}

// Store indicates an expected call of Store
func (mr *MockcacheMockRecorder) Store(key, value interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*Mockcache)(nil).Store), key, value)
}
