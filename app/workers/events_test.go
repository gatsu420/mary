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
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			if tc.withSetCalledMethods {
				tempvalue.SetCalledMethods("GetFood")

				s.mockUsecase.EXPECT().CreateEvent(
					mock.Anything,
					mock.AnythingOfType("*events.CreateEventParams"),
				).Return(nil).Once()
			}

			ticker := make(chan time.Time)
			go s.worker.Create(ctx, ticker)
			ticker <- time.Now()

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
