package food_test

import (
	"context"
	"testing"
	"time"

	"github.com/gatsu420/mary/app/usecases/food"
	"github.com/gatsu420/mary/common/ctxvalue"
	mockcache "github.com/gatsu420/mary/mocks/app/cache"
	mockrepository "github.com/gatsu420/mary/mocks/app/repository"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/suite"
)

type testSuite struct {
	suite.Suite
	mockRepo  *mockrepository.MockQuerier
	mockCache *mockcache.MockStorer
	usecase   food.Usecase

	fakeUser ctxvalue.User
	ctx      context.Context

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
	s.fakeUser = ctxvalue.User{
		UserID: "test_user",
	}
	s.ctx = ctxvalue.SetUser(context.Background(), s.fakeUser)

	s.pgText = pgtype.Text{String: "test", Valid: true}
	s.pgTimestamptz = pgtype.Timestamptz{Time: time.Now(), Valid: true}
	s.pgInt4 = pgtype.Int4{Int32: 99, Valid: true}
}

func (s *testSuite) SetupTest() {
	s.mockRepo = mockrepository.NewMockQuerier(s.T())
	s.mockCache = mockcache.NewMockStorer(s.T())
	s.usecase = food.NewUsecase(s.mockRepo, s.mockCache)
}

func Test(t *testing.T) {
	suite.Run(t, &testSuite{})
}
