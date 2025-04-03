package users_test

import (
	"github.com/gatsu420/mary/common/errors"
	"github.com/stretchr/testify/mock"
)

func (s *testSuite) Test_CheckUserIsExisting() {
	testCases := []struct {
		testName       string
		repoIsExisting bool
		repoErr        error
		expectedErr    error
	}{
		{
			testName:       "repo error",
			repoIsExisting: false,
			repoErr:        errors.New(errors.InternalServerError, "DB error"),
			expectedErr:    errors.New(errors.InternalServerError, "DB failed to check if user is existing"),
		},
		{
			testName:       "user is not found",
			repoIsExisting: false,
			repoErr:        nil,
			expectedErr:    errors.New(errors.NotFoundError, "user not found"),
		},
		{
			testName:       "user is found",
			repoIsExisting: true,
			repoErr:        nil,
			expectedErr:    nil,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.testName, func() {
			s.mockRepo.EXPECT().CheckUserIsExisting(
				mock.Anything,
				mock.AnythingOfType("string"),
			).Return(tc.repoIsExisting, tc.repoErr).Once()

			err := s.usecase.CheckUserIsExisting(s.ctx, "test")
			s.Equal(err, tc.expectedErr)
		})
	}
}
