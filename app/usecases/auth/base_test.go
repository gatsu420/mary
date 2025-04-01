package auth_test

import (
	"context"
	"testing"

	"github.com/gatsu420/mary/app/usecases/auth"
	mockrepository "github.com/gatsu420/mary/mocks/db/repository"
	"github.com/stretchr/testify/suite"
)

type testSuite struct {
	suite.Suite
	mockRepo *mockrepository.MockQuerier
	usecases auth.Usecases

	ctx       context.Context
	jwtSecret string
}

func (s *testSuite) SetupSuite() {
	s.ctx = context.Background()

	s.jwtSecret = "secret"
}

func (s *testSuite) SetupTest() {
	s.mockRepo = mockrepository.NewMockQuerier(s.T())
	s.usecases = auth.NewUsecases(s.jwtSecret, s.mockRepo)
}

func Test(t *testing.T) {
	suite.Run(t, &testSuite{})
}
