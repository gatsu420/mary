package workers_test

import (
	"context"
	"testing"

	"github.com/gatsu420/mary/app/workers"
	mockevents "github.com/gatsu420/mary/mocks/app/usecases/events"
	"github.com/stretchr/testify/suite"
)

type testSuite struct {
	suite.Suite
	mockUsecase *mockevents.MockUsecase
	worker      workers.Worker

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
	s.mockUsecase = mockevents.NewMockUsecase(s.T())
	s.worker = workers.New(s.mockUsecase)
}

func Test(t *testing.T) {
	suite.Run(t, &testSuite{})
}
