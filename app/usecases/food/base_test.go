package food_test

import (
	"context"
	"testing"
	"time"

	"github.com/gatsu420/mary/app/usecases/food"
	mockrepository "github.com/gatsu420/mary/mocks/app/repository"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/suite"
)

type testSuite struct {
	suite.Suite
	mockRepo *mockrepository.MockQuerier
	usecase  food.Usecase

	ctx           context.Context
	pgText        pgtype.Text
	pgTimestamptz pgtype.Timestamptz
	pgInt4        pgtype.Int4
}

var (
	_ suite.TestingSuite   = (*testSuite)(nil)
	_ suite.SetupAllSuite  = (*testSuite)(nil)
	_ suite.SetupTestSuite = (*testSuite)(nil)
)

func (s *testSuite) SetupSuite() {
	s.ctx = context.Background()
	s.pgText = pgtype.Text{String: "test", Valid: true}
	s.pgTimestamptz = pgtype.Timestamptz{Time: time.Now(), Valid: true}
	s.pgInt4 = pgtype.Int4{Int32: 99, Valid: true}
}

func (s *testSuite) SetupTest() {
	s.mockRepo = mockrepository.NewMockQuerier(s.T())
	s.usecase = food.NewUsecase(s.mockRepo)
}

func Test(t *testing.T) {
	suite.Run(t, &testSuite{})
}
