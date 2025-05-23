package food_test

import (
	"github.com/gatsu420/mary/app/usecases/food"
	"github.com/gatsu420/mary/common/errors"
	"github.com/stretchr/testify/mock"
)

func (s *testSuite) Test_UpdateFood() {
	testCases := []struct {
		testName    string
		arg         *food.UpdateFoodParams
		repoRows    int64
		repoErr     error
		eventErr    error
		expectedErr error
	}{
		{
			testName: "repo error",
			arg: &food.UpdateFoodParams{
				Name:           s.pgText,
				TypeID:         s.pgInt4,
				IntakeStatusID: s.pgInt4,
				FeederID:       s.pgInt4,
				LocationID:     s.pgInt4,
				Remarks:        s.pgText,
				ID:             99,
			},
			repoRows:    0,
			repoErr:     errors.New(errors.InternalServerError, "DB error"),
			expectedErr: errors.New(errors.InternalServerError, "DB failed to update food"),
		},
		{
			testName: "ID is not found",
			arg: &food.UpdateFoodParams{
				Name:           s.pgText,
				TypeID:         s.pgInt4,
				IntakeStatusID: s.pgInt4,
				FeederID:       s.pgInt4,
				LocationID:     s.pgInt4,
				Remarks:        s.pgText,
				ID:             99,
			},
			repoRows:    0,
			repoErr:     nil,
			expectedErr: errors.New(errors.NotFoundError, "food not found"),
		},
		{
			testName: "event error",
			arg: &food.UpdateFoodParams{
				Name:           s.pgText,
				TypeID:         s.pgInt4,
				IntakeStatusID: s.pgInt4,
				FeederID:       s.pgInt4,
				LocationID:     s.pgInt4,
				Remarks:        s.pgText,
				ID:             99,
			},
			repoRows:    1,
			repoErr:     nil,
			eventErr:    errors.New(errors.InternalServerError, "cache error"),
			expectedErr: errors.New(errors.InternalServerError, "cache failed to create event"),
		},
		{
			testName: "success",
			arg: &food.UpdateFoodParams{
				Name:           s.pgText,
				TypeID:         s.pgInt4,
				IntakeStatusID: s.pgInt4,
				FeederID:       s.pgInt4,
				LocationID:     s.pgInt4,
				Remarks:        s.pgText,
				ID:             99,
			},
			repoRows:    1,
			repoErr:     nil,
			eventErr:    nil,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.testName, func() {
			s.mockRepo.EXPECT().UpdateFood(
				mock.Anything,
				mock.AnythingOfType("repository.UpdateFoodParams"),
			).Return(tc.repoRows, tc.repoErr).Once()

			if tc.repoErr == nil && tc.repoRows != 0 {
				s.mockCache.EXPECT().CreateEvent(
					mock.Anything,
					mock.AnythingOfType("cache.CreateEventParams"),
				).Return(tc.eventErr).Once()
			}

			err := s.usecase.UpdateFood(s.ctx, tc.arg)
			s.Equal(tc.expectedErr, err)
		})
	}
}
