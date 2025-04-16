package auth_test

import "github.com/gatsu420/mary/common/errors"

func (s *testSuite) Test_ValidateToken() {
	mockToken, _ := s.auth.IssueToken(s.username)
	testCases := []struct {
		testName       string
		signedToken    string
		expectedUserID string
		expectedErr    error
	}{
		{
			testName:       "token is expired on april 7 2025",
			signedToken:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDQwMjExNzcsImlhdCI6MTc0NDAyMTE2Miwic3ViIjoiY29jb19kdW1teSJ9.RvgdWBzgW3j-YqHQlBQtUegzqJ1eCdo6ES2vcNInEyA",
			expectedUserID: "",
			expectedErr:    errors.New(errors.AuthError, "unable to parse or verify token"),
		},
		{
			testName:       "token is tampered",
			signedToken:    "#there must be no additional string here# eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDQwMjExNzcsImlhdCI6MTc0NDAyMTE2Miwic3ViIjoiY29jb19kdW1teSJ9.RvgdWBzgW3j-YqHQlBQtUegzqJ1eCdo6ES2vcNInEyA",
			expectedUserID: "",
			expectedErr:    errors.New(errors.AuthError, "unable to parse or verify token"),
		},
		{
			testName:       "token is generated",
			signedToken:    mockToken,
			expectedUserID: "es_teh_manis",
			expectedErr:    nil,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.testName, func() {
			userID, err := s.auth.ValidateToken(tc.signedToken)
			s.Equal(tc.expectedUserID, userID)
			s.Equal(tc.expectedErr, err)
		})
	}
}
