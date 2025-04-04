package auth_test

import (
	"testing"

	"github.com/gatsu420/mary/auth"
	"github.com/stretchr/testify/suite"
)

type testSuite struct {
	suite.Suite
	auth auth.Auth
}

func (s *testSuite) SetupSuite() {}

func (s *testSuite) SetupTest() {
	s.auth = auth.NewAuth("")
}

func Test(t *testing.T) {
	suite.Run(t, &testSuite{})
}
