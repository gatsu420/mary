package food_test

import (
	"context"
	"testing"
	"time"

	"github.com/gatsu420/mary/app/usecases/food"
	mockrepository "github.com/gatsu420/mary/mocks/db/repository"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/suite"
)

type testSuite struct {
	suite.Suite
	mockRepo *mockrepository.MockQuerier
	usecases food.Usecase

	ctx               context.Context
	mockPGText        pgtype.Text
	mockPGTimestamptz pgtype.Timestamptz
	mockPGInt4        pgtype.Int4
}

func (s *testSuite) SetupSuite() {
	s.ctx = context.Background()
	s.mockPGText = pgtype.Text{String: "test", Valid: true}
	s.mockPGTimestamptz = pgtype.Timestamptz{Time: time.Now(), Valid: true}
	s.mockPGInt4 = pgtype.Int4{Int32: 99, Valid: true}
}

func (s *testSuite) SetupTest() {
	s.mockRepo = mockrepository.NewMockQuerier(s.T())
	s.usecases = food.NewUsecase(s.mockRepo)
}

func Test(t *testing.T) {
	suite.Run(t, &testSuite{})
}
