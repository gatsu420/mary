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
	usecase  auth.Usecase

	ctx    context.Context
	secret interface{}
}

func (s *testSuite) SetupSuite() {
	s.ctx = context.Background()
}

func (s *testSuite) SetupTest() {
	s.mockRepo = mockrepository.NewMockQuerier(s.T())
	s.usecase = auth.NewUsecase(s.secret, s.mockRepo)
}

func Test(t *testing.T) {
	suite.Run(t, &testSuite{})
}
