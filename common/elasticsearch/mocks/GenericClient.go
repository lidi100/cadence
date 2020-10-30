// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	elasticsearch "github.com/uber/cadence/common/elasticsearch"

	persistence "github.com/uber/cadence/common/persistence"
)

// GenericClient is an autogenerated mock type for the GenericClient type
type GenericClient struct {
	mock.Mock
}

// CountByQuery provides a mock function with given fields: ctx, index, query
func (_m *GenericClient) CountByQuery(ctx context.Context, index string, query string) (int64, error) {
	ret := _m.Called(ctx, index, query)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, string, string) int64); ok {
		r0 = rf(ctx, index, query)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, index, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateIndex provides a mock function with given fields: ctx, index
func (_m *GenericClient) CreateIndex(ctx context.Context, index string) error {
	ret := _m.Called(ctx, index)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, index)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IsNotFoundError provides a mock function with given fields: err
func (_m *GenericClient) IsNotFoundError(err error) bool {
	ret := _m.Called(err)

	var r0 bool
	if rf, ok := ret.Get(0).(func(error) bool); ok {
		r0 = rf(err)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// PutMapping provides a mock function with given fields: ctx, index, root, key, valueType
func (_m *GenericClient) PutMapping(ctx context.Context, index string, root string, key string, valueType string) error {
	ret := _m.Called(ctx, index, root, key, valueType)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string) error); ok {
		r0 = rf(ctx, index, root, key, valueType)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RunBulkProcessor provides a mock function with given fields: ctx, p
func (_m *GenericClient) RunBulkProcessor(ctx context.Context, p *elasticsearch.BulkProcessorParameters) (elasticsearch.GenericBulkProcessor, error) {
	ret := _m.Called(ctx, p)

	var r0 elasticsearch.GenericBulkProcessor
	if rf, ok := ret.Get(0).(func(context.Context, *elasticsearch.BulkProcessorParameters) elasticsearch.GenericBulkProcessor); ok {
		r0 = rf(ctx, p)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(elasticsearch.GenericBulkProcessor)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *elasticsearch.BulkProcessorParameters) error); ok {
		r1 = rf(ctx, p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ScanByQuery provides a mock function with given fields: ctx, request
func (_m *GenericClient) ScanByQuery(ctx context.Context, request *elasticsearch.ScanByQueryRequest) (*persistence.InternalListWorkflowExecutionsResponse, error) {
	ret := _m.Called(ctx, request)

	var r0 *persistence.InternalListWorkflowExecutionsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *elasticsearch.ScanByQueryRequest) *persistence.InternalListWorkflowExecutionsResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*persistence.InternalListWorkflowExecutionsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *elasticsearch.ScanByQueryRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Search provides a mock function with given fields: ctx, request
func (_m *GenericClient) Search(ctx context.Context, request *elasticsearch.SearchRequest) (*persistence.InternalListWorkflowExecutionsResponse, error) {
	ret := _m.Called(ctx, request)

	var r0 *persistence.InternalListWorkflowExecutionsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *elasticsearch.SearchRequest) *persistence.InternalListWorkflowExecutionsResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*persistence.InternalListWorkflowExecutionsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *elasticsearch.SearchRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchByQuery provides a mock function with given fields: ctx, request
func (_m *GenericClient) SearchByQuery(ctx context.Context, request *elasticsearch.SearchByQueryRequest) (*persistence.InternalListWorkflowExecutionsResponse, error) {
	ret := _m.Called(ctx, request)

	var r0 *persistence.InternalListWorkflowExecutionsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *elasticsearch.SearchByQueryRequest) *persistence.InternalListWorkflowExecutionsResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*persistence.InternalListWorkflowExecutionsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *elasticsearch.SearchByQueryRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchForOneClosedExecution provides a mock function with given fields: ctx, index, request
func (_m *GenericClient) SearchForOneClosedExecution(ctx context.Context, index string, request *persistence.InternalGetClosedWorkflowExecutionRequest) (*persistence.InternalGetClosedWorkflowExecutionResponse, error) {
	ret := _m.Called(ctx, index, request)

	var r0 *persistence.InternalGetClosedWorkflowExecutionResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, *persistence.InternalGetClosedWorkflowExecutionRequest) *persistence.InternalGetClosedWorkflowExecutionResponse); ok {
		r0 = rf(ctx, index, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*persistence.InternalGetClosedWorkflowExecutionResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, *persistence.InternalGetClosedWorkflowExecutionRequest) error); ok {
		r1 = rf(ctx, index, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
