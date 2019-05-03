// Code generated by MockGen. DO NOT EDIT.
// Source: social_protocol.go

// Package pub is a generated GoMock package.
package pub

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	http "net/http"
	reflect "reflect"
)

// MockSocialProtocol is a mock of SocialProtocol interface
type MockSocialProtocol struct {
	ctrl     *gomock.Controller
	recorder *MockSocialProtocolMockRecorder
}

// MockSocialProtocolMockRecorder is the mock recorder for MockSocialProtocol
type MockSocialProtocolMockRecorder struct {
	mock *MockSocialProtocol
}

// NewMockSocialProtocol creates a new mock instance
func NewMockSocialProtocol(ctrl *gomock.Controller) *MockSocialProtocol {
	mock := &MockSocialProtocol{ctrl: ctrl}
	mock.recorder = &MockSocialProtocolMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSocialProtocol) EXPECT() *MockSocialProtocolMockRecorder {
	return m.recorder
}

// AuthenticatePostOutbox mocks base method
func (m *MockSocialProtocol) AuthenticatePostOutbox(c context.Context, w http.ResponseWriter, r *http.Request) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthenticatePostOutbox", c, w, r)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthenticatePostOutbox indicates an expected call of AuthenticatePostOutbox
func (mr *MockSocialProtocolMockRecorder) AuthenticatePostOutbox(c, w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthenticatePostOutbox", reflect.TypeOf((*MockSocialProtocol)(nil).AuthenticatePostOutbox), c, w, r)
}

// Callbacks mocks base method
func (m *MockSocialProtocol) Callbacks(c context.Context) (SocialWrappedCallbacks, []interface{}) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Callbacks", c)
	ret0, _ := ret[0].(SocialWrappedCallbacks)
	ret1, _ := ret[1].([]interface{})
	return ret0, ret1
}

// Callbacks indicates an expected call of Callbacks
func (mr *MockSocialProtocolMockRecorder) Callbacks(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Callbacks", reflect.TypeOf((*MockSocialProtocol)(nil).Callbacks), c)
}

// DefaultCallback mocks base method
func (m *MockSocialProtocol) DefaultCallback(c context.Context, activity Activity) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DefaultCallback", c, activity)
	ret0, _ := ret[0].(error)
	return ret0
}

// DefaultCallback indicates an expected call of DefaultCallback
func (mr *MockSocialProtocolMockRecorder) DefaultCallback(c, activity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefaultCallback", reflect.TypeOf((*MockSocialProtocol)(nil).DefaultCallback), c, activity)
}
