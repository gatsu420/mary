// Code generated by mockery v2.53.3. DO NOT EDIT.

package mockusers

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockUsecase is an autogenerated mock type for the Usecase type
type MockUsecase struct {
	mock.Mock
}

type MockUsecase_Expecter struct {
	mock *mock.Mock
}

func (_m *MockUsecase) EXPECT() *MockUsecase_Expecter {
	return &MockUsecase_Expecter{mock: &_m.Mock}
}

// CheckUserIsExisting provides a mock function with given fields: ctx, username
func (_m *MockUsecase) CheckUserIsExisting(ctx context.Context, username string) error {
	ret := _m.Called(ctx, username)

	if len(ret) == 0 {
		panic("no return value specified for CheckUserIsExisting")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, username)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockUsecase_CheckUserIsExisting_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CheckUserIsExisting'
type MockUsecase_CheckUserIsExisting_Call struct {
	*mock.Call
}

// CheckUserIsExisting is a helper method to define mock.On call
//   - ctx context.Context
//   - username string
func (_e *MockUsecase_Expecter) CheckUserIsExisting(ctx interface{}, username interface{}) *MockUsecase_CheckUserIsExisting_Call {
	return &MockUsecase_CheckUserIsExisting_Call{Call: _e.mock.On("CheckUserIsExisting", ctx, username)}
}

func (_c *MockUsecase_CheckUserIsExisting_Call) Run(run func(ctx context.Context, username string)) *MockUsecase_CheckUserIsExisting_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockUsecase_CheckUserIsExisting_Call) Return(_a0 error) *MockUsecase_CheckUserIsExisting_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUsecase_CheckUserIsExisting_Call) RunAndReturn(run func(context.Context, string) error) *MockUsecase_CheckUserIsExisting_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockUsecase creates a new instance of MockUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockUsecase {
	mock := &MockUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
