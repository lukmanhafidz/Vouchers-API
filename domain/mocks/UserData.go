// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "Test/domain"

	mock "github.com/stretchr/testify/mock"
)

// UserData is an autogenerated mock type for the UserData type
type UserData struct {
	mock.Mock
}

// CheckDuplicate provides a mock function with given fields: newUser
func (_m *UserData) CheckDuplicate(newUser domain.User) bool {
	ret := _m.Called(newUser)

	var r0 bool
	if rf, ok := ret.Get(0).(func(domain.User) bool); ok {
		r0 = rf(newUser)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// GetPasswordData provides a mock function with given fields: name
func (_m *UserData) GetPasswordData(name string) string {
	ret := _m.Called(name)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Login provides a mock function with given fields: userLogin
func (_m *UserData) Login(userLogin domain.User) domain.User {
	ret := _m.Called(userLogin)

	var r0 domain.User
	if rf, ok := ret.Get(0).(func(domain.User) domain.User); ok {
		r0 = rf(userLogin)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	return r0
}

// RegisterData provides a mock function with given fields: newUser
func (_m *UserData) RegisterData(newUser domain.User) domain.User {
	ret := _m.Called(newUser)

	var r0 domain.User
	if rf, ok := ret.Get(0).(func(domain.User) domain.User); ok {
		r0 = rf(newUser)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	return r0
}

type mockConstructorTestingTNewUserData interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserData creates a new instance of UserData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserData(t mockConstructorTestingTNewUserData) *UserData {
	mock := &UserData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
