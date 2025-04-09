package handlers_test

import (
	apiauthv1 "github.com/gatsu420/mary/api/gen/go/auth/v1"
	"github.com/gatsu420/mary/common/errors"
	"github.com/stretchr/testify/mock"
)

func (s *testSuite) Test_IssueToken() {
	testCases := []struct {
		testName        string
		req             *apiauthv1.IssueTokenRequest
		usecaseErr      error
		authSignedToken string
		authErr         error
		expectedResp    *apiauthv1.IssueTokenResponse
		expectedErr     error
	}{
		{
			testName:     "usecase error",
			req:          &apiauthv1.IssueTokenRequest{Username: "test"},
			usecaseErr:   errors.New(errors.BadRequestError, "bad request"),
			expectedResp: nil,
			expectedErr:  errors.New(errors.BadRequestError, "bad request"),
		},
		{
			testName:        "failed when issuing token",
			req:             &apiauthv1.IssueTokenRequest{Username: "test"},
			usecaseErr:      nil,
			authSignedToken: "",
			authErr:         errors.New(errors.AuthError, "auth error"),
			expectedResp:    nil,
			expectedErr:     errors.New(errors.AuthError, "auth error"),
		},
		{
			testName:        "success",
			req:             &apiauthv1.IssueTokenRequest{Username: "test"},
			usecaseErr:      nil,
			authSignedToken: "some token",
			authErr:         nil,
			expectedResp:    &apiauthv1.IssueTokenResponse{SignedToken: "some token"},
			expectedErr:     nil,
		},
	}

	for _, tc := range testCases {
		s.mockUsersUsecase.EXPECT().CheckUserIsExisting(
			mock.Anything,
			mock.AnythingOfType("string"),
		).Return(tc.usecaseErr).Once()

		if tc.usecaseErr == nil {
			s.mockAuth.EXPECT().IssueToken(
				mock.Anything,
			).Return(tc.authSignedToken, tc.authErr).Once()
		}

		resp, err := s.authServer.IssueToken(s.stubCtx, tc.req)
		s.Equal(tc.expectedResp, resp)
		s.Equal(tc.expectedErr, err)
	}
}
