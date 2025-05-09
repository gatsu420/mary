package auth_test

import (
	"testing"

	"github.com/gatsu420/mary/app/auth"
	"github.com/gatsu420/mary/common/config"
	mockrepository "github.com/gatsu420/mary/mocks/app/repository"
	"github.com/stretchr/testify/suite"
)

type testSuite struct {
	suite.Suite
	auth auth.Auth

	config   *config.Config
	query    *mockrepository.MockQuerier
	username string
}

var (
	_ suite.TestingSuite   = (*testSuite)(nil)
	_ suite.SetupAllSuite  = (*testSuite)(nil)
	_ suite.SetupTestSuite = (*testSuite)(nil)
)

func (s *testSuite) SetupSuite() {
	s.config, _ = config.New("../../.env.example")
	s.query = mockrepository.NewMockQuerier(s.T())
	s.username = "es_teh_manis"
}

func (s *testSuite) SetupTest() {
	s.auth = auth.NewAuth(s.config, s.query)
}

func Test(t *testing.T) {
	suite.Run(t, &testSuite{})
}
