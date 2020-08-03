// Code generated by MockGen. DO NOT EDIT.
// Source: ./usecase/channel_repository.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	gomock "github.com/golang/mock/gomock"
	domain "github.com/kindai-csg/D-Chat/domain"
	reflect "reflect"
)

// MockChannelRepository is a mock of ChannelRepository interface.
type MockChannelRepository struct {
	ctrl     *gomock.Controller
	recorder *MockChannelRepositoryMockRecorder
}

// MockChannelRepositoryMockRecorder is the mock recorder for MockChannelRepository.
type MockChannelRepositoryMockRecorder struct {
	mock *MockChannelRepository
}

// NewMockChannelRepository creates a new mock instance.
func NewMockChannelRepository(ctrl *gomock.Controller) *MockChannelRepository {
	mock := &MockChannelRepository{ctrl: ctrl}
	mock.recorder = &MockChannelRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChannelRepository) EXPECT() *MockChannelRepositoryMockRecorder {
	return m.recorder
}

// StoreChannel mocks base method.
func (m *MockChannelRepository) StoreChannel(arg0 domain.Channel) (domain.Channel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreChannel", arg0)
	ret0, _ := ret[0].(domain.Channel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StoreChannel indicates an expected call of StoreChannel.
func (mr *MockChannelRepositoryMockRecorder) StoreChannel(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreChannel", reflect.TypeOf((*MockChannelRepository)(nil).StoreChannel), arg0)
}

// DeleteChannel mocks base method.
func (m *MockChannelRepository) DeleteChannel(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteChannel", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteChannel indicates an expected call of DeleteChannel.
func (mr *MockChannelRepositoryMockRecorder) DeleteChannel(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteChannel", reflect.TypeOf((*MockChannelRepository)(nil).DeleteChannel), arg0)
}
