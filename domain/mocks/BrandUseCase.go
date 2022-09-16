// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "Test/domain"

	mock "github.com/stretchr/testify/mock"
)

// BrandUseCase is an autogenerated mock type for the BrandUseCase type
type BrandUseCase struct {
	mock.Mock
}

// CreateBrand provides a mock function with given fields: newBrand
func (_m *BrandUseCase) CreateBrand(newBrand domain.Brand) int {
	ret := _m.Called(newBrand)

	var r0 int
	if rf, ok := ret.Get(0).(func(domain.Brand) int); ok {
		r0 = rf(newBrand)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

type mockConstructorTestingTNewBrandUseCase interface {
	mock.TestingT
	Cleanup(func())
}

// NewBrandUseCase creates a new instance of BrandUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBrandUseCase(t mockConstructorTestingTNewBrandUseCase) *BrandUseCase {
	mock := &BrandUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}