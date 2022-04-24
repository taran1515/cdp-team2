// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/swiggy-2022-bootcamp/cdp-team2/Shipping/dao (interfaces: IDao)

// Package mock_dao is a generated GoMock package.
package mock_dao

import (
	gomock "github.com/golang/mock/gomock"
	models "github.com/swiggy-2022-bootcamp/cdp-team2/Shipping/dao/models"
	reflect "reflect"
)

// MockIDao is a mock of IDao interface
type MockIDao struct {
	ctrl     *gomock.Controller
	recorder *MockIDaoMockRecorder
}

// MockIDaoMockRecorder is the mock recorder for MockIDao
type MockIDaoMockRecorder struct {
	mock *MockIDao
}

// NewMockIDao creates a new mock instance
func NewMockIDao(ctrl *gomock.Controller) *MockIDao {
	mock := &MockIDao{ctrl: ctrl}
	mock.recorder = &MockIDaoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIDao) EXPECT() *MockIDaoMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockIDao) Create(arg0 models.ShippingAddress) (*models.ShippingAddress, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(*models.ShippingAddress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockIDaoMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIDao)(nil).Create), arg0)
}

// GetByCustomerId mocks base method
func (m *MockIDao) GetByCustomerId(arg0 int) ([]models.ShippingAddress, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByCustomerId", arg0)
	ret0, _ := ret[0].([]models.ShippingAddress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByCustomerId indicates an expected call of GetByCustomerId
func (mr *MockIDaoMockRecorder) GetByCustomerId(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByCustomerId", reflect.TypeOf((*MockIDao)(nil).GetByCustomerId), arg0)
}
