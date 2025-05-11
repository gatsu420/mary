package workers_test

import (
	"context"
	"testing"

	"github.com/gatsu420/mary/app/workers"
	mockauthn "github.com/gatsu420/mary/mocks/app/usecases/authn"
	mockevents "github.com/gatsu420/mary/mocks/app/usecases/events"
	"github.com/stretchr/testify/suite"
)

type testSuite struct {
	suite.Suite
	mockAuthnUsecase *mockauthn.MockUsecase
	mockUsecase      *mockevents.MockUsecase
	worker           workers.Worker

	ctx context.Context
}

var (
	_ suite.TestingSuite   = (*testSuite)(nil)
	_ suite.SetupAllSuite  = (*testSuite)(nil)
	_ suite.SetupTestSuite = (*testSuite)(nil)
)

func (s *testSuite) SetupSuite() {
	s.ctx = context.Background()
}

func (s *testSuite) SetupTest() {
	s.mockAuthnUsecase = mockauthn.NewMockUsecase(s.T())
	s.mockUsecase = mockevents.NewMockUsecase(s.T())
	s.worker = workers.New(s.mockAuthnUsecase, s.mockUsecase)
}

func Test(t *testing.T) {
	suite.Run(t, &testSuite{})
}
