package workers_test

import (
	"context"

	"github.com/stretchr/testify/mock"
)

func (s *testSuite) Test_CreateMembershipRegistry() {
	// For now, can only simulate if runtime does not call log.Fatal()
	s.Run("create membership registry successfully", func() {
		s.mockAuthnUsecase.EXPECT().CreateMembershipRegistry(
			mock.Anything,
		).Return(nil).Once()

		s.worker.CreateMembershipRegistry(context.Background())

		s.mockAuthnUsecase.AssertCalled(s.T(), "CreateMembershipRegistry",
			mock.Anything,
		)
	})
}
