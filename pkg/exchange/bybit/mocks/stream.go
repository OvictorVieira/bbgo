// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/OvictorVieira/bbgo/pkg/exchange/bybit (interfaces: StreamDataProvider)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	bybitapi "github.com/OvictorVieira/bbgo/pkg/exchange/bybit/bybitapi"
	types "github.com/OvictorVieira/bbgo/pkg/types"
	gomock "github.com/golang/mock/gomock"
)

// MockStreamDataProvider is a mock of StreamDataProvider interface.
type MockStreamDataProvider struct {
	ctrl     *gomock.Controller
	recorder *MockStreamDataProviderMockRecorder
}

// MockStreamDataProviderMockRecorder is the mock recorder for MockStreamDataProvider.
type MockStreamDataProviderMockRecorder struct {
	mock *MockStreamDataProvider
}

// NewMockStreamDataProvider creates a new mock instance.
func NewMockStreamDataProvider(ctrl *gomock.Controller) *MockStreamDataProvider {
	mock := &MockStreamDataProvider{ctrl: ctrl}
	mock.recorder = &MockStreamDataProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStreamDataProvider) EXPECT() *MockStreamDataProviderMockRecorder {
	return m.recorder
}

// GetAllFeeRates mocks base method.
func (m *MockStreamDataProvider) GetAllFeeRates(arg0 context.Context) (bybitapi.FeeRates, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllFeeRates", arg0)
	ret0, _ := ret[0].(bybitapi.FeeRates)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllFeeRates indicates an expected call of GetAllFeeRates.
func (mr *MockStreamDataProviderMockRecorder) GetAllFeeRates(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllFeeRates", reflect.TypeOf((*MockStreamDataProvider)(nil).GetAllFeeRates), arg0)
}

// QueryAccountBalances mocks base method.
func (m *MockStreamDataProvider) QueryAccountBalances(arg0 context.Context) (types.BalanceMap, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryAccountBalances", arg0)
	ret0, _ := ret[0].(types.BalanceMap)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryAccountBalances indicates an expected call of QueryAccountBalances.
func (mr *MockStreamDataProviderMockRecorder) QueryAccountBalances(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryAccountBalances", reflect.TypeOf((*MockStreamDataProvider)(nil).QueryAccountBalances), arg0)
}

// QueryMarkets mocks base method.
func (m *MockStreamDataProvider) QueryMarkets(arg0 context.Context) (types.MarketMap, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryMarkets", arg0)
	ret0, _ := ret[0].(types.MarketMap)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryMarkets indicates an expected call of QueryMarkets.
func (mr *MockStreamDataProviderMockRecorder) QueryMarkets(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryMarkets", reflect.TypeOf((*MockStreamDataProvider)(nil).QueryMarkets), arg0)
}
