// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import (
	domain "github.com/odpf/siren/domain"
	mock "github.com/stretchr/testify/mock"
)

// SlackNotifierService is an autogenerated mock type for the SlackNotifierService type
type SlackNotifierService struct {
	mock.Mock
}

// Notify provides a mock function with given fields: _a0
func (_m *SlackNotifierService) Notify(_a0 *domain.SlackMessage) (*domain.SlackMessageSendResponse, error) {
	ret := _m.Called(_a0)

	var r0 *domain.SlackMessageSendResponse
	if rf, ok := ret.Get(0).(func(*domain.SlackMessage) *domain.SlackMessageSendResponse); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.SlackMessageSendResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.SlackMessage) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
