package auth_test

import "github.com/gatsu420/mary/common/errors"

func (s *testSuite) Test_CheckMembership() {
	testCases := []struct {
		testName    string
		userName    string
		expectedErr error
	}{
		{
			testName:    "user is not found",
			userName:    "some_non_existent_user",
			expectedErr: errors.New(errors.AuthError, "user is not found"),
		},
		{
			testName:    "success",
			userName:    "testUserA",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.testName, func() {
			registry := s.auth.CreateMembershipRegistry([]string{"testUserA", "testUserB"})
			err := s.auth.CheckMembership(registry, tc.userName)
			s.Equal(tc.expectedErr, err)
		})
	}
}
