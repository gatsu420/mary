package food_test

import (
	"github.com/gatsu420/mary/app/usecases/food"
	"github.com/gatsu420/mary/common/errors"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/mock"
)

func (s *testSuite) Test_CreateFood() {
	testCases := []struct {
		testName      string
		arg           *food.CreateFoodParams
		repoErr       error
		expectedError error
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
			repoErr:       errors.New(errors.InternalServerError, "DB error"),
			expectedError: errors.New(errors.InternalServerError, "DB failed to create food"),
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
			repoErr:       nil,
			expectedError: nil,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.testName, func() {
			s.mockRepo.EXPECT().CreateFood(
				mock.Anything,
				mock.AnythingOfType("*repository.CreateFoodParams"),
			).Return(tc.repoErr).Once()

			err := s.usecases.CreateFood(s.ctx, tc.arg)
			s.Equal(tc.expectedError, err)
		})
	}
}
