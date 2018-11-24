// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/astoliarov/jackal/pkg/interfaces (interfaces: IFetchService)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIFetchService is a mock of IFetchService interface
type MockIFetchService struct {
	ctrl     *gomock.Controller
	recorder *MockIFetchServiceMockRecorder
}

// MockIFetchServiceMockRecorder is the mock recorder for MockIFetchService
type MockIFetchServiceMockRecorder struct {
	mock *MockIFetchService
}

// NewMockIFetchService creates a new mock instance
func NewMockIFetchService(ctrl *gomock.Controller) *MockIFetchService {
	mock := &MockIFetchService{ctrl: ctrl}
	mock.recorder = &MockIFetchServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIFetchService) EXPECT() *MockIFetchServiceMockRecorder {
	return m.recorder
}

// GetBodyFromUrl mocks base method
func (m *MockIFetchService) GetBodyFromUrl(arg0 string) ([]byte, error) {
	ret := m.ctrl.Call(m, "GetBodyFromUrl", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBodyFromUrl indicates an expected call of GetBodyFromUrl
func (mr *MockIFetchServiceMockRecorder) GetBodyFromUrl(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBodyFromUrl", reflect.TypeOf((*MockIFetchService)(nil).GetBodyFromUrl), arg0)
}
