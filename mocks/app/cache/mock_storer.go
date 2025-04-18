// Code generated by mockery v2.53.3. DO NOT EDIT.

package mockcache

import (
	context "context"

	cache "github.com/gatsu420/mary/app/cache"

	mock "github.com/stretchr/testify/mock"
)

// MockStorer is an autogenerated mock type for the Storer type
type MockStorer struct {
	mock.Mock
}

type MockStorer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockStorer) EXPECT() *MockStorer_Expecter {
	return &MockStorer_Expecter{mock: &_m.Mock}
}

// CreateEvent provides a mock function with given fields: ctx, arg
func (_m *MockStorer) CreateEvent(ctx context.Context, arg cache.CreateEventParams) error {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for CreateEvent")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, cache.CreateEventParams) error); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockStorer_CreateEvent_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateEvent'
type MockStorer_CreateEvent_Call struct {
	*mock.Call
}

// CreateEvent is a helper method to define mock.On call
//   - ctx context.Context
//   - arg cache.CreateEventParams
func (_e *MockStorer_Expecter) CreateEvent(ctx interface{}, arg interface{}) *MockStorer_CreateEvent_Call {
	return &MockStorer_CreateEvent_Call{Call: _e.mock.On("CreateEvent", ctx, arg)}
}

func (_c *MockStorer_CreateEvent_Call) Run(run func(ctx context.Context, arg cache.CreateEventParams)) *MockStorer_CreateEvent_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(cache.CreateEventParams))
	})
	return _c
}

func (_c *MockStorer_CreateEvent_Call) Return(_a0 error) *MockStorer_CreateEvent_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStorer_CreateEvent_Call) RunAndReturn(run func(context.Context, cache.CreateEventParams) error) *MockStorer_CreateEvent_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockStorer creates a new instance of MockStorer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockStorer(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockStorer {
	mock := &MockStorer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
