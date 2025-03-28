package food_test

import (
	"github.com/gatsu420/mary/app/usecases/food"
	"github.com/gatsu420/mary/common/errors"
	"github.com/gatsu420/mary/db/repository"
	"github.com/stretchr/testify/mock"
)

func (s *testSuite) Test_ListFood() {
	testCases := []struct {
		testName         string
		arg              *food.ListFoodParams
		repoFoodList     []repository.ListFoodRow
		repoErr          error
		expectedFoodList []repository.ListFoodRow
		expectedErr      error
	}{
		{
			testName: "repo error",
			arg: &food.ListFoodParams{
				StartTimestamp: s.mockPGTimestamptz,
				EndTimestamp:   s.mockPGTimestamptz,
				Type:           s.mockPGText,
				IntakeStatus:   s.mockPGText,
				Feeder:         s.mockPGText,
				Location:       s.mockPGText,
			},
			repoFoodList:     nil,
			repoErr:          errors.New(errors.InternalServerError, "DB error"),
			expectedFoodList: nil,
			expectedErr:      errors.New(errors.InternalServerError, "DB failed to list food"),
		},
		{
			testName: "success",
			arg: &food.ListFoodParams{
				StartTimestamp: s.mockPGTimestamptz,
				EndTimestamp:   s.mockPGTimestamptz,
				Type:           s.mockPGText,
				IntakeStatus:   s.mockPGText,
				Feeder:         s.mockPGText,
				Location:       s.mockPGText,
			},
			repoFoodList: []repository.ListFoodRow{
				{
					ID:           99,
					Name:         "test",
					Type:         s.mockPGText,
					IntakeStatus: s.mockPGText,
					Feeder:       s.mockPGText,
					Location:     s.mockPGText,
					Remarks:      s.mockPGText,
					CreatedAt:    s.mockPGTimestamptz,
					UpdatedAt:    s.mockPGTimestamptz,
				},
			},
			repoErr: nil,
			expectedFoodList: []repository.ListFoodRow{
				{
					ID:           99,
					Name:         "test",
					Type:         s.mockPGText,
					IntakeStatus: s.mockPGText,
					Feeder:       s.mockPGText,
					Location:     s.mockPGText,
					Remarks:      s.mockPGText,
					CreatedAt:    s.mockPGTimestamptz,
					UpdatedAt:    s.mockPGTimestamptz,
				},
			},
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.testName, func() {
			s.mockRepo.EXPECT().ListFood(
				mock.Anything,
				mock.AnythingOfType("*repository.ListFoodParams"),
			).Return(tc.repoFoodList, tc.repoErr).Once()

			food, err := s.usecases.ListFood(s.ctx, tc.arg)
			s.Equal(tc.expectedFoodList, food)
			s.Equal(tc.expectedErr, err)
		})
	}
}
