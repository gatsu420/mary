package events_test

import (
	"github.com/gatsu420/mary/app/cache"
	"github.com/gatsu420/mary/app/usecases/events"
	"github.com/gatsu420/mary/common/errors"
	"github.com/stretchr/testify/mock"
)

func (s *testSuite) Test_CreateEvent() {
	testCases := []struct {
		testName             string
		arg                  *events.CreateEventParams
		cacheEvents          []cache.GetEventResponse
		cacheGetEventsErr    error
		repoEventRownum      int64
		repoCreateEventErr   error
		cacheDeleteEventsErr error
		expectedErr          error
	}{
		{
			testName: "cache error when getting events",
			arg: &events.CreateEventParams{
				Name: map[string]struct{}{
					"GetFood": {},
				},
			},
			cacheEvents:       nil,
			cacheGetEventsErr: errors.New(errors.InternalServerError, "cache error"),
			expectedErr:       errors.New(errors.InternalServerError, "cache failed to get events"),
		},
		{
			testName: "repo failed when creating event",
			arg: &events.CreateEventParams{
				Name: map[string]struct{}{
					"GetFood": {},
				},
			},
			cacheEvents: []cache.GetEventResponse{
				{
					Name:   "GetFood",
					UserID: "test_user",
					CreatedAtEpoch: []int64{
						1746075368,
						1746075377,
					},
				},
			},
			cacheGetEventsErr:  nil,
			repoCreateEventErr: errors.New(errors.InternalServerError, "repo error"),
			expectedErr:        errors.New(errors.InternalServerError, "DB failed to create event"),
		},
		{
			testName: "cache failed when deleting events",
			arg: &events.CreateEventParams{
				Name: map[string]struct{}{
					"GetFood": {},
				},
			},
			cacheEvents: []cache.GetEventResponse{
				{
					Name:   "GetFood",
					UserID: "test_user",
					CreatedAtEpoch: []int64{
						1746075368,
						1746075377,
					},
				},
			},
			cacheGetEventsErr:    nil,
			repoEventRownum:      0,
			repoCreateEventErr:   nil,
			cacheDeleteEventsErr: errors.New(errors.InternalServerError, "cache error"),
			expectedErr:          errors.New(errors.InternalServerError, "cache failed to delete event"),
		},
		{
			testName: "success",
			arg: &events.CreateEventParams{
				Name: map[string]struct{}{
					"GetFood": {},
				},
			},
			cacheEvents: []cache.GetEventResponse{
				{
					Name:   "GetFood",
					UserID: "test_user",
					CreatedAtEpoch: []int64{
						1746075368,
						1746075377,
					},
				},
			},
			cacheGetEventsErr:    nil,
			repoEventRownum:      0,
			repoCreateEventErr:   nil,
			cacheDeleteEventsErr: nil,
			expectedErr:          nil,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.testName, func() {
			s.mockCache.EXPECT().GetEvents(
				mock.Anything,
				mock.AnythingOfType("cache.GetEventParams"),
			).Return(tc.cacheEvents, tc.cacheGetEventsErr).Once()

			if tc.cacheGetEventsErr == nil {
				s.mockRepo.EXPECT().CreateEvent(
					mock.Anything,
					mock.AnythingOfType("[]repository.CreateEventParams"),
				).Return(tc.repoEventRownum, tc.repoCreateEventErr).Once()
			}

			if tc.cacheGetEventsErr == nil && tc.repoCreateEventErr == nil {
				s.mockCache.EXPECT().DeleteEvents(
					mock.Anything,
					mock.AnythingOfType("cache.DeleteEventParams"),
				).Return(tc.cacheDeleteEventsErr).Once()
			}

			err := s.usecase.CreateEvent(s.ctx, tc.arg)
			s.Equal(tc.expectedErr, err)
		})
	}
}
