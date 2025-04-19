package food_test

import (
	"github.com/gatsu420/mary/common/errors"
	"github.com/stretchr/testify/mock"
)

func (s *testSuite) Test_DeleteFood() {
	testCases := []struct {
		testName    string
		id          int32
		repoRows    int64
		repoErr     error
		eventErr    error
		expectedErr error
	}{
		{
			testName:    "repo error",
			id:          99,
			repoRows:    0,
			repoErr:     errors.New(errors.InternalServerError, "DB error"),
			expectedErr: errors.New(errors.InternalServerError, "DB failed to delete food"),
		},
		{
			testName:    "ID is not found",
			id:          99,
			repoRows:    0,
			repoErr:     nil,
			expectedErr: errors.New(errors.NotFoundError, "food not found"),
		},
		{
			testName:    "event error",
			id:          99,
			repoRows:    1,
			repoErr:     nil,
			eventErr:    errors.New(errors.InternalServerError, "cache error"),
			expectedErr: errors.New(errors.InternalServerError, "cache failed to create event"),
		},
		{
			testName:    "success",
			id:          99,
			repoRows:    1,
			repoErr:     nil,
			eventErr:    nil,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.testName, func() {
			s.mockRepo.EXPECT().DeleteFood(
				mock.Anything,
				mock.AnythingOfType("int32"),
			).Return(tc.repoRows, tc.repoErr).Once()

			if tc.repoErr == nil && tc.repoRows != 0 {
				s.mockCache.EXPECT().CreateEvent(
					mock.Anything,
					mock.AnythingOfType("cache.CreateEventParams"),
				).Return(tc.eventErr).Once()
			}

			err := s.usecase.DeleteFood(s.ctx, tc.id)
			s.Equal(tc.expectedErr, err)
		})
	}
}
