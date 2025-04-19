package food_test

import (
	"github.com/gatsu420/mary/app/repository"
	"github.com/gatsu420/mary/app/usecases/food"
	"github.com/gatsu420/mary/common/errors"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/mock"
)

func (s *testSuite) Test_GetFood() {
	testCases := []struct {
		testName     string
		id           int32
		repoFood     repository.GetFoodRow
		repoErr      error
		eventErr     error
		expectedFood *food.GetFoodRow
		expectedErr  error
	}{
		{
			testName:     "repo error",
			id:           99,
			repoFood:     repository.GetFoodRow{},
			repoErr:      errors.New(errors.InternalServerError, "DB error"),
			expectedFood: nil,
			expectedErr:  errors.New(errors.InternalServerError, "DB failed to get food"),
		},
		{
			testName:     "ID is not found",
			id:           99,
			repoFood:     repository.GetFoodRow{},
			repoErr:      pgx.ErrNoRows,
			expectedFood: nil,
			expectedErr:  errors.New(errors.NotFoundError, "food not found"),
		},
		{
			testName: "event error",
			id:       99,
			repoFood: repository.GetFoodRow{
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
			repoErr:      nil,
			eventErr:     errors.New(errors.InternalServerError, "cache error"),
			expectedFood: nil,
			expectedErr:  errors.New(errors.InternalServerError, "cache failed to create event"),
		},
		{
			testName: "success",
			id:       99,
			repoFood: repository.GetFoodRow{
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
			repoErr:  nil,
			eventErr: nil,
			expectedFood: &food.GetFoodRow{
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
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.testName, func() {
			s.mockRepo.EXPECT().GetFood(
				mock.Anything,
				mock.AnythingOfType("int32"),
			).Return(tc.repoFood, tc.repoErr).Once()

			if tc.repoErr == nil {
				s.mockCache.EXPECT().CreateEvent(
					mock.Anything,
					mock.AnythingOfType("cache.CreateEventParams"),
				).Return(tc.eventErr).Once()
			}

			food, err := s.usecase.GetFood(s.ctx, tc.id)
			s.Equal(tc.expectedFood, food)
			s.Equal(tc.expectedErr, err)
		})
	}
}
