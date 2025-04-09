package handlers_test

import (
	"context"
	"testing"

	apifoodv1 "github.com/gatsu420/mary/api/gen/go/food/v1"
	"github.com/gatsu420/mary/app/handlers"
	"github.com/gatsu420/mary/common/errors"
	mockfood "github.com/gatsu420/mary/mocks/app/usecases/food"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type testSuite struct {
	suite.Suite
	mockUsecase *mockfood.MockUsecase
	server      handlers.FoodServer

	ctx context.Context
}

func (s *testSuite) SetupSuite() {
	s.ctx = context.Background()
}

func (s *testSuite) SetupTest() {
	s.mockUsecase = mockfood.NewMockUsecase(s.T())
	s.server = *handlers.NewFoodServer(s.mockUsecase)
}

func Test(t *testing.T) {
	suite.Run(t, &testSuite{})
}

func (s *testSuite) Test_Create() {
	testCases := []struct {
		testName     string
		req          *apifoodv1.CreateRequest
		usecaseErr   error
		expectedResp *apifoodv1.CreateResponse
		expectedErr  error
	}{
		{
			testName: "usecase error",
			req: &apifoodv1.CreateRequest{
				Name:           "some food",
				TypeId:         99,
				IntakeStatusId: 99,
				FeederId:       99,
				LocationId:     99,
				Remarks: &wrapperspb.StringValue{
					Value: "some remarks",
				},
			},
			usecaseErr:   errors.New(errors.BadRequestError, "bad request"),
			expectedResp: nil,
			expectedErr:  errors.New(errors.BadRequestError, "bad request"),
		},
		{
			testName: "success",
			req: &apifoodv1.CreateRequest{
				Name:           "some food",
				TypeId:         99,
				IntakeStatusId: 99,
				FeederId:       99,
				LocationId:     99,
				Remarks: &wrapperspb.StringValue{
					Value: "some remarks",
				},
			},
			usecaseErr:   nil,
			expectedResp: &apifoodv1.CreateResponse{},
			expectedErr:  nil,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.testName, func() {
			s.mockUsecase.EXPECT().CreateFood(
				mock.Anything,
				mock.AnythingOfType("*food.CreateFoodParams"),
			).Return(tc.usecaseErr).Once()

			resp, err := s.server.Create(s.ctx, tc.req)
			s.Equal(tc.expectedResp, resp)
			s.Equal(tc.expectedErr, err)
		})
	}
}
