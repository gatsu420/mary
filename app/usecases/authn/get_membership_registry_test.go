package authn_test

import (
	"github.com/gatsu420/mary/common/errors"
	"github.com/stretchr/testify/mock"
)

func (s *testSuite) Test_GetMembershipRegistry() {
	testCases := []struct {
		testName            string
		cacheRegistry       []string
		cacheGetRegistryErr error
		expectedRegistry    []string
		expectedErr         error
	}{
		{
			testName:            "cache failed to get membership registry",
			cacheRegistry:       nil,
			cacheGetRegistryErr: errors.New(errors.InternalServerError, "some error"),
			expectedRegistry:    nil,
			expectedErr:         errors.New(errors.InternalServerError, "cache failed to get membership registry"),
		},
		{
			testName:            "success",
			cacheRegistry:       []string{"1", "1", "0", "0"},
			cacheGetRegistryErr: nil,
			expectedRegistry:    []string{"1", "1", "0", "0"},
			expectedErr:         nil,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.testName, func() {
			s.mockCache.EXPECT().GetMembershipRegistry(
				mock.Anything,
			).Return(tc.cacheRegistry, tc.cacheGetRegistryErr).Once()

			registry, err := s.usecase.GetMembershipRegistry(s.ctx)
			s.Equal(tc.expectedRegistry, registry)
			s.Equal(tc.expectedErr, err)
		})
	}
}
