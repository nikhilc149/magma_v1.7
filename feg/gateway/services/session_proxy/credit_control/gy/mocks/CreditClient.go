// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

// Run make gen at FeG to re-generate

package mocks

import (
	time "time"

	mock "github.com/stretchr/testify/mock"

	diameter "magma/feg/gateway/diameter"
	gy "magma/feg/gateway/services/session_proxy/credit_control/gy"
)

// CreditClient is an autogenerated mock type for the CreditClient type
type CreditClient struct {
	mock.Mock
}

// DisableConnections provides a mock function with given fields: period
func (_m *CreditClient) DisableConnections(period time.Duration) {
	_m.Called(period)
}

// EnableConnections provides a mock function with given fields:
func (_m *CreditClient) EnableConnections() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IgnoreAnswer provides a mock function with given fields: request
func (_m *CreditClient) IgnoreAnswer(request *gy.CreditControlRequest) {
	_m.Called(request)
}

// SendCreditControlRequest provides a mock function with given fields: server, done, request
func (_m *CreditClient) SendCreditControlRequest(server *diameter.DiameterServerConfig, done chan interface{}, request *gy.CreditControlRequest) error {
	ret := _m.Called(server, done, request)

	var r0 error
	if rf, ok := ret.Get(0).(func(*diameter.DiameterServerConfig, chan interface{}, *gy.CreditControlRequest) error); ok {
		r0 = rf(server, done, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
