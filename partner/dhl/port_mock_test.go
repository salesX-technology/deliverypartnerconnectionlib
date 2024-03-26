// Code generated by MockGen. DO NOT EDIT.
// Source: port.go
//
// Generated by this command:
//
//	mockgen -source=port.go -package=dhl -destination=port_mock_test.go
//

// Package dhl is a generated GoMock package.
package dhl

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockAuthenticator is a mock of Authenticator interface.
type MockAuthenticator struct {
	ctrl     *gomock.Controller
	recorder *MockAuthenticatorMockRecorder
}

// MockAuthenticatorMockRecorder is the mock recorder for MockAuthenticator.
type MockAuthenticatorMockRecorder struct {
	mock *MockAuthenticator
}

// NewMockAuthenticator creates a new mock instance.
func NewMockAuthenticator(ctrl *gomock.Controller) *MockAuthenticator {
	mock := &MockAuthenticator{ctrl: ctrl}
	mock.recorder = &MockAuthenticatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthenticator) EXPECT() *MockAuthenticatorMockRecorder {
	return m.recorder
}

// Authenticate mocks base method.
func (m *MockAuthenticator) Authenticate() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authenticate")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Authenticate indicates an expected call of Authenticate.
func (mr *MockAuthenticatorMockRecorder) Authenticate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authenticate", reflect.TypeOf((*MockAuthenticator)(nil).Authenticate))
}

// MockDHLOrderCreatorAPI is a mock of DHLOrderCreatorAPI interface.
type MockDHLOrderCreatorAPI struct {
	ctrl     *gomock.Controller
	recorder *MockDHLOrderCreatorAPIMockRecorder
}

// MockDHLOrderCreatorAPIMockRecorder is the mock recorder for MockDHLOrderCreatorAPI.
type MockDHLOrderCreatorAPIMockRecorder struct {
	mock *MockDHLOrderCreatorAPI
}

// NewMockDHLOrderCreatorAPI creates a new mock instance.
func NewMockDHLOrderCreatorAPI(ctrl *gomock.Controller) *MockDHLOrderCreatorAPI {
	mock := &MockDHLOrderCreatorAPI{ctrl: ctrl}
	mock.recorder = &MockDHLOrderCreatorAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDHLOrderCreatorAPI) EXPECT() *MockDHLOrderCreatorAPIMockRecorder {
	return m.recorder
}

// Post mocks base method.
func (m *MockDHLOrderCreatorAPI) Post(endpoint string, headers map[string]string, request DHLCreateOrderAPIRequest) (DHLCreateOrderAPIResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Post", endpoint, headers, request)
	ret0, _ := ret[0].(DHLCreateOrderAPIResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Post indicates an expected call of Post.
func (mr *MockDHLOrderCreatorAPIMockRecorder) Post(endpoint, headers, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Post", reflect.TypeOf((*MockDHLOrderCreatorAPI)(nil).Post), endpoint, headers, request)
}

// MockDHLAuthenticationAPI is a mock of DHLAuthenticationAPI interface.
type MockDHLAuthenticationAPI struct {
	ctrl     *gomock.Controller
	recorder *MockDHLAuthenticationAPIMockRecorder
}

// MockDHLAuthenticationAPIMockRecorder is the mock recorder for MockDHLAuthenticationAPI.
type MockDHLAuthenticationAPIMockRecorder struct {
	mock *MockDHLAuthenticationAPI
}

// NewMockDHLAuthenticationAPI creates a new mock instance.
func NewMockDHLAuthenticationAPI(ctrl *gomock.Controller) *MockDHLAuthenticationAPI {
	mock := &MockDHLAuthenticationAPI{ctrl: ctrl}
	mock.recorder = &MockDHLAuthenticationAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDHLAuthenticationAPI) EXPECT() *MockDHLAuthenticationAPIMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockDHLAuthenticationAPI) Get(endpoint string, headers map[string]string, request DHLAuthenticationAPIRequest) (DHLAuthenticationAPIResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", endpoint, headers, request)
	ret0, _ := ret[0].(DHLAuthenticationAPIResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockDHLAuthenticationAPIMockRecorder) Get(endpoint, headers, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDHLAuthenticationAPI)(nil).Get), endpoint, headers, request)
}

// MockDHLOrderDeletorAPI is a mock of DHLOrderDeletorAPI interface.
type MockDHLOrderDeletorAPI struct {
	ctrl     *gomock.Controller
	recorder *MockDHLOrderDeletorAPIMockRecorder
}

// MockDHLOrderDeletorAPIMockRecorder is the mock recorder for MockDHLOrderDeletorAPI.
type MockDHLOrderDeletorAPIMockRecorder struct {
	mock *MockDHLOrderDeletorAPI
}

// NewMockDHLOrderDeletorAPI creates a new mock instance.
func NewMockDHLOrderDeletorAPI(ctrl *gomock.Controller) *MockDHLOrderDeletorAPI {
	mock := &MockDHLOrderDeletorAPI{ctrl: ctrl}
	mock.recorder = &MockDHLOrderDeletorAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDHLOrderDeletorAPI) EXPECT() *MockDHLOrderDeletorAPIMockRecorder {
	return m.recorder
}

// Post mocks base method.
func (m *MockDHLOrderDeletorAPI) Post(endpoint string, headers map[string]string, request DHLDeleteOrderAPIRequest) (DHLDeleteOrderAPIResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Post", endpoint, headers, request)
	ret0, _ := ret[0].(DHLDeleteOrderAPIResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Post indicates an expected call of Post.
func (mr *MockDHLOrderDeletorAPIMockRecorder) Post(endpoint, headers, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Post", reflect.TypeOf((*MockDHLOrderDeletorAPI)(nil).Post), endpoint, headers, request)
}

// MockDHLOrderUpdatorAPI is a mock of DHLOrderUpdatorAPI interface.
type MockDHLOrderUpdatorAPI struct {
	ctrl     *gomock.Controller
	recorder *MockDHLOrderUpdatorAPIMockRecorder
}

// MockDHLOrderUpdatorAPIMockRecorder is the mock recorder for MockDHLOrderUpdatorAPI.
type MockDHLOrderUpdatorAPIMockRecorder struct {
	mock *MockDHLOrderUpdatorAPI
}

// NewMockDHLOrderUpdatorAPI creates a new mock instance.
func NewMockDHLOrderUpdatorAPI(ctrl *gomock.Controller) *MockDHLOrderUpdatorAPI {
	mock := &MockDHLOrderUpdatorAPI{ctrl: ctrl}
	mock.recorder = &MockDHLOrderUpdatorAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDHLOrderUpdatorAPI) EXPECT() *MockDHLOrderUpdatorAPIMockRecorder {
	return m.recorder
}

// Post mocks base method.
func (m *MockDHLOrderUpdatorAPI) Post(endpoint string, headers map[string]string, request DHLUpdateOrderAPIRequest) (DHLUpdateOrderAPIResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Post", endpoint, headers, request)
	ret0, _ := ret[0].(DHLUpdateOrderAPIResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Post indicates an expected call of Post.
func (mr *MockDHLOrderUpdatorAPIMockRecorder) Post(endpoint, headers, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Post", reflect.TypeOf((*MockDHLOrderUpdatorAPI)(nil).Post), endpoint, headers, request)
}
