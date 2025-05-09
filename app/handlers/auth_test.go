package handlers_test

import (
	apiauthv1 "github.com/gatsu420/mary/api/gen/go/auth/v1"
	"github.com/gatsu420/mary/common/errors"
	"github.com/stretchr/testify/mock"
)

func (s *testSuite) Test_IssueToken() {
	testCases := []struct {
		testName            string
		req                 *apiauthv1.IssueTokenRequest
		registry            []int
		registryCreationErr error
		membershipCheckErr  error
		authSignedToken     string
		authErr             error
		expectedResp        *apiauthv1.IssueTokenResponse
		expectedErr         error
	}{
		{
			testName:            "failed when creating membership registry",
			req:                 &apiauthv1.IssueTokenRequest{Username: "test"},
			registry:            nil,
			registryCreationErr: errors.New(errors.InternalServerError, "some error"),
			expectedResp:        nil,
			expectedErr:         errors.New(errors.InternalServerError, "some error"),
		},
		{
			testName:            "user is not registered",
			req:                 &apiauthv1.IssueTokenRequest{Username: "test"},
			registry:            []int{0, 0, 1},
			registryCreationErr: nil,
			membershipCheckErr:  errors.New(errors.AuthError, "some error"),
			expectedResp:        nil,
			expectedErr:         errors.New(errors.AuthError, "some error"),
		},
		{
			testName:            "failed when issuing token",
			req:                 &apiauthv1.IssueTokenRequest{Username: "test"},
			registry:            []int{0, 0, 1},
			registryCreationErr: nil,
			membershipCheckErr:  nil,
			authSignedToken:     "",
			authErr:             errors.New(errors.AuthError, "auth error"),
			expectedResp:        nil,
			expectedErr:         errors.New(errors.AuthError, "auth error"),
		},
		{
			testName:            "success",
			req:                 &apiauthv1.IssueTokenRequest{Username: "test"},
			registry:            []int{0, 0, 1},
			registryCreationErr: nil,
			membershipCheckErr:  nil,
			authSignedToken:     "some token",
			authErr:             nil,
			expectedResp:        &apiauthv1.IssueTokenResponse{SignedToken: "some token"},
			expectedErr:         nil,
		},
	}

	for _, tc := range testCases {
		s.mockAuth.EXPECT().CreateMembershipRegistry().Return(tc.registry, tc.registryCreationErr).Once()

		if tc.registryCreationErr == nil {
			s.mockAuth.EXPECT().CheckMembership(
				mock.AnythingOfType("[]int"),
				mock.AnythingOfType("string"),
			).Return(tc.membershipCheckErr).Once()
		}

		if tc.registryCreationErr == nil && tc.membershipCheckErr == nil {
			s.mockAuth.EXPECT().IssueToken(
				mock.Anything,
			).Return(tc.authSignedToken, tc.authErr).Once()
		}

		resp, err := s.authServer.IssueToken(s.ctx, tc.req)
		s.Equal(tc.expectedResp, resp)
		s.Equal(tc.expectedErr, err)
	}
}
