package food_test

import (
	"time"

	"github.com/gatsu420/mary/app/usecases/food"
	"github.com/gatsu420/mary/common/errors"
	"github.com/gatsu420/mary/db/repository"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/mock"
)

func (s *testSuite) Test_CreateFood() {
	testCases := []struct {
		testName    string
		arg         *food.CreateFoodParams
		repoErr     error
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
			repoErr:     errors.New(errors.InternalServerError, "repo error"),
			expectedErr: errors.New(errors.InternalServerError, "DB failed to create food"),
		},
		{
			testName: "created food successfully",
			arg: &food.CreateFoodParams{
				Name:           "test",
				TypeID:         99,
				IntakeStatusID: 99,
				FeederID:       99,
				LocationID:     99,
				Remarks:        pgtype.Text{String: "test", Valid: true},
			},
			repoErr:     nil,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.testName, func() {
			mockRepo := s.mockrepo
			mockRepo.EXPECT().CreateFood(
				mock.Anything,
				mock.AnythingOfType("*repository.CreateFoodParams"),
			).Return(tc.repoErr).Once()

			usecases := s.usecases
			err := usecases.CreateFood(s.ctx, tc.arg)
			s.Equal(tc.expectedErr, err)
		})
	}
}

func (s *testSuite) Test_ListFood() {
	testCases := []struct {
		testName         string
		arg              *food.ListFoodParams
		expectedFoodList []repository.ListFoodRow
		repoErr          error
		expectedErr      error
	}{
		{
			testName: "repo error",
			arg: &food.ListFoodParams{
				StartTimestamp: pgtype.Timestamptz{Time: time.Now(), Valid: true},
				EndTimestamp:   pgtype.Timestamptz{Time: time.Now(), Valid: true},
				Type:           pgtype.Text{String: "test", Valid: true},
				IntakeStatus:   pgtype.Text{String: "test", Valid: true},
				Feeder:         pgtype.Text{String: "test", Valid: true},
				Location:       pgtype.Text{String: "test", Valid: true},
			},
			expectedFoodList: nil,
			repoErr:          errors.New(errors.InternalServerError, "repo error"),
			expectedErr:      errors.New(errors.InternalServerError, "DB failed to list food"),
		},
		{
			testName: "listed food successfully",
			arg: &food.ListFoodParams{
				StartTimestamp: pgtype.Timestamptz{Time: time.Now(), Valid: true},
				EndTimestamp:   pgtype.Timestamptz{Time: time.Now(), Valid: true},
				Type:           pgtype.Text{String: "test", Valid: true},
				IntakeStatus:   pgtype.Text{String: "test", Valid: true},
				Feeder:         pgtype.Text{String: "test", Valid: true},
				Location:       pgtype.Text{String: "test", Valid: true},
			},
			expectedFoodList: []repository.ListFoodRow{
				{
					ID:           99,
					Name:         "test food",
					Type:         pgtype.Text{String: "test", Valid: true},
					IntakeStatus: pgtype.Text{String: "test", Valid: true},
					Feeder:       pgtype.Text{String: "test", Valid: true},
					Location:     pgtype.Text{String: "test", Valid: true},
					Remarks:      pgtype.Text{String: "test", Valid: true},
					CreatedAt:    pgtype.Timestamptz{Time: time.Now(), Valid: true},
					UpdatedAt:    pgtype.Timestamptz{Time: time.Now(), Valid: true},
				},
			},
			repoErr:     nil,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.testName, func() {
			mockRepo := s.mockrepo
			mockRepo.EXPECT().ListFood(
				mock.Anything,
				mock.AnythingOfType("*repository.ListFoodParams"),
			).Return(tc.expectedFoodList, tc.repoErr).Once()

			usecases := s.usecases
			foodList, err := usecases.ListFood(s.ctx, tc.arg)
			s.Equal(tc.expectedFoodList, foodList)
			s.Equal(tc.expectedErr, err)
		})
	}
}

func (s *testSuite) Test_GetFood() {
	testCases := []struct {
		testName     string
		id           int32
		expectedFood repository.GetFoodRow
		repoErr      error
		expectedErr  error
	}{
		{
			testName:     "repo error",
			id:           99,
			expectedFood: repository.GetFoodRow{},
			repoErr:      errors.New(errors.InternalServerError, "DB failed to get food"),
			expectedErr:  errors.New(errors.InternalServerError, "DB failed to get food"),
		},
		{
			testName:     "no rows error",
			id:           99,
			expectedFood: repository.GetFoodRow{},
			repoErr:      pgx.ErrNoRows,
			expectedErr:  errors.New(errors.NotFoundError, "food not found"),
		},
		{
			testName: "got food successfully",
			id:       99,
			expectedFood: repository.GetFoodRow{
				ID:           99,
				Name:         "test",
				Type:         pgtype.Text{String: "test", Valid: true},
				IntakeStatus: pgtype.Text{String: "test", Valid: true},
				Feeder:       pgtype.Text{String: "test", Valid: true},
				Location:     pgtype.Text{String: "test", Valid: true},
				Remarks:      pgtype.Text{String: "test", Valid: true},
				CreatedAt:    pgtype.Timestamptz{Time: time.Now(), Valid: true},
				UpdatedAt:    pgtype.Timestamptz{Time: time.Now(), Valid: true},
			},
			repoErr:     nil,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.testName, func() {
			mockRepo := s.mockrepo
			mockRepo.EXPECT().GetFood(
				mock.Anything,
				mock.AnythingOfType("int32"),
			).Return(tc.expectedFood, tc.repoErr).Once()

			usecases := s.usecases
			food, err := usecases.GetFood(s.ctx, tc.id)
			s.Equal(tc.expectedFood, food)
			s.Equal(tc.expectedErr, err)
		})
	}
}
