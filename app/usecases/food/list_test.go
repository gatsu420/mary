package food_test

import (
	"github.com/gatsu420/mary/app/repository"
	"github.com/gatsu420/mary/app/usecases/food"
	"github.com/gatsu420/mary/common/errors"
	"github.com/stretchr/testify/mock"
)

func (s *testSuite) Test_ListFood() {
	testCases := []struct {
		testName         string
		arg              *food.ListFoodParams
		repoFoodList     []repository.ListFoodRow
		repoErr          error
		eventErr         error
		expectedFoodList []food.ListFoodRow
		expectedErr      error
	}{
		{
			testName: "repo error",
			arg: &food.ListFoodParams{
				StartTimestamp: s.pgTimestamptz,
				EndTimestamp:   s.pgTimestamptz,
				Type:           s.pgText,
				IntakeStatus:   s.pgText,
				Feeder:         s.pgText,
				Location:       s.pgText,
			},
			repoFoodList:     nil,
			repoErr:          errors.New(errors.InternalServerError, "DB error"),
			expectedFoodList: nil,
			expectedErr:      errors.New(errors.InternalServerError, "DB failed to list food"),
		},
		{
			testName: "event error",
			arg: &food.ListFoodParams{
				StartTimestamp: s.pgTimestamptz,
				EndTimestamp:   s.pgTimestamptz,
				Type:           s.pgText,
				IntakeStatus:   s.pgText,
				Feeder:         s.pgText,
				Location:       s.pgText,
			},
			repoFoodList: []repository.ListFoodRow{
				{
					ID:           99,
					Name:         "test",
					Type:         s.pgText,
					IntakeStatus: s.pgText,
					Feeder:       s.pgText,
					Location:     s.pgText,
					Remarks:      s.pgText,
					CreatedAt:    s.pgTimestamptz,
					UpdatedAt:    s.pgTimestamptz,
				},
			},
			repoErr:          nil,
			eventErr:         errors.New(errors.InternalServerError, "cache error"),
			expectedFoodList: nil,
			expectedErr:      errors.New(errors.InternalServerError, "cache failed to create event"),
		},
		{
			testName: "success",
			arg: &food.ListFoodParams{
				StartTimestamp: s.pgTimestamptz,
				EndTimestamp:   s.pgTimestamptz,
				Type:           s.pgText,
				IntakeStatus:   s.pgText,
				Feeder:         s.pgText,
				Location:       s.pgText,
			},
			repoFoodList: []repository.ListFoodRow{
				{
					ID:           99,
					Name:         "test",
					Type:         s.pgText,
					IntakeStatus: s.pgText,
					Feeder:       s.pgText,
					Location:     s.pgText,
					Remarks:      s.pgText,
					CreatedAt:    s.pgTimestamptz,
					UpdatedAt:    s.pgTimestamptz,
				},
			},
			repoErr:  nil,
			eventErr: nil,
			expectedFoodList: []food.ListFoodRow{
				{
					ID:           99,
					Name:         "test",
					Type:         s.pgText,
					IntakeStatus: s.pgText,
					Feeder:       s.pgText,
					Location:     s.pgText,
					Remarks:      s.pgText,
					CreatedAt:    s.pgTimestamptz,
					UpdatedAt:    s.pgTimestamptz,
				},
			},
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.testName, func() {
			s.mockRepo.EXPECT().ListFood(
				mock.Anything,
				mock.AnythingOfType("repository.ListFoodParams"),
			).Return(tc.repoFoodList, tc.repoErr).Once()

			if tc.repoErr == nil {
				s.mockCache.EXPECT().CreateEvent(
					mock.Anything,
					mock.AnythingOfType("cache.CreateEventParams"),
				).Return(tc.eventErr).Once()
			}

			food, err := s.usecase.ListFood(s.ctx, tc.arg)
			s.Equal(tc.expectedFoodList, food)
			s.Equal(tc.expectedErr, err)
		})
	}
}
