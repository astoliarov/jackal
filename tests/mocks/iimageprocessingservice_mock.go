// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/astoliarov/jackal/pkg/interfaces (interfaces: IImageProcessingService)

// Package mocks is a generated GoMock package.
package mocks

import (
	interfaces "github.com/astoliarov/jackal/pkg/interfaces"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIImageProcessingService is a mock of IImageProcessingService interface
type MockIImageProcessingService struct {
	ctrl     *gomock.Controller
	recorder *MockIImageProcessingServiceMockRecorder
}

// MockIImageProcessingServiceMockRecorder is the mock recorder for MockIImageProcessingService
type MockIImageProcessingServiceMockRecorder struct {
	mock *MockIImageProcessingService
}

// NewMockIImageProcessingService creates a new mock instance
func NewMockIImageProcessingService(ctrl *gomock.Controller) *MockIImageProcessingService {
	mock := &MockIImageProcessingService{ctrl: ctrl}
	mock.recorder = &MockIImageProcessingServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIImageProcessingService) EXPECT() *MockIImageProcessingServiceMockRecorder {
	return m.recorder
}

// CropCentered mocks base method
func (m *MockIImageProcessingService) CropCentered(arg0 []byte, arg1, arg2 int, arg3 interfaces.CropType) ([]byte, string, error) {
	ret := m.ctrl.Call(m, "CropCentered", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CropCentered indicates an expected call of CropCentered
func (mr *MockIImageProcessingServiceMockRecorder) CropCentered(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CropCentered", reflect.TypeOf((*MockIImageProcessingService)(nil).CropCentered), arg0, arg1, arg2, arg3)
}
