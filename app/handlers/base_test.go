package handlers_test

import (
	"context"
	"testing"
	"time"

	"github.com/gatsu420/mary/app/handlers"
	mockfood "github.com/gatsu420/mary/mocks/app/usecases/food"
	mockusers "github.com/gatsu420/mary/mocks/app/usecases/users"
	mockauth "github.com/gatsu420/mary/mocks/auth"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/suite"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type testSuite struct {
	suite.Suite
	mockAuth         *mockauth.MockAuth
	mockUsersUsecase *mockusers.MockUsecase
	mockFoodUsecase  *mockfood.MockUsecase
	authServer       *handlers.AuthServer
	foodServer       *handlers.FoodServer

	ctx      context.Context
	timeSec  int
	timeNano int

	stringWrapper *wrapperspb.StringValue
	stringDummy   string
	pgText        pgtype.Text

	timestampWrapper *timestamppb.Timestamp
	pgTimestamptz    pgtype.Timestamptz

	int32Wrapper *wrapperspb.Int32Value
}

func (s *testSuite) SetupSuite() {
	s.ctx = context.Background()
	s.timeSec = 1744208782
	s.timeNano = 99

	s.stringWrapper = &wrapperspb.StringValue{Value: "test"}
	s.stringDummy = "test"
	s.pgText = pgtype.Text{String: "test", Valid: true}

	s.timestampWrapper = &timestamppb.Timestamp{
		Seconds: int64(s.timeSec),
		Nanos:   int32(s.timeNano),
	}
	s.pgTimestamptz = pgtype.Timestamptz{
		Time:  time.Unix(int64(s.timeSec), int64(s.timeNano)),
		Valid: true,
	}

	s.int32Wrapper = &wrapperspb.Int32Value{Value: 99}
}

func (s *testSuite) SetupTest() {
	s.mockAuth = mockauth.NewMockAuth(s.T())
	s.mockUsersUsecase = mockusers.NewMockUsecase(s.T())
	s.mockFoodUsecase = mockfood.NewMockUsecase(s.T())
	s.authServer = handlers.NewAuthServer(s.mockAuth, s.mockUsersUsecase)
	s.foodServer = handlers.NewFoodServer(s.mockFoodUsecase)
}

func Test(t *testing.T) {
	suite.Run(t, &testSuite{})
}
