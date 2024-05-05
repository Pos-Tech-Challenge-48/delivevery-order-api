// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/interfaces/usecases/ordergetter/order_getter.go

// Package ordergetterymock is a generated GoMock package.
package ordergetterymock

import (
	context "context"
	reflect "reflect"

	entities "github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/entities"
	gomock "go.uber.org/mock/gomock"
)

// MockOrderGetter is a mock of OrderGetter interface.
type MockOrderGetter struct {
	ctrl     *gomock.Controller
	recorder *MockOrderGetterMockRecorder
}

// MockOrderGetterMockRecorder is the mock recorder for MockOrderGetter.
type MockOrderGetterMockRecorder struct {
	mock *MockOrderGetter
}

// NewMockOrderGetter creates a new mock instance.
func NewMockOrderGetter(ctrl *gomock.Controller) *MockOrderGetter {
	mock := &MockOrderGetter{ctrl: ctrl}
	mock.recorder = &MockOrderGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderGetter) EXPECT() *MockOrderGetterMockRecorder {
	return m.recorder
}

// GetAll mocks base method.
func (m *MockOrderGetter) GetAll(ctx context.Context, sortBy string) ([]entities.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx, sortBy)
	ret0, _ := ret[0].([]entities.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockOrderGetterMockRecorder) GetAll(ctx, sortBy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockOrderGetter)(nil).GetAll), ctx, sortBy)
}
