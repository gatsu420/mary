package authn_test

import (
	"context"
	"testing"

	"github.com/gatsu420/mary/app/usecases/authn"
	mockauth "github.com/gatsu420/mary/mocks/app/auth"
	mockcache "github.com/gatsu420/mary/mocks/app/cache"
	mockrepository "github.com/gatsu420/mary/mocks/app/repository"
	"github.com/stretchr/testify/suite"
)

type testSuite struct {
	suite.Suite
	mockAuth  *mockauth.MockAuth
	mockRepo  *mockrepository.MockQuerier
	mockCache *mockcache.MockStorer
	usecase   authn.Usecase

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
	s.mockAuth = mockauth.NewMockAuth(s.T())
	s.mockRepo = mockrepository.NewMockQuerier(s.T())
	s.mockCache = mockcache.NewMockStorer(s.T())
	s.usecase = authn.NewUsecase(s.mockAuth, s.mockRepo, s.mockCache)
}

func Test(t *testing.T) {
	suite.Run(t, &testSuite{})
}
