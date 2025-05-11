package auth_test

import (
	"testing"

	"github.com/gatsu420/mary/app/auth"
	"github.com/gatsu420/mary/common/config"
	"github.com/stretchr/testify/suite"
)

type testSuite struct {
	suite.Suite
	auth auth.Auth

	config   *config.Config
	username string
}

var (
	_ suite.TestingSuite   = (*testSuite)(nil)
	_ suite.SetupAllSuite  = (*testSuite)(nil)
	_ suite.SetupTestSuite = (*testSuite)(nil)
)

func (s *testSuite) SetupSuite() {
	s.config, _ = config.New("../../.env.example")
	s.username = "es_teh_manis"
}

func (s *testSuite) SetupTest() {
	s.auth = auth.NewAuth(s.config)
}

func Test(t *testing.T) {
	suite.Run(t, &testSuite{})
}
