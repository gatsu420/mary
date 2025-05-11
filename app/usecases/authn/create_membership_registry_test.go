package authn_test

import (
	"github.com/gatsu420/mary/common/errors"
	"github.com/stretchr/testify/mock"
)

func (s *testSuite) Test_CreateMembershipRegistry() {
	testCases := []struct {
		testName               string
		repoUsers              []string
		repoListUsersErr       error
		cacheCreateRegistryErr error
		expectedErr            error
	}{
		{
			testName:         "repo failed to list users",
			repoUsers:        nil,
			repoListUsersErr: errors.New(errors.InternalServerError, "some error"),
			expectedErr:      errors.New(errors.InternalServerError, "DB failed to list users"),
		},
		{
			testName:               "cache failed to create membership registry",
			repoUsers:              []string{"testUserA", "testUserB"},
			repoListUsersErr:       nil,
			cacheCreateRegistryErr: errors.New(errors.InternalServerError, "some error"),
			expectedErr:            errors.New(errors.InternalServerError, "cache failed to create membership registry"),
		},
		{
			testName:               "success",
			repoUsers:              []string{"testUserA", "testUserB"},
			repoListUsersErr:       nil,
			cacheCreateRegistryErr: nil,
			expectedErr:            nil,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.testName, func() {
			s.mockRepo.EXPECT().ListUsers(
				mock.Anything,
			).Return(tc.repoUsers, tc.repoListUsersErr).Once()

			if tc.repoListUsersErr == nil {
				s.mockAuth.EXPECT().CreateMembershipRegistry(
					mock.AnythingOfType("[]string"),
				).Return([]string{}).Once()

				s.mockCache.EXPECT().CreateMembershipRegistry(
					mock.Anything,
					mock.AnythingOfType("cache.CreateMembershipRegistryParams"),
				).Return(tc.cacheCreateRegistryErr).Once()
			}

			err := s.usecase.CreateMembershipRegistry(s.ctx)
			s.Equal(tc.expectedErr, err)
		})
	}
}
