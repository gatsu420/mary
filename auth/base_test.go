package auth_test

import (
	"testing"

	"github.com/gatsu420/mary/auth"
	"github.com/stretchr/testify/suite"
)

type testSuite struct {
	suite.Suite
	auth auth.Auth

	secret   string
	username string
}

func (s *testSuite) SetupSuite() {
	s.secret = "nasi_gila"
	s.username = "es_teh_manis"
}

func (s *testSuite) SetupTest() {
	s.auth = auth.NewAuth(s.secret)
}

func Test(t *testing.T) {
	suite.Run(t, &testSuite{})
}
