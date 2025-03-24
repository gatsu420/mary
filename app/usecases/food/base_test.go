package food_test

import (
	"context"
	"testing"

	"github.com/gatsu420/mary/app/usecases/food"
	"github.com/gatsu420/mary/common/errors"
	mockrepository "github.com/gatsu420/mary/mocks/db/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCheckFoodIsRemoved(t *testing.T) {
	testCases := []struct {
		testName    string
		id          int32
		isRemoved   bool
		repoErr     error
		expectedErr error
	}{
		{
			testName:    "repository error",
			id:          99,
			isRemoved:   true,
			repoErr:     errors.New(errors.InternalServerError, "repository error"),
			expectedErr: errors.New(errors.InternalServerError, "DB failed to check if food is removed"),
		},
		{
			testName:    "food is not removed",
			id:          99,
			isRemoved:   false,
			repoErr:     nil,
			expectedErr: nil,
		},
		{
			testName:    "food is removed",
			id:          99,
			isRemoved:   true,
			repoErr:     nil,
			expectedErr: errors.New(errors.NotFoundError, "food not found"),
		},
	}
	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			mockRepo := mockrepository.NewMockQuerier(t)
			mockRepo.EXPECT().CheckFoodIsRemoved(
				mock.Anything,
				mock.AnythingOfType("int32"),
			).Return(tc.isRemoved, tc.repoErr).Once()

			usecases := food.NewUsecases(mockRepo)
			err := usecases.CheckFoodIsRemoved(ctx, tc.id)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}
