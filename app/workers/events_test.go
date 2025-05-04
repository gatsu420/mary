package workers_test

import (
	"context"
	"time"

	"github.com/gatsu420/mary/common/tempvalue"
	"github.com/stretchr/testify/mock"
)

// Create doesn't return anything, so we need to assert its behavior rather than return value.
// The test works by mocking usecase, calling Create, and checking if that usecase has been
// called or not.
func (s *testSuite) Test_Create() {
	testCases := []struct {
		testName             string
		withSetCalledMethods bool
	}{
		{
			testName:             "calledMethods are not set",
			withSetCalledMethods: false,
		},
		{
			testName:             "calledMethods are set",
			withSetCalledMethods: true,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.testName, func() {
			ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
			defer cancel()

			if tc.withSetCalledMethods {
				tempvalue.SetCalledMethods("GetFood")

				s.mockUsecase.EXPECT().CreateEvent(
					mock.Anything,
					mock.AnythingOfType("*events.CreateEventParams"),
				).Return(nil).Once()
			}

			// If we run Create as ordinary function, unit test need to wait for loop to exit
			// according to timeout. However, we can spawn it into another goroutine and both
			// asserts can kick in in the first loop due to they being not concurrent-safe.
			//
			// 11 second lag is added because:
			// 	1. 10 second for first loop (and will be caught by asserts)
			// 	2. 1 second for goroutine overhead
			go s.worker.Create(ctx)
			time.Sleep(11 * time.Second)

			if tc.withSetCalledMethods {
				s.mockUsecase.AssertCalled(s.T(), "CreateEvent",
					mock.Anything,
					mock.AnythingOfType("*events.CreateEventParams"),
				)
				return
			}
			s.mockUsecase.AssertNotCalled(s.T(), "CreateEvent",
				mock.Anything,
				mock.AnythingOfType("*events.CreateEventParams"),
			)
		})
	}
}
