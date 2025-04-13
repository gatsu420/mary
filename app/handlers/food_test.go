package handlers_test

import (
	apifoodv1 "github.com/gatsu420/mary/api/gen/go/food/v1"
	"github.com/gatsu420/mary/app/usecases/food"
	"github.com/gatsu420/mary/common/errors"
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
				Remarks:        s.stringWrapper,
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
				Remarks:        s.stringWrapper,
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

			resp, err := s.foodServer.Create(s.ctx, tc.req)
			s.Equal(tc.expectedResp, resp)
			s.Equal(tc.expectedErr, err)
		})
	}
}

func (s *testSuite) Test_List() {
	testCases := []struct {
		testName      string
		req           *apifoodv1.ListRequest
		usecaseDBRows []food.ListFoodRow
		usecaseErr    error
		expectedResp  *apifoodv1.ListResponse
		expectedErr   error
	}{
		{
			testName: "usecase error",
			req: &apifoodv1.ListRequest{
				StartTimestamp: s.timestampWrapper,
				EndTimestamp:   s.timestampWrapper,
				Type:           s.stringWrapper,
				IntakeStatus:   s.stringWrapper,
				Feeder:         s.stringWrapper,
				Location:       s.stringWrapper,
			},
			usecaseDBRows: nil,
			usecaseErr:    errors.New(errors.BadRequestError, "bad request"),
			expectedResp:  nil,
			expectedErr:   errors.New(errors.BadRequestError, "bad request"),
		},
		{
			testName: "success",
			req: &apifoodv1.ListRequest{
				StartTimestamp: s.timestampWrapper,
				EndTimestamp:   s.timestampWrapper,
				Type:           s.stringWrapper,
				IntakeStatus:   s.stringWrapper,
				Feeder:         s.stringWrapper,
				Location:       s.stringWrapper,
			},
			usecaseDBRows: []food.ListFoodRow{
				{
					ID:           99,
					Name:         "some food",
					Type:         s.pgText,
					IntakeStatus: s.pgText,
					Feeder:       s.pgText,
					Location:     s.pgText,
					Remarks:      s.pgText,
					CreatedAt:    s.pgTimestamptz,
					UpdatedAt:    s.pgTimestamptz,
				},
				{
					ID:           999,
					Name:         "some other food",
					Type:         s.pgText,
					IntakeStatus: s.pgText,
					Feeder:       s.pgText,
					Location:     s.pgText,
					Remarks:      s.pgText,
					CreatedAt:    s.pgTimestamptz,
					UpdatedAt:    s.pgTimestamptz,
				},
			},
			usecaseErr: nil,
			expectedResp: &apifoodv1.ListResponse{
				Food: []*apifoodv1.ListResponse_Row{
					{
						Id:           99,
						Name:         "some food",
						Type:         s.stringDummy,
						IntakeStatus: s.stringDummy,
						Feeder:       s.stringDummy,
						Location:     s.stringDummy,
						Remarks:      s.stringDummy,
						CreatedAt:    s.timestampWrapper,
						UpdatedAt:    s.timestampWrapper,
					},
					{
						Id:           999,
						Name:         "some other food",
						Type:         s.stringDummy,
						IntakeStatus: s.stringDummy,
						Feeder:       s.stringDummy,
						Location:     s.stringDummy,
						Remarks:      s.stringDummy,
						CreatedAt:    s.timestampWrapper,
						UpdatedAt:    s.timestampWrapper,
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

			resp, err := s.foodServer.List(s.ctx, tc.req)
			s.Equal(tc.expectedResp, resp)
			s.Equal(tc.expectedErr, err)
		})
	}
}

func (s *testSuite) Test_Get() {
	testCases := []struct {
		testName     string
		req          *apifoodv1.GetRequest
		usecaseDBRow *food.GetFoodRow
		usecaseErr   error
		expectedResp *apifoodv1.GetResponse
		expectedErr  error
	}{
		{
			testName:     "usecase error",
			req:          &apifoodv1.GetRequest{Id: 99},
			usecaseDBRow: &food.GetFoodRow{},
			usecaseErr:   errors.New(errors.BadRequestError, "bad request"),
			expectedResp: nil,
			expectedErr:  errors.New(errors.BadRequestError, "bad request"),
		},
		{
			testName: "success",
			req:      &apifoodv1.GetRequest{Id: 99},
			usecaseDBRow: &food.GetFoodRow{
				ID:           99,
				Name:         "some food",
				Type:         s.pgText,
				IntakeStatus: s.pgText,
				Feeder:       s.pgText,
				Location:     s.pgText,
				Remarks:      s.pgText,
				CreatedAt:    s.pgTimestamptz,
				UpdatedAt:    s.pgTimestamptz,
			},
			usecaseErr: nil,
			expectedResp: &apifoodv1.GetResponse{
				Id:           99,
				Name:         "some food",
				Type:         s.stringDummy,
				IntakeStatus: s.stringDummy,
				Feeder:       s.stringDummy,
				Location:     s.stringDummy,
				Remarks:      s.stringDummy,
				CreatedAt:    s.timestampWrapper,
				UpdatedAt:    s.timestampWrapper,
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

			resp, err := s.foodServer.Get(s.ctx, tc.req)
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
				Name:           s.stringWrapper,
				TypeId:         s.int32Wrapper,
				IntakeStatusId: s.int32Wrapper,
				FeederId:       s.int32Wrapper,
				LocationId:     s.int32Wrapper,
				Remarks:        s.stringWrapper,
				Id:             99,
			},
			usecaseErr:   errors.New(errors.BadRequestError, "bad request"),
			expectedResp: nil,
			expectedErr:  errors.New(errors.BadRequestError, "bad request"),
		},
		{
			testName: "success",
			req: &apifoodv1.UpdateRequest{
				Name:           s.stringWrapper,
				TypeId:         s.int32Wrapper,
				IntakeStatusId: s.int32Wrapper,
				FeederId:       s.int32Wrapper,
				LocationId:     s.int32Wrapper,
				Remarks:        s.stringWrapper,
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

			resp, err := s.foodServer.Update(s.ctx, tc.req)
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

			resp, err := s.foodServer.Delete(s.ctx, tc.req)
			s.Equal(tc.expectedResp, resp)
			s.Equal(tc.expectedErr, err)
		})
	}
}
