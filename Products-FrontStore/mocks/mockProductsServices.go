// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/swiggy-2022-bootcamp/cdp-team2/Products-FrontStore/internal/core/ports (interfaces: IProductsServices)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/swiggy-2022-bootcamp/cdp-team2/Products-FrontStore/internal/core/domain"
	errors "github.com/swiggy-2022-bootcamp/cdp-team2/Products-FrontStore/pkg/errors"
)

// MockIProductsServices is a mock of IProductsServices interface.
type MockIProductsServices struct {
	ctrl     *gomock.Controller
	recorder *MockIProductsServicesMockRecorder
}

// MockIProductsServicesMockRecorder is the mock recorder for MockIProductsServices.
type MockIProductsServicesMockRecorder struct {
	mock *MockIProductsServices
}

// NewMockIProductsServices creates a new mock instance.
func NewMockIProductsServices(ctrl *gomock.Controller) *MockIProductsServices {
	mock := &MockIProductsServices{ctrl: ctrl}
	mock.recorder = &MockIProductsServicesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIProductsServices) EXPECT() *MockIProductsServicesMockRecorder {
	return m.recorder
}

// GetProductById mocks base method.
func (m *MockIProductsServices) GetProductById(arg0 int64) (*domain.Product, *errors.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductById", arg0)
	ret0, _ := ret[0].(*domain.Product)
	ret1, _ := ret[1].(*errors.AppError)
	return ret0, ret1
}

// GetProductById indicates an expected call of GetProductById.
func (mr *MockIProductsServicesMockRecorder) GetProductById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductById", reflect.TypeOf((*MockIProductsServices)(nil).GetProductById), arg0)
}

// GetProductList mocks base method.
func (m *MockIProductsServices) GetProductList() ([]*domain.Product, *errors.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductList")
	ret0, _ := ret[0].([]*domain.Product)
	ret1, _ := ret[1].(*errors.AppError)
	return ret0, ret1
}

// GetProductList indicates an expected call of GetProductList.
func (mr *MockIProductsServicesMockRecorder) GetProductList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductList", reflect.TypeOf((*MockIProductsServices)(nil).GetProductList))
}

// GetProductsByCategoryId mocks base method.
func (m *MockIProductsServices) GetProductsByCategoryId(arg0 int64) ([]*domain.Product, *errors.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductsByCategoryId", arg0)
	ret0, _ := ret[0].([]*domain.Product)
	ret1, _ := ret[1].(*errors.AppError)
	return ret0, ret1
}

// GetProductsByCategoryId indicates an expected call of GetProductsByCategoryId.
func (mr *MockIProductsServicesMockRecorder) GetProductsByCategoryId(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductsByCategoryId", reflect.TypeOf((*MockIProductsServices)(nil).GetProductsByCategoryId), arg0)
}