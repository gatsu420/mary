package events_test

import (
	"context"
	"testing"

	"github.com/gatsu420/mary/app/usecases/events"
	mockcache "github.com/gatsu420/mary/mocks/app/cache"
	mockrepository "github.com/gatsu420/mary/mocks/app/repository"
	"github.com/stretchr/testify/suite"
)

type testSuite struct {
	suite.Suite
	mockRepo  *mockrepository.MockQuerier
	mockCache *mockcache.MockStorer
	usecase   events.Usecase

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
	s.mockRepo = mockrepository.NewMockQuerier(s.T())
	s.mockCache = mockcache.NewMockStorer(s.T())
	s.usecase = events.NewUsecase(s.mockRepo, s.mockCache)
}

func Test(t *testing.T) {
	suite.Run(t, &testSuite{})
}
