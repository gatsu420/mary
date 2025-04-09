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
			s.mockFoodUsecase.EXPECT().CreateFood(
				mock.Anything,
				mock.AnythingOfType("*food.CreateFoodParams"),
			).Return(tc.usecaseErr).Once()

			resp, err := s.foodServer.Create(s.stubCtx, tc.req)
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
			s.mockFoodUsecase.EXPECT().ListFood(
				mock.Anything,
				mock.AnythingOfType("*food.ListFoodParams"),
			).Return(tc.usecaseDBRows, tc.usecaseErr).Once()

			resp, err := s.foodServer.List(s.stubCtx, tc.req)
			s.Equal(tc.expectedResp, resp)
			s.Equal(tc.expectedErr, err)
		})
	}
}

func (s *testSuite) Test_Get() {
	testCases := []struct {
		testName     string
		req          *apifoodv1.GetRequest
		usecaseDBRow repository.GetFoodRow
		usecaseErr   error
		expectedResp *apifoodv1.GetResponse
		expectedErr  error
	}{
		{
			testName:     "usecase error",
			req:          &apifoodv1.GetRequest{Id: 99},
			usecaseDBRow: repository.GetFoodRow{},
			usecaseErr:   errors.New(errors.BadRequestError, "bad request"),
			expectedResp: nil,
			expectedErr:  errors.New(errors.BadRequestError, "bad request"),
		},
		{
			testName: "success",
			req:      &apifoodv1.GetRequest{Id: 99},
			usecaseDBRow: repository.GetFoodRow{
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
			usecaseErr: nil,
			expectedResp: &apifoodv1.GetResponse{
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
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.testName, func() {
			s.mockFoodUsecase.EXPECT().GetFood(
				mock.Anything,
				mock.AnythingOfType("int32"),
			).Return(tc.usecaseDBRow, tc.usecaseErr).Once()

			resp, err := s.foodServer.Get(s.stubCtx, tc.req)
			s.Equal(tc.expectedResp, resp)
			s.Equal(tc.expectedErr, err)
		})
	}
}

func (s *testSuite) Test_Update() {
	testCases := []struct {
		testName     string
		req          *apifoodv1.UpdateRequest
		usecaseErr   error
		expectedResp *apifoodv1.UpdateResponse
		expectedErr  error
	}{
		{
			testName: "usecase error",
			req: &apifoodv1.UpdateRequest{
				Name:           s.stubStringWrapper,
				TypeId:         s.stubInt32Wrapper,
				IntakeStatusId: s.stubInt32Wrapper,
				FeederId:       s.stubInt32Wrapper,
				LocationId:     s.stubInt32Wrapper,
				Remarks:        s.stubStringWrapper,
				Id:             99,
			},
			usecaseErr:   errors.New(errors.BadRequestError, "bad request"),
			expectedResp: nil,
			expectedErr:  errors.New(errors.BadRequestError, "bad request"),
		},
		{
			testName: "success",
			req: &apifoodv1.UpdateRequest{
				Name:           s.stubStringWrapper,
				TypeId:         s.stubInt32Wrapper,
				IntakeStatusId: s.stubInt32Wrapper,
				FeederId:       s.stubInt32Wrapper,
				LocationId:     s.stubInt32Wrapper,
				Remarks:        s.stubStringWrapper,
				Id:             99,
			},
			usecaseErr:   nil,
			expectedResp: &apifoodv1.UpdateResponse{},
			expectedErr:  nil,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.testName, func() {
			s.mockFoodUsecase.EXPECT().UpdateFood(
				mock.Anything,
				mock.AnythingOfType("*food.UpdateFoodParams"),
			).Return(tc.usecaseErr).Once()

			resp, err := s.foodServer.Update(s.stubCtx, tc.req)
			s.Equal(tc.expectedResp, resp)
			s.Equal(tc.expectedErr, err)
		})
	}
}

func (s *testSuite) Test_Delete() {
	testCases := []struct {
		testName     string
		req          *apifoodv1.DeleteRequest
		usecaseErr   error
		expectedResp *apifoodv1.DeleteResponse
		expectedErr  error
	}{
		{
			testName:     "usecase error",
			req:          &apifoodv1.DeleteRequest{Id: 99},
			usecaseErr:   errors.New(errors.BadRequestError, "bad request"),
			expectedResp: nil,
			expectedErr:  errors.New(errors.BadRequestError, "bad request"),
		},
		{
			testName:     "success",
			req:          &apifoodv1.DeleteRequest{Id: 99},
			usecaseErr:   nil,
			expectedResp: &apifoodv1.DeleteResponse{},
			expectedErr:  nil,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.testName, func() {
			s.mockFoodUsecase.EXPECT().DeleteFood(
				mock.Anything,
				mock.AnythingOfType("int32"),
			).Return(tc.usecaseErr).Once()

			resp, err := s.foodServer.Delete(s.stubCtx, tc.req)
			s.Equal(tc.expectedResp, resp)
			s.Equal(tc.expectedErr, err)
		})
	}
}
