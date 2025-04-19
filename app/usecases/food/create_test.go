package food_test

import (
	"github.com/gatsu420/mary/app/usecases/food"
	"github.com/gatsu420/mary/common/errors"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/mock"
)

func (s *testSuite) Test_CreateFood() {
	testCases := []struct {
		testName    string
		arg         *food.CreateFoodParams
		repoErr     error
		eventErr    error
		expectedErr error
	}{
		{
			testName: "repo error",
			arg: &food.CreateFoodParams{
				Name:           "test",
				TypeID:         99,
				IntakeStatusID: 99,
				FeederID:       99,
				LocationID:     99,
				Remarks:        pgtype.Text{String: "test", Valid: true},
			},
			repoErr:     errors.New(errors.InternalServerError, "DB error"),
			expectedErr: errors.New(errors.InternalServerError, "DB failed to create food"),
		},
		{
			testName: "event error",
			arg: &food.CreateFoodParams{
				Name:           "test",
				TypeID:         99,
				IntakeStatusID: 99,
				FeederID:       99,
				LocationID:     99,
				Remarks:        pgtype.Text{String: "test", Valid: true},
			},
			repoErr:     nil,
			eventErr:    errors.New(errors.InternalServerError, "cache error"),
			expectedErr: errors.New(errors.InternalServerError, "cache failed to create event"),
		},
		{
			testName: "success",
			arg: &food.CreateFoodParams{
				Name:           "test",
				TypeID:         99,
				IntakeStatusID: 99,
				FeederID:       99,
				LocationID:     99,
				Remarks:        pgtype.Text{String: "test", Valid: true},
			},
			repoErr:     nil,
			eventErr:    nil,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.testName, func() {
			s.mockRepo.EXPECT().CreateFood(
				mock.Anything,
				mock.AnythingOfType("repository.CreateFoodParams"),
			).Return(tc.repoErr).Once()

			if tc.repoErr == nil {
				s.mockCache.EXPECT().CreateEvent(
					mock.Anything,
					mock.AnythingOfType("cache.CreateEventParams"),
				).Return(tc.eventErr).Once()
			}

			err := s.usecase.CreateFood(s.ctx, tc.arg)
			s.Equal(tc.expectedErr, err)
		})
	}
}
