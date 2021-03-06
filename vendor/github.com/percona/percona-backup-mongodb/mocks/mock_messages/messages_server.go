// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/percona/percona-backup-mongodb/proto/messages (interfaces: MessagesServer)

// Package mock_messages is a generated GoMock package.
package mock_messages

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	messages "github.com/percona/percona-backup-mongodb/proto/messages"
	reflect "reflect"
)

// MockMessagesServer is a mock of MessagesServer interface
type MockMessagesServer struct {
	ctrl     *gomock.Controller
	recorder *MockMessagesServerMockRecorder
}

// MockMessagesServerMockRecorder is the mock recorder for MockMessagesServer
type MockMessagesServerMockRecorder struct {
	mock *MockMessagesServer
}

// NewMockMessagesServer creates a new mock instance
func NewMockMessagesServer(ctrl *gomock.Controller) *MockMessagesServer {
	mock := &MockMessagesServer{ctrl: ctrl}
	mock.recorder = &MockMessagesServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMessagesServer) EXPECT() *MockMessagesServerMockRecorder {
	return m.recorder
}

// DBBackupFinished mocks base method
func (m *MockMessagesServer) DBBackupFinished(arg0 context.Context, arg1 *messages.DBBackupFinishStatus) (*messages.DBBackupFinishedAck, error) {
	ret := m.ctrl.Call(m, "DBBackupFinished", arg0, arg1)
	ret0, _ := ret[0].(*messages.DBBackupFinishedAck)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DBBackupFinished indicates an expected call of DBBackupFinished
func (mr *MockMessagesServerMockRecorder) DBBackupFinished(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DBBackupFinished", reflect.TypeOf((*MockMessagesServer)(nil).DBBackupFinished), arg0, arg1)
}

// Logging mocks base method
func (m *MockMessagesServer) Logging(arg0 messages.Messages_LoggingServer) error {
	ret := m.ctrl.Call(m, "Logging", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Logging indicates an expected call of Logging
func (mr *MockMessagesServerMockRecorder) Logging(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logging", reflect.TypeOf((*MockMessagesServer)(nil).Logging), arg0)
}

// MessagesChat mocks base method
func (m *MockMessagesServer) MessagesChat(arg0 messages.Messages_MessagesChatServer) error {
	ret := m.ctrl.Call(m, "MessagesChat", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// MessagesChat indicates an expected call of MessagesChat
func (mr *MockMessagesServerMockRecorder) MessagesChat(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MessagesChat", reflect.TypeOf((*MockMessagesServer)(nil).MessagesChat), arg0)
}

// OplogBackupFinished mocks base method
func (m *MockMessagesServer) OplogBackupFinished(arg0 context.Context, arg1 *messages.OplogBackupFinishStatus) (*messages.OplogBackupFinishedAck, error) {
	ret := m.ctrl.Call(m, "OplogBackupFinished", arg0, arg1)
	ret0, _ := ret[0].(*messages.OplogBackupFinishedAck)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OplogBackupFinished indicates an expected call of OplogBackupFinished
func (mr *MockMessagesServerMockRecorder) OplogBackupFinished(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OplogBackupFinished", reflect.TypeOf((*MockMessagesServer)(nil).OplogBackupFinished), arg0, arg1)
}

// RestoreCompleted mocks base method
func (m *MockMessagesServer) RestoreCompleted(arg0 context.Context, arg1 *messages.RestoreComplete) (*messages.RestoreCompletedAck, error) {
	ret := m.ctrl.Call(m, "RestoreCompleted", arg0, arg1)
	ret0, _ := ret[0].(*messages.RestoreCompletedAck)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RestoreCompleted indicates an expected call of RestoreCompleted
func (mr *MockMessagesServerMockRecorder) RestoreCompleted(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RestoreCompleted", reflect.TypeOf((*MockMessagesServer)(nil).RestoreCompleted), arg0, arg1)
}
