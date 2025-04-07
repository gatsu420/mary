package auth_test

func (s *testSuite) Test_IssueToken() {
	testCases := []struct {
		testName    string
		expectedErr error
	}{
		{
			testName:    "successfully issued a token",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.testName, func() {
			signedToken, err := s.auth.IssueToken(s.username)
			s.NotEqual("", signedToken)
			s.Equal(tc.expectedErr, err)
		})
	}
}
