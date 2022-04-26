// Code generated by mockery v2.10.1. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	errors "github.com/swiggy-2022-bootcamp/cdp-team2/cart/internal/errors"

	models "github.com/swiggy-2022-bootcamp/cdp-team2/cart/internal/dao/models"
)

// DynamoDAO is an autogenerated mock type for the DynamoDAO type
type DynamoDAO struct {
	mock.Mock
}

// AddCartItem provides a mock function with given fields: product
func (_m *DynamoDAO) AddCartItem(product models.Product) *errors.ServerError {
	ret := _m.Called(product)

	var r0 *errors.ServerError
	if rf, ok := ret.Get(0).(func(models.Product) *errors.ServerError); ok {
		r0 = rf(product)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*errors.ServerError)
		}
	}

	return r0
}

// DeleteCartItem provides a mock function with given fields: customerId, productId
func (_m *DynamoDAO) DeleteCartItem(customerId string, productId string) *errors.ServerError {
	ret := _m.Called(customerId, productId)

	var r0 *errors.ServerError
	if rf, ok := ret.Get(0).(func(string, string) *errors.ServerError); ok {
		r0 = rf(customerId, productId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*errors.ServerError)
		}
	}

	return r0
}

// EmptyCart provides a mock function with given fields: customerId
func (_m *DynamoDAO) EmptyCart(customerId string) *errors.ServerError {
	ret := _m.Called(customerId)

	var r0 *errors.ServerError
	if rf, ok := ret.Get(0).(func(string) *errors.ServerError); ok {
		r0 = rf(customerId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*errors.ServerError)
		}
	}

	return r0
}

// GetCart provides a mock function with given fields: customerId
func (_m *DynamoDAO) GetCart(customerId string) (models.Cart, *errors.ServerError) {
	ret := _m.Called(customerId)

	var r0 models.Cart
	if rf, ok := ret.Get(0).(func(string) models.Cart); ok {
		r0 = rf(customerId)
	} else {
		r0 = ret.Get(0).(models.Cart)
	}

	var r1 *errors.ServerError
	if rf, ok := ret.Get(1).(func(string) *errors.ServerError); ok {
		r1 = rf(customerId)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errors.ServerError)
		}
	}

	return r0, r1
}

// UpdateCartItem provides a mock function with given fields: customerId, productId, newQuantity
func (_m *DynamoDAO) UpdateCartItem(customerId string, productId string, newQuantity int32) *errors.ServerError {
	ret := _m.Called(customerId, productId, newQuantity)

	var r0 *errors.ServerError
	if rf, ok := ret.Get(0).(func(string, string, int32) *errors.ServerError); ok {
		r0 = rf(customerId, productId, newQuantity)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*errors.ServerError)
		}
	}

	return r0
}