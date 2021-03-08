// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	context "context"

	core "github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/core"
	mock "github.com/stretchr/testify/mock"
)

// TaskReader is an autogenerated mock type for the TaskReader type
type TaskReader struct {
	mock.Mock
}

type TaskReader_Read struct {
	*mock.Call
}

func (_m TaskReader_Read) Return(_a0 *core.TaskTemplate, _a1 error) *TaskReader_Read {
	return &TaskReader_Read{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *TaskReader) OnRead(ctx context.Context) *TaskReader_Read {
	c := _m.On("Read", ctx)
	return &TaskReader_Read{Call: c}
}

func (_m *TaskReader) OnReadMatch(matchers ...interface{}) *TaskReader_Read {
	c := _m.On("Read", matchers...)
	return &TaskReader_Read{Call: c}
}

// Read provides a mock function with given fields: ctx
func (_m *TaskReader) Read(ctx context.Context) (*core.TaskTemplate, error) {
	ret := _m.Called(ctx)

	var r0 *core.TaskTemplate
	if rf, ok := ret.Get(0).(func(context.Context) *core.TaskTemplate); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.TaskTemplate)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
