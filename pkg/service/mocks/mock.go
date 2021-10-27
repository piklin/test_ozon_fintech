// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/piklin/test_ozon_fintech/models"
)

// MockShortURL is a mock of ShortURL interface.
type MockShortURL struct {
	ctrl     *gomock.Controller
	recorder *MockShortURLMockRecorder
}

// MockShortURLMockRecorder is the mock recorder for MockShortURL.
type MockShortURLMockRecorder struct {
	mock *MockShortURL
}

// NewMockShortURL creates a new mock instance.
func NewMockShortURL(ctrl *gomock.Controller) *MockShortURL {
	mock := &MockShortURL{ctrl: ctrl}
	mock.recorder = &MockShortURLMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockShortURL) EXPECT() *MockShortURLMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockShortURL) Create(url models.URLRequest) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", url)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockShortURLMockRecorder) Create(url interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockShortURL)(nil).Create), url)
}

// GenerateShortURL mocks base method.
func (m *MockShortURL) GenerateShortURL() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateShortURL")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateShortURL indicates an expected call of GenerateShortURL.
func (mr *MockShortURLMockRecorder) GenerateShortURL() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateShortURL", reflect.TypeOf((*MockShortURL)(nil).GenerateShortURL))
}

// GetFullURL mocks base method.
func (m *MockShortURL) GetFullURL(shortURL string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFullURL", shortURL)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFullURL indicates an expected call of GetFullURL.
func (mr *MockShortURLMockRecorder) GetFullURL(shortURL interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFullURL", reflect.TypeOf((*MockShortURL)(nil).GetFullURL), shortURL)
}
