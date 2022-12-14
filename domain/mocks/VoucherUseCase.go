// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "Test/domain"

	mock "github.com/stretchr/testify/mock"
)

// VoucherUseCase is an autogenerated mock type for the VoucherUseCase type
type VoucherUseCase struct {
	mock.Mock
}

// CreateVoucher provides a mock function with given fields: newVoucher
func (_m *VoucherUseCase) CreateVoucher(newVoucher domain.Voucher) int {
	ret := _m.Called(newVoucher)

	var r0 int
	if rf, ok := ret.Get(0).(func(domain.Voucher) int); ok {
		r0 = rf(newVoucher)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// GetAllVoucher provides a mock function with given fields: brandID
func (_m *VoucherUseCase) GetAllVoucher(brandID int) ([]map[string]interface{}, int) {
	ret := _m.Called(brandID)

	var r0 []map[string]interface{}
	if rf, ok := ret.Get(0).(func(int) []map[string]interface{}); ok {
		r0 = rf(brandID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]map[string]interface{})
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(int) int); ok {
		r1 = rf(brandID)
	} else {
		r1 = ret.Get(1).(int)
	}

	return r0, r1
}

// GetByIDVoucher provides a mock function with given fields: id
func (_m *VoucherUseCase) GetByIDVoucher(id int) (map[string]interface{}, int) {
	ret := _m.Called(id)

	var r0 map[string]interface{}
	if rf, ok := ret.Get(0).(func(int) map[string]interface{}); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(int) int); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Get(1).(int)
	}

	return r0, r1
}

type mockConstructorTestingTNewVoucherUseCase interface {
	mock.TestingT
	Cleanup(func())
}

// NewVoucherUseCase creates a new instance of VoucherUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewVoucherUseCase(t mockConstructorTestingTNewVoucherUseCase) *VoucherUseCase {
	mock := &VoucherUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
