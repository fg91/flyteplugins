// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	bitarray "github.com/flyteorg/flytestdlib/bitarray"

	mock "github.com/stretchr/testify/mock"
)

// DownloadResponse is an autogenerated mock type for the DownloadResponse type
type DownloadResponse struct {
	mock.Mock
}

type DownloadResponse_GetCachedCount struct {
	*mock.Call
}

func (_m DownloadResponse_GetCachedCount) Return(_a0 int) *DownloadResponse_GetCachedCount {
	return &DownloadResponse_GetCachedCount{Call: _m.Call.Return(_a0)}
}

func (_m *DownloadResponse) OnGetCachedCount() *DownloadResponse_GetCachedCount {
	c := _m.On("GetCachedCount")
	return &DownloadResponse_GetCachedCount{Call: c}
}

func (_m *DownloadResponse) OnGetCachedCountMatch(matchers ...interface{}) *DownloadResponse_GetCachedCount {
	c := _m.On("GetCachedCount", matchers...)
	return &DownloadResponse_GetCachedCount{Call: c}
}

// GetCachedCount provides a mock function with given fields:
func (_m *DownloadResponse) GetCachedCount() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

type DownloadResponse_GetCachedResults struct {
	*mock.Call
}

func (_m DownloadResponse_GetCachedResults) Return(_a0 *bitarray.BitSet) *DownloadResponse_GetCachedResults {
	return &DownloadResponse_GetCachedResults{Call: _m.Call.Return(_a0)}
}

func (_m *DownloadResponse) OnGetCachedResults() *DownloadResponse_GetCachedResults {
	c := _m.On("GetCachedResults")
	return &DownloadResponse_GetCachedResults{Call: c}
}

func (_m *DownloadResponse) OnGetCachedResultsMatch(matchers ...interface{}) *DownloadResponse_GetCachedResults {
	c := _m.On("GetCachedResults", matchers...)
	return &DownloadResponse_GetCachedResults{Call: c}
}

// GetCachedResults provides a mock function with given fields:
func (_m *DownloadResponse) GetCachedResults() *bitarray.BitSet {
	ret := _m.Called()

	var r0 *bitarray.BitSet
	if rf, ok := ret.Get(0).(func() *bitarray.BitSet); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bitarray.BitSet)
		}
	}

	return r0
}

type DownloadResponse_GetResultsSize struct {
	*mock.Call
}

func (_m DownloadResponse_GetResultsSize) Return(_a0 int) *DownloadResponse_GetResultsSize {
	return &DownloadResponse_GetResultsSize{Call: _m.Call.Return(_a0)}
}

func (_m *DownloadResponse) OnGetResultsSize() *DownloadResponse_GetResultsSize {
	c := _m.On("GetResultsSize")
	return &DownloadResponse_GetResultsSize{Call: c}
}

func (_m *DownloadResponse) OnGetResultsSizeMatch(matchers ...interface{}) *DownloadResponse_GetResultsSize {
	c := _m.On("GetResultsSize", matchers...)
	return &DownloadResponse_GetResultsSize{Call: c}
}

// GetResultsSize provides a mock function with given fields:
func (_m *DownloadResponse) GetResultsSize() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}
