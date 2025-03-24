package food_test

import (
	"context"
	"testing"
	"time"

	"github.com/gatsu420/mary/app/usecases/food"
	"github.com/gatsu420/mary/common/errors"
	"github.com/gatsu420/mary/db/repository"
	mockrepository "github.com/gatsu420/mary/mocks/db/repository"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateFood(t *testing.T) {
	testCases := []struct {
		testName    string
		arg         *food.CreateFoodParams
		repoErr     error
		expectedErr error
	}{
		{
			testName: "repository error",
			arg: &food.CreateFoodParams{
				Name:           "test",
				TypeID:         99,
				IntakeStatusID: 99,
				FeederID:       99,
				LocationID:     99,
				Remarks: pgtype.Text{
					String: "test",
					Valid:  true,
				},
			},
			repoErr:     errors.New(errors.InternalServerError, "repository error"),
			expectedErr: errors.New(errors.InternalServerError, "DB failed to create food"),
		},
		{
			testName: "create food successfully",
			arg: &food.CreateFoodParams{
				Name:           "test",
				TypeID:         99,
				IntakeStatusID: 99,
				FeederID:       99,
				LocationID:     99,
				Remarks: pgtype.Text{
					String: "test",
					Valid:  true,
				},
			},
			repoErr:     nil,
			expectedErr: nil,
		},
	}
	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			mockRepo := mockrepository.NewMockQuerier(t)
			mockRepo.EXPECT().CreateFood(
				mock.Anything,
				mock.AnythingOfType("*repository.CreateFoodParams"),
			).Return(tc.repoErr).Once()

			usecases := food.NewUsecases(mockRepo)
			err := usecases.CreateFood(ctx, tc.arg)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}

func TestListFood(t *testing.T) {
	testCases := []struct {
		testName         string
		arg              *food.ListFoodParams
		expectedFoodList []repository.ListFoodRow
		repoErr          error
		expectedErr      error
	}{
		{
			testName: "repository error",
			arg: &food.ListFoodParams{
				StartTimestamp: pgtype.Timestamptz{Time: time.Now(), Valid: true},
				EndTimestamp:   pgtype.Timestamptz{Time: time.Now(), Valid: true},
				Type:           pgtype.Text{String: "test", Valid: true},
				IntakeStatus:   pgtype.Text{String: "test", Valid: true},
				Feeder:         pgtype.Text{String: "test", Valid: true},
				Location:       pgtype.Text{String: "test", Valid: true},
			},
			expectedFoodList: nil,
			repoErr:          errors.New(errors.InternalServerError, "repository error"),
			expectedErr:      errors.New(errors.InternalServerError, "DB failed to list food"),
		},
		{
			testName: "list food successfully",
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
	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			mockRepo := mockrepository.NewMockQuerier(t)
			mockRepo.EXPECT().ListFood(
				mock.Anything,
				mock.AnythingOfType("*repository.ListFoodParams"),
			).Return(tc.expectedFoodList, tc.repoErr).Once()

			usecases := food.NewUsecases(mockRepo)
			foodList, err := usecases.ListFood(ctx, tc.arg)
			assert.Equal(t, tc.expectedFoodList, foodList)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}

func TestGetFood(t *testing.T) {
	testCases := []struct {
		testName     string
		id           int32
		expectedFood repository.GetFoodRow
		repoErr      error
		expectedErr  error
	}{
		{
			testName:     "repository error",
			id:           99,
			expectedFood: repository.GetFoodRow{},
			repoErr:      errors.New(errors.InternalServerError, "repository error"),
			expectedErr:  errors.New(errors.InternalServerError, "DB failed to get food"),
		},
		{
			testName:     "error no rows when geting food",
			id:           99,
			expectedFood: repository.GetFoodRow{},
			repoErr:      pgx.ErrNoRows,
			expectedErr:  errors.New(errors.NotFoundError, "food not found"),
		},
		{
			testName: "get food successfully",
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
		},
	}
	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			mockRepo := mockrepository.NewMockQuerier(t)
			mockRepo.EXPECT().GetFood(
				mock.Anything,
				mock.AnythingOfType("int32"),
			).Return(tc.expectedFood, tc.repoErr).Once()

			usecases := food.NewUsecases(mockRepo)
			food, err := usecases.GetFood(ctx, tc.id)
			assert.Equal(t, tc.expectedFood, food)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}
