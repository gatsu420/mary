package handlers_test

import (
	"context"
	"testing"
	"time"

	apifoodv1 "github.com/gatsu420/mary/api/gen/go/food/v1"
	"github.com/gatsu420/mary/app/handlers"
	"github.com/gatsu420/mary/common/errors"
	"github.com/gatsu420/mary/db/repository"
	mockfood "github.com/gatsu420/mary/mocks/app/usecases/food"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type testSuite struct {
	suite.Suite
	mockUsecase *mockfood.MockUsecase
	server      handlers.FoodServer

	stubCtx      context.Context
	stubTimeSec  int
	stubTimeNano int

	stubStringWrapper *wrapperspb.StringValue
	stubString        string
	stubPGText        pgtype.Text

	stubTimestampWrapper *timestamppb.Timestamp
	stubPGTimestamptz    pgtype.Timestamptz
}

func (s *testSuite) SetupSuite() {
	s.stubCtx = context.Background()
	s.stubTimeSec = 1744208782
	s.stubTimeNano = 99

	s.stubStringWrapper = &wrapperspb.StringValue{Value: "test"}
	s.stubString = "test"
	s.stubPGText = pgtype.Text{String: "test", Valid: true}

	s.stubTimestampWrapper = &timestamppb.Timestamp{
		Seconds: int64(s.stubTimeSec),
		Nanos:   int32(s.stubTimeNano),
	}
	s.stubPGTimestamptz = pgtype.Timestamptz{
		Time:  time.Unix(int64(s.stubTimeSec), int64(s.stubTimeNano)),
		Valid: true,
	}
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
				Remarks:        s.stubStringWrapper,
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
				Remarks:        s.stubStringWrapper,
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

			resp, err := s.server.Create(s.stubCtx, tc.req)
			s.Equal(tc.expectedResp, resp)
			s.Equal(tc.expectedErr, err)
		})
	}
}

func (s *testSuite) Test_List() {
	testCases := []struct {
		testName      string
		req           *apifoodv1.ListRequest
		usecaseDBRows []repository.ListFoodRow
		usecaseErr    error
		expectedResp  *apifoodv1.ListResponse
		expectedErr   error
	}{
		{
			testName: "usecase error",
			req: &apifoodv1.ListRequest{
				StartTimestamp: s.stubTimestampWrapper,
				EndTimestamp:   s.stubTimestampWrapper,
				Type:           s.stubStringWrapper,
				IntakeStatus:   s.stubStringWrapper,
				Feeder:         s.stubStringWrapper,
				Location:       s.stubStringWrapper,
			},
			usecaseDBRows: nil,
			usecaseErr:    errors.New(errors.BadRequestError, "bad request"),
			expectedResp:  nil,
			expectedErr:   errors.New(errors.BadRequestError, "bad request"),
		},
		{
			testName: "success",
			req: &apifoodv1.ListRequest{
				StartTimestamp: s.stubTimestampWrapper,
				EndTimestamp:   s.stubTimestampWrapper,
				Type:           s.stubStringWrapper,
				IntakeStatus:   s.stubStringWrapper,
				Feeder:         s.stubStringWrapper,
				Location:       s.stubStringWrapper,
			},
			usecaseDBRows: []repository.ListFoodRow{
				{
					ID:           99,
					Name:         "some food",
					Type:         s.stubPGText,
					IntakeStatus: s.stubPGText,
					Feeder:       s.stubPGText,
					Location:     s.stubPGText,
					Remarks:      s.stubPGText,
					CreatedAt:    s.stubPGTimestamptz,
					UpdatedAt:    s.stubPGTimestamptz,
				},
				{
					ID:           999,
					Name:         "some other food",
					Type:         s.stubPGText,
					IntakeStatus: s.stubPGText,
					Feeder:       s.stubPGText,
					Location:     s.stubPGText,
					Remarks:      s.stubPGText,
					CreatedAt:    s.stubPGTimestamptz,
					UpdatedAt:    s.stubPGTimestamptz,
				},
			},
			usecaseErr: nil,
			expectedResp: &apifoodv1.ListResponse{
				Food: []*apifoodv1.ListResponse_Row{
					{
						Id:           99,
						Name:         "some food",
						Type:         s.stubString,
						IntakeStatus: s.stubString,
						Feeder:       s.stubString,
						Location:     s.stubString,
						Remarks:      s.stubString,
						CreatedAt:    s.stubTimestampWrapper,
						UpdatedAt:    s.stubTimestampWrapper,
					},
					{
						Id:           999,
						Name:         "some other food",
						Type:         s.stubString,
						IntakeStatus: s.stubString,
						Feeder:       s.stubString,
						Location:     s.stubString,
						Remarks:      s.stubString,
						CreatedAt:    s.stubTimestampWrapper,
						UpdatedAt:    s.stubTimestampWrapper,
					},
				},
			},
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.testName, func() {
			s.mockUsecase.EXPECT().ListFood(
				mock.Anything,
				mock.AnythingOfType("*food.ListFoodParams"),
			).Return(tc.usecaseDBRows, tc.usecaseErr).Once()

			resp, err := s.server.List(s.stubCtx, tc.req)
			s.Equal(tc.expectedResp, resp)
			s.Equal(tc.expectedErr, err)
		})
	}
}
