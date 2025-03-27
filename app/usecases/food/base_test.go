package food_test

import (
	"context"
	"testing"

	"github.com/gatsu420/mary/app/usecases/food"
	"github.com/gatsu420/mary/common/errors"
	mockfood "github.com/gatsu420/mary/mocks/app/usecases"
	mockrepository "github.com/gatsu420/mary/mocks/db/repository"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type testSuite struct {
	suite.Suite
	ctx          context.Context
	mockrepo     *mockrepository.MockQuerier
	mockusecases *mockfood.MockUsecases
	usecases     food.Usecases
}

func (s *testSuite) SetupSuite() {
	s.ctx = context.Background()
}

func (s *testSuite) SetupTest() {
	s.mockrepo = mockrepository.NewMockQuerier(s.T())
	s.mockusecases = mockfood.NewMockUsecases(s.T())
	s.usecases = food.NewUsecases(s.mockrepo)
}

func Test(t *testing.T) {
	suite.Run(t, &testSuite{})
}

func (s *testSuite) Test_CheckFoodIsRemoved() {
	testCases := []struct {
		testName    string
		id          int32
		isRemoved   bool
		repoErr     error
		expectedErr error
	}{
		{
			testName:    "repo error",
			id:          99,
			isRemoved:   true,
			repoErr:     errors.New(errors.InternalServerError, "repo error"),
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

	for _, tc := range testCases {
		s.Run(tc.testName, func() {
			mockRepo := s.mockrepo
			mockRepo.EXPECT().CheckFoodIsRemoved(
				mock.Anything,
				mock.AnythingOfType("int32"),
			).Return(tc.isRemoved, tc.repoErr).Once()

			usecases := s.usecases
			err := usecases.CheckFoodIsRemoved(s.ctx, tc.id)
			s.Equal(tc.expectedErr, err)
		})
	}
}
