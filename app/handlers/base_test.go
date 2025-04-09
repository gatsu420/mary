package handlers_test

import (
	"context"
	"testing"
	"time"

	"github.com/gatsu420/mary/app/handlers"
	mockfood "github.com/gatsu420/mary/mocks/app/usecases/food"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/suite"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type testSuite struct {
	suite.Suite
	mockUsecase *mockfood.MockUsecase
	server      handlers.FoodServer

	stubCtx      context.Context
	stubTimeSec  int
	stubTimeNano int

	stubStringWrapper *wrapperspb.StringValue
	stubString        string
	stubPGText        pgtype.Text

	stubTimestampWrapper *timestamppb.Timestamp
	stubPGTimestamptz    pgtype.Timestamptz

	stubInt32Wrapper *wrapperspb.Int32Value
}

func (s *testSuite) SetupSuite() {
	s.stubCtx = context.Background()
	s.stubTimeSec = 1744208782
	s.stubTimeNano = 99

	s.stubStringWrapper = &wrapperspb.StringValue{Value: "test"}
	s.stubString = "test"
	s.stubPGText = pgtype.Text{String: "test", Valid: true}

	s.stubTimestampWrapper = &timestamppb.Timestamp{
		Seconds: int64(s.stubTimeSec),
		Nanos:   int32(s.stubTimeNano),
	}
	s.stubPGTimestamptz = pgtype.Timestamptz{
		Time:  time.Unix(int64(s.stubTimeSec), int64(s.stubTimeNano)),
		Valid: true,
	}

	s.stubInt32Wrapper = &wrapperspb.Int32Value{Value: 99}
}

func (s *testSuite) SetupTest() {
	s.mockUsecase = mockfood.NewMockUsecase(s.T())
	s.server = *handlers.NewFoodServer(s.mockUsecase)
}

func Test(t *testing.T) {
	suite.Run(t, &testSuite{})
}
