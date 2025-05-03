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
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			if tc.withSetCalledMethods {
				tempvalue.SetCalledMethods("GetFood")

				s.mockUsecase.EXPECT().CreateEvent(
					mock.Anything,
					mock.AnythingOfType("*events.CreateEventParams"),
				).Return(nil).Once()
			}

			// Since Create deliberately contains lag while looping, we need to finish the lag
			// first if we run it as ordinary function and the loop will be done immediately
			// due to timeout. However, we can spawn it into another goroutine and both asserts
			// can kick in due to they being not concurrent-safe.
			//
			// Another 1 second lag is added because spawning goroutine has some overhead, but
			// very short in time.
			go s.worker.Create(ctx)
			time.Sleep(1 * time.Second)

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
