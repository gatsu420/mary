package auth_test

import (
	"github.com/gatsu420/mary/app/usecases/auth"
	"github.com/gatsu420/mary/common/errors"
)

func (s *testSuite) Test_IssueToken() {
	testCases := []struct {
		testName    string
		username    string
		secret      interface{}
		expectedErr error
	}{
		{
			testName:    "secret is not []byte",
			username:    "test",
			secret:      "test",
			expectedErr: errors.New(errors.AuthError, "unable to sign JWT"),
		},
		{
			testName:    "success",
			username:    "test",
			secret:      []byte("test"),
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.testName, func() {
			s.usecase = auth.NewUsecase(tc.secret, s.mockRepo)

			signedToken, err := s.usecase.IssueToken(tc.username)
			s.Equal(tc.expectedErr, err)
			if tc.expectedErr == nil {
				s.NotEqual("", signedToken)
			} else {
				s.Equal("", signedToken)
			}
		})
	}
}
