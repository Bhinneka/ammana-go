// Code generated by mockery v2.4.0. DO NOT EDIT.

package mocks

import (
	ammana "github.com/Bhinneka/ammana-go"
	mock "github.com/stretchr/testify/mock"
)

// AmmanaService is an autogenerated mock type for the AmmanaService type
type AmmanaService struct {
	mock.Mock
}

// GetAuth provides a mock function with given fields:
func (_m *AmmanaService) GetAuth() (ammana.GenerateAccessTokenResponse, error) {
	ret := _m.Called()

	var r0 ammana.GenerateAccessTokenResponse
	if rf, ok := ret.Get(0).(func() ammana.GenerateAccessTokenResponse); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(ammana.GenerateAccessTokenResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCardDetail provides a mock function with given fields: request
func (_m *AmmanaService) GetCardDetail(request ammana.GetCardDetailRequest) (ammana.GetCardDetailResponse, error) {
	ret := _m.Called(request)

	var r0 ammana.GetCardDetailResponse
	if rf, ok := ret.Get(0).(func(ammana.GetCardDetailRequest) ammana.GetCardDetailResponse); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Get(0).(ammana.GetCardDetailResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(ammana.GetCardDetailRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrderURL provides a mock function with given fields: request
func (_m *AmmanaService) GetOrderURL(request ammana.OrderRequest) (ammana.OrderResponse, error) {
	ret := _m.Called(request)

	var r0 ammana.OrderResponse
	if rf, ok := ret.Get(0).(func(ammana.OrderRequest) ammana.OrderResponse); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Get(0).(ammana.OrderResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(ammana.OrderRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPaymentDetail provides a mock function with given fields: invoiceCode
func (_m *AmmanaService) GetPaymentDetail(invoiceCode string) (ammana.GetPaymentDetailResponse, error) {
	ret := _m.Called(invoiceCode)

	var r0 ammana.GetPaymentDetailResponse
	if rf, ok := ret.Get(0).(func(string) ammana.GetPaymentDetailResponse); ok {
		r0 = rf(invoiceCode)
	} else {
		r0 = ret.Get(0).(ammana.GetPaymentDetailResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(invoiceCode)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetSettleOrder provides a mock function with given fields: invoiceCode
func (_m *AmmanaService) SetSettleOrder(invoiceCode string) (ammana.SetSettlePaymentResponse, error) {
	ret := _m.Called(invoiceCode)

	var r0 ammana.SetSettlePaymentResponse
	if rf, ok := ret.Get(0).(func(string) ammana.SetSettlePaymentResponse); ok {
		r0 = rf(invoiceCode)
	} else {
		r0 = ret.Get(0).(ammana.SetSettlePaymentResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(invoiceCode)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetVoidOrder provides a mock function with given fields: invoiceCode
func (_m *AmmanaService) SetVoidOrder(invoiceCode string) (ammana.SetVoidPaymentResponse, error) {
	ret := _m.Called(invoiceCode)

	var r0 ammana.SetVoidPaymentResponse
	if rf, ok := ret.Get(0).(func(string) ammana.SetVoidPaymentResponse); ok {
		r0 = rf(invoiceCode)
	} else {
		r0 = ret.Get(0).(ammana.SetVoidPaymentResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(invoiceCode)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
