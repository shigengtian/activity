// Code generated by MockGen. DO NOT EDIT.
// Source: delegate_actor.go

// Package pub is a generated GoMock package.
package pub

import (
	context "context"
	vocab "github.com/go-fed/activity/streams/vocab"
	gomock "github.com/golang/mock/gomock"
	http "net/http"
	url "net/url"
	reflect "reflect"
)

// MockDelegateActor is a mock of DelegateActor interface
type MockDelegateActor struct {
	ctrl     *gomock.Controller
	recorder *MockDelegateActorMockRecorder
}

// MockDelegateActorMockRecorder is the mock recorder for MockDelegateActor
type MockDelegateActorMockRecorder struct {
	mock *MockDelegateActor
}

// NewMockDelegateActor creates a new mock instance
func NewMockDelegateActor(ctrl *gomock.Controller) *MockDelegateActor {
	mock := &MockDelegateActor{ctrl: ctrl}
	mock.recorder = &MockDelegateActorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDelegateActor) EXPECT() *MockDelegateActorMockRecorder {
	return m.recorder
}

// PostInboxRequestBodyHook mocks base method
func (m *MockDelegateActor) PostInboxRequestBodyHook(c context.Context, r *http.Request, activity Activity) (context.Context, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostInboxRequestBodyHook", c, r, activity)
	ret0, _ := ret[0].(context.Context)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostInboxRequestBodyHook indicates an expected call of PostInboxRequestBodyHook
func (mr *MockDelegateActorMockRecorder) PostInboxRequestBodyHook(c, r, activity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostInboxRequestBodyHook", reflect.TypeOf((*MockDelegateActor)(nil).PostInboxRequestBodyHook), c, r, activity)
}

// PostOutboxRequestBodyHook mocks base method
func (m *MockDelegateActor) PostOutboxRequestBodyHook(c context.Context, r *http.Request, data vocab.Type) (context.Context, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostOutboxRequestBodyHook", c, r, data)
	ret0, _ := ret[0].(context.Context)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostOutboxRequestBodyHook indicates an expected call of PostOutboxRequestBodyHook
func (mr *MockDelegateActorMockRecorder) PostOutboxRequestBodyHook(c, r, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostOutboxRequestBodyHook", reflect.TypeOf((*MockDelegateActor)(nil).PostOutboxRequestBodyHook), c, r, data)
}

// AuthenticatePostInbox mocks base method
func (m *MockDelegateActor) AuthenticatePostInbox(c context.Context, w http.ResponseWriter, r *http.Request) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthenticatePostInbox", c, w, r)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthenticatePostInbox indicates an expected call of AuthenticatePostInbox
func (mr *MockDelegateActorMockRecorder) AuthenticatePostInbox(c, w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthenticatePostInbox", reflect.TypeOf((*MockDelegateActor)(nil).AuthenticatePostInbox), c, w, r)
}

// AuthenticateGetInbox mocks base method
func (m *MockDelegateActor) AuthenticateGetInbox(c context.Context, w http.ResponseWriter, r *http.Request) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthenticateGetInbox", c, w, r)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthenticateGetInbox indicates an expected call of AuthenticateGetInbox
func (mr *MockDelegateActorMockRecorder) AuthenticateGetInbox(c, w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthenticateGetInbox", reflect.TypeOf((*MockDelegateActor)(nil).AuthenticateGetInbox), c, w, r)
}

// AuthorizePostInbox mocks base method
func (m *MockDelegateActor) AuthorizePostInbox(c context.Context, w http.ResponseWriter, activity Activity) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthorizePostInbox", c, w, activity)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthorizePostInbox indicates an expected call of AuthorizePostInbox
func (mr *MockDelegateActorMockRecorder) AuthorizePostInbox(c, w, activity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthorizePostInbox", reflect.TypeOf((*MockDelegateActor)(nil).AuthorizePostInbox), c, w, activity)
}

// PostInbox mocks base method
func (m *MockDelegateActor) PostInbox(c context.Context, inboxIRI *url.URL, activity Activity) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostInbox", c, inboxIRI, activity)
	ret0, _ := ret[0].(error)
	return ret0
}

// PostInbox indicates an expected call of PostInbox
func (mr *MockDelegateActorMockRecorder) PostInbox(c, inboxIRI, activity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostInbox", reflect.TypeOf((*MockDelegateActor)(nil).PostInbox), c, inboxIRI, activity)
}

// InboxForwarding mocks base method
func (m *MockDelegateActor) InboxForwarding(c context.Context, inboxIRI *url.URL, activity Activity) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InboxForwarding", c, inboxIRI, activity)
	ret0, _ := ret[0].(error)
	return ret0
}

// InboxForwarding indicates an expected call of InboxForwarding
func (mr *MockDelegateActorMockRecorder) InboxForwarding(c, inboxIRI, activity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InboxForwarding", reflect.TypeOf((*MockDelegateActor)(nil).InboxForwarding), c, inboxIRI, activity)
}

// PostOutbox mocks base method
func (m *MockDelegateActor) PostOutbox(c context.Context, a Activity, outboxIRI *url.URL, rawJSON map[string]interface{}) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostOutbox", c, a, outboxIRI, rawJSON)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostOutbox indicates an expected call of PostOutbox
func (mr *MockDelegateActorMockRecorder) PostOutbox(c, a, outboxIRI, rawJSON interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostOutbox", reflect.TypeOf((*MockDelegateActor)(nil).PostOutbox), c, a, outboxIRI, rawJSON)
}

// AddNewIds mocks base method
func (m *MockDelegateActor) AddNewIds(c context.Context, a Activity) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddNewIds", c, a)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddNewIds indicates an expected call of AddNewIds
func (mr *MockDelegateActorMockRecorder) AddNewIds(c, a interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddNewIds", reflect.TypeOf((*MockDelegateActor)(nil).AddNewIds), c, a)
}

// Deliver mocks base method
func (m *MockDelegateActor) Deliver(c context.Context, outbox *url.URL, activity Activity) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deliver", c, outbox, activity)
	ret0, _ := ret[0].(error)
	return ret0
}

// Deliver indicates an expected call of Deliver
func (mr *MockDelegateActorMockRecorder) Deliver(c, outbox, activity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deliver", reflect.TypeOf((*MockDelegateActor)(nil).Deliver), c, outbox, activity)
}

// AuthenticatePostOutbox mocks base method
func (m *MockDelegateActor) AuthenticatePostOutbox(c context.Context, w http.ResponseWriter, r *http.Request) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthenticatePostOutbox", c, w, r)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthenticatePostOutbox indicates an expected call of AuthenticatePostOutbox
func (mr *MockDelegateActorMockRecorder) AuthenticatePostOutbox(c, w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthenticatePostOutbox", reflect.TypeOf((*MockDelegateActor)(nil).AuthenticatePostOutbox), c, w, r)
}

// AuthenticateGetOutbox mocks base method
func (m *MockDelegateActor) AuthenticateGetOutbox(c context.Context, w http.ResponseWriter, r *http.Request) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthenticateGetOutbox", c, w, r)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthenticateGetOutbox indicates an expected call of AuthenticateGetOutbox
func (mr *MockDelegateActorMockRecorder) AuthenticateGetOutbox(c, w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthenticateGetOutbox", reflect.TypeOf((*MockDelegateActor)(nil).AuthenticateGetOutbox), c, w, r)
}

// WrapInCreate mocks base method
func (m *MockDelegateActor) WrapInCreate(c context.Context, value vocab.Type, outboxIRI *url.URL) (vocab.ActivityStreamsCreate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WrapInCreate", c, value, outboxIRI)
	ret0, _ := ret[0].(vocab.ActivityStreamsCreate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WrapInCreate indicates an expected call of WrapInCreate
func (mr *MockDelegateActorMockRecorder) WrapInCreate(c, value, outboxIRI interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WrapInCreate", reflect.TypeOf((*MockDelegateActor)(nil).WrapInCreate), c, value, outboxIRI)
}

// GetOutbox mocks base method
func (m *MockDelegateActor) GetOutbox(c context.Context, r *http.Request) (vocab.ActivityStreamsOrderedCollectionPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOutbox", c, r)
	ret0, _ := ret[0].(vocab.ActivityStreamsOrderedCollectionPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOutbox indicates an expected call of GetOutbox
func (mr *MockDelegateActorMockRecorder) GetOutbox(c, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOutbox", reflect.TypeOf((*MockDelegateActor)(nil).GetOutbox), c, r)
}

// GetInbox mocks base method
func (m *MockDelegateActor) GetInbox(c context.Context, r *http.Request) (vocab.ActivityStreamsOrderedCollectionPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInbox", c, r)
	ret0, _ := ret[0].(vocab.ActivityStreamsOrderedCollectionPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInbox indicates an expected call of GetInbox
func (mr *MockDelegateActorMockRecorder) GetInbox(c, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInbox", reflect.TypeOf((*MockDelegateActor)(nil).GetInbox), c, r)
}
