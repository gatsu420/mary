package handlers_test

import (
	apifoodv1 "github.com/gatsu420/mary/api/gen/go/food/v1"
	"github.com/gatsu420/mary/common/errors"
	"github.com/gatsu420/mary/db/repository"
	"github.com/stretchr/testify/mock"
)

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
