// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/martin-helmich/go-tips/example-mocks (interfaces: CreditScoreProvider)

// Package customer_mocks is a generated GoMock package.
package customer_mocks

import (
	gomock "github.com/golang/mock/gomock"
	example_mocks "github.com/martin-helmich/go-tips/example-mocks"
	reflect "reflect"
)

// MockCreditScoreProvider is a mock of CreditScoreProvider interface
type MockCreditScoreProvider struct {
	ctrl     *gomock.Controller
	recorder *MockCreditScoreProviderMockRecorder
}

// MockCreditScoreProviderMockRecorder is the mock recorder for MockCreditScoreProvider
type MockCreditScoreProviderMockRecorder struct {
	mock *MockCreditScoreProvider
}

// NewMockCreditScoreProvider creates a new mock instance
func NewMockCreditScoreProvider(ctrl *gomock.Controller) *MockCreditScoreProvider {
	mock := &MockCreditScoreProvider{ctrl: ctrl}
	mock.recorder = &MockCreditScoreProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCreditScoreProvider) EXPECT() *MockCreditScoreProviderMockRecorder {
	return m.recorder
}

// CreditScoreForCustomer mocks base method
func (m *MockCreditScoreProvider) CreditScoreForCustomer(arg0 *example_mocks.Customer) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreditScoreForCustomer", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreditScoreForCustomer indicates an expected call of CreditScoreForCustomer
func (mr *MockCreditScoreProviderMockRecorder) CreditScoreForCustomer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreditScoreForCustomer", reflect.TypeOf((*MockCreditScoreProvider)(nil).CreditScoreForCustomer), arg0)
}
