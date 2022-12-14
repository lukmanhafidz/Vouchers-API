// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "Test/domain"

	mock "github.com/stretchr/testify/mock"
)

// TransactionData is an autogenerated mock type for the TransactionData type
type TransactionData struct {
	mock.Mock
}

// CreateRedeemData provides a mock function with given fields: newTrans
func (_m *TransactionData) CreateRedeemData(newTrans domain.Transaction) domain.Transaction {
	ret := _m.Called(newTrans)

	var r0 domain.Transaction
	if rf, ok := ret.Get(0).(func(domain.Transaction) domain.Transaction); ok {
		r0 = rf(newTrans)
	} else {
		r0 = ret.Get(0).(domain.Transaction)
	}

	return r0
}

// GetRedeemData provides a mock function with given fields: id
func (_m *TransactionData) GetRedeemData(id int) domain.Transaction_Junction {
	ret := _m.Called(id)

	var r0 domain.Transaction_Junction
	if rf, ok := ret.Get(0).(func(int) domain.Transaction_Junction); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(domain.Transaction_Junction)
	}

	return r0
}

// GetVoucherData provides a mock function with given fields: id
func (_m *TransactionData) GetVoucherData(id int) int {
	ret := _m.Called(id)

	var r0 int
	if rf, ok := ret.Get(0).(func(int) int); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

type mockConstructorTestingTNewTransactionData interface {
	mock.TestingT
	Cleanup(func())
}

// NewTransactionData creates a new instance of TransactionData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTransactionData(t mockConstructorTestingTNewTransactionData) *TransactionData {
	mock := &TransactionData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
