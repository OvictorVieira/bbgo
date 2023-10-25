// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/OvictorVieira/bbgo/pkg/bbgo (interfaces: OrderExecutorExtended)

// Package mocks is a generated GoMock package.
package mocks

import (
	"context"
	"reflect"

	"github.com/golang/mock/gomock"

	"github.com/OvictorVieira/bbgo/pkg/core"
	"github.com/OvictorVieira/bbgo/pkg/types"
)

// MockOrderExecutorExtended is a mock of OrderExecutorExtended interface.
type MockOrderExecutorExtended struct {
	ctrl     *gomock.Controller
	recorder *MockOrderExecutorExtendedMockRecorder
}

// MockOrderExecutorExtendedMockRecorder is the mock recorder for MockOrderExecutorExtended.
type MockOrderExecutorExtendedMockRecorder struct {
	mock *MockOrderExecutorExtended
}

// NewMockOrderExecutorExtended creates a new mock instance.
func NewMockOrderExecutorExtended(ctrl *gomock.Controller) *MockOrderExecutorExtended {
	mock := &MockOrderExecutorExtended{ctrl: ctrl}
	mock.recorder = &MockOrderExecutorExtendedMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderExecutorExtended) EXPECT() *MockOrderExecutorExtendedMockRecorder {
	return m.recorder
}

// CancelOrders mocks base method.
func (m *MockOrderExecutorExtended) CancelOrders(arg0 context.Context, arg1 ...types.Order) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CancelOrders", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CancelOrders indicates an expected call of CancelOrders.
func (mr *MockOrderExecutorExtendedMockRecorder) CancelOrders(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelOrders", reflect.TypeOf((*MockOrderExecutorExtended)(nil).CancelOrders), varargs...)
}

// Position mocks base method.
func (m *MockOrderExecutorExtended) Position() *types.Position {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Position")
	ret0, _ := ret[0].(*types.Position)
	return ret0
}

// Position indicates an expected call of Position.
func (mr *MockOrderExecutorExtendedMockRecorder) Position() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Position", reflect.TypeOf((*MockOrderExecutorExtended)(nil).Position))
}

// SubmitOrders mocks base method.
func (m *MockOrderExecutorExtended) SubmitOrders(arg0 context.Context, arg1 ...types.SubmitOrder) (types.OrderSlice, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SubmitOrders", varargs...)
	ret0, _ := ret[0].(types.OrderSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubmitOrders indicates an expected call of SubmitOrders.
func (mr *MockOrderExecutorExtendedMockRecorder) SubmitOrders(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitOrders", reflect.TypeOf((*MockOrderExecutorExtended)(nil).SubmitOrders), varargs...)
}

// TradeCollector mocks base method.
func (m *MockOrderExecutorExtended) TradeCollector() *core.TradeCollector {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TradeCollector")
	ret0, _ := ret[0].(*core.TradeCollector)
	return ret0
}

// TradeCollector indicates an expected call of TradeCollector.
func (mr *MockOrderExecutorExtendedMockRecorder) TradeCollector() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TradeCollector", reflect.TypeOf((*MockOrderExecutorExtended)(nil).TradeCollector))
}
