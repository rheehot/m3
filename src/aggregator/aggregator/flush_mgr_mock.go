// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/m3db/m3aggregator/aggregator/flush_mgr.go

// Package aggregator is a generated GoMock package.
package aggregator

import (
	"reflect"
	"time"

	"github.com/golang/mock/gomock"
)

// MockFlushManager is a mock of FlushManager interface
type MockFlushManager struct {
	ctrl     *gomock.Controller
	recorder *MockFlushManagerMockRecorder
}

// MockFlushManagerMockRecorder is the mock recorder for MockFlushManager
type MockFlushManagerMockRecorder struct {
	mock *MockFlushManager
}

// NewMockFlushManager creates a new mock instance
func NewMockFlushManager(ctrl *gomock.Controller) *MockFlushManager {
	mock := &MockFlushManager{ctrl: ctrl}
	mock.recorder = &MockFlushManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFlushManager) EXPECT() *MockFlushManagerMockRecorder {
	return m.recorder
}

// Reset mocks base method
func (m *MockFlushManager) Reset() error {
	ret := m.ctrl.Call(m, "Reset")
	ret0, _ := ret[0].(error)
	return ret0
}

// Reset indicates an expected call of Reset
func (mr *MockFlushManagerMockRecorder) Reset() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reset", reflect.TypeOf((*MockFlushManager)(nil).Reset))
}

// Open mocks base method
func (m *MockFlushManager) Open() error {
	ret := m.ctrl.Call(m, "Open")
	ret0, _ := ret[0].(error)
	return ret0
}

// Open indicates an expected call of Open
func (mr *MockFlushManagerMockRecorder) Open() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Open", reflect.TypeOf((*MockFlushManager)(nil).Open))
}

// Status mocks base method
func (m *MockFlushManager) Status() FlushStatus {
	ret := m.ctrl.Call(m, "Status")
	ret0, _ := ret[0].(FlushStatus)
	return ret0
}

// Status indicates an expected call of Status
func (mr *MockFlushManagerMockRecorder) Status() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockFlushManager)(nil).Status))
}

// Register mocks base method
func (m *MockFlushManager) Register(flusher flushingMetricList) error {
	ret := m.ctrl.Call(m, "Register", flusher)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register
func (mr *MockFlushManagerMockRecorder) Register(flusher interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockFlushManager)(nil).Register), flusher)
}

// Unregister mocks base method
func (m *MockFlushManager) Unregister(flusher flushingMetricList) error {
	ret := m.ctrl.Call(m, "Unregister", flusher)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unregister indicates an expected call of Unregister
func (mr *MockFlushManagerMockRecorder) Unregister(flusher interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unregister", reflect.TypeOf((*MockFlushManager)(nil).Unregister), flusher)
}

// Close mocks base method
func (m *MockFlushManager) Close() error {
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockFlushManagerMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockFlushManager)(nil).Close))
}

// MockflushTask is a mock of flushTask interface
type MockflushTask struct {
	ctrl     *gomock.Controller
	recorder *MockflushTaskMockRecorder
}

// MockflushTaskMockRecorder is the mock recorder for MockflushTask
type MockflushTaskMockRecorder struct {
	mock *MockflushTask
}

// NewMockflushTask creates a new mock instance
func NewMockflushTask(ctrl *gomock.Controller) *MockflushTask {
	mock := &MockflushTask{ctrl: ctrl}
	mock.recorder = &MockflushTaskMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockflushTask) EXPECT() *MockflushTaskMockRecorder {
	return m.recorder
}

// Run mocks base method
func (m *MockflushTask) Run() {
	m.ctrl.Call(m, "Run")
}

// Run indicates an expected call of Run
func (mr *MockflushTaskMockRecorder) Run() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockflushTask)(nil).Run))
}

// MockroleBasedFlushManager is a mock of roleBasedFlushManager interface
type MockroleBasedFlushManager struct {
	ctrl     *gomock.Controller
	recorder *MockroleBasedFlushManagerMockRecorder
}

// MockroleBasedFlushManagerMockRecorder is the mock recorder for MockroleBasedFlushManager
type MockroleBasedFlushManagerMockRecorder struct {
	mock *MockroleBasedFlushManager
}

// NewMockroleBasedFlushManager creates a new mock instance
func NewMockroleBasedFlushManager(ctrl *gomock.Controller) *MockroleBasedFlushManager {
	mock := &MockroleBasedFlushManager{ctrl: ctrl}
	mock.recorder = &MockroleBasedFlushManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockroleBasedFlushManager) EXPECT() *MockroleBasedFlushManagerMockRecorder {
	return m.recorder
}

// Open mocks base method
func (m *MockroleBasedFlushManager) Open() {
	m.ctrl.Call(m, "Open")
}

// Open indicates an expected call of Open
func (mr *MockroleBasedFlushManagerMockRecorder) Open() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Open", reflect.TypeOf((*MockroleBasedFlushManager)(nil).Open))
}

// Init mocks base method
func (m *MockroleBasedFlushManager) Init(buckets []*flushBucket) {
	m.ctrl.Call(m, "Init", buckets)
}

// Init indicates an expected call of Init
func (mr *MockroleBasedFlushManagerMockRecorder) Init(buckets interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockroleBasedFlushManager)(nil).Init), buckets)
}

// Prepare mocks base method
func (m *MockroleBasedFlushManager) Prepare(buckets []*flushBucket) (flushTask, time.Duration) {
	ret := m.ctrl.Call(m, "Prepare", buckets)
	ret0, _ := ret[0].(flushTask)
	ret1, _ := ret[1].(time.Duration)
	return ret0, ret1
}

// Prepare indicates an expected call of Prepare
func (mr *MockroleBasedFlushManagerMockRecorder) Prepare(buckets interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Prepare", reflect.TypeOf((*MockroleBasedFlushManager)(nil).Prepare), buckets)
}

// OnBucketAdded mocks base method
func (m *MockroleBasedFlushManager) OnBucketAdded(bucketIdx int, bucket *flushBucket) {
	m.ctrl.Call(m, "OnBucketAdded", bucketIdx, bucket)
}

// OnBucketAdded indicates an expected call of OnBucketAdded
func (mr *MockroleBasedFlushManagerMockRecorder) OnBucketAdded(bucketIdx, bucket interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnBucketAdded", reflect.TypeOf((*MockroleBasedFlushManager)(nil).OnBucketAdded), bucketIdx, bucket)
}

// OnFlusherAdded mocks base method
func (m *MockroleBasedFlushManager) OnFlusherAdded(bucketIdx int, bucket *flushBucket, flusher flushingMetricList) {
	m.ctrl.Call(m, "OnFlusherAdded", bucketIdx, bucket, flusher)
}

// OnFlusherAdded indicates an expected call of OnFlusherAdded
func (mr *MockroleBasedFlushManagerMockRecorder) OnFlusherAdded(bucketIdx, bucket, flusher interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnFlusherAdded", reflect.TypeOf((*MockroleBasedFlushManager)(nil).OnFlusherAdded), bucketIdx, bucket, flusher)
}

// CanLead mocks base method
func (m *MockroleBasedFlushManager) CanLead() bool {
	ret := m.ctrl.Call(m, "CanLead")
	ret0, _ := ret[0].(bool)
	return ret0
}

// CanLead indicates an expected call of CanLead
func (mr *MockroleBasedFlushManagerMockRecorder) CanLead() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanLead", reflect.TypeOf((*MockroleBasedFlushManager)(nil).CanLead))
}

// Close mocks base method
func (m *MockroleBasedFlushManager) Close() {
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockroleBasedFlushManagerMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockroleBasedFlushManager)(nil).Close))
}
