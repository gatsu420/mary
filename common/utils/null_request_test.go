package utils_test

import (
	"testing"

	"github.com/gatsu420/mary/common/utils"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func Test_NullStringWrapperToPGText(t *testing.T) {
	testCases := []struct {
		testName       string
		mockWrapper    *wrapperspb.StringValue
		expectedPGText pgtype.Text
	}{
		{
			testName:    "wrapper is nil",
			mockWrapper: nil,
			expectedPGText: pgtype.Text{
				Valid: false,
			},
		},
		{
			testName: "wrapper is not nil",
			mockWrapper: &wrapperspb.StringValue{
				Value: "anna & elsa",
			},
			expectedPGText: pgtype.Text{
				String: "anna & elsa",
				Valid:  true,
			},
		},
	}

	for _, tc := range testCases {
		pgText := utils.NullStringWrapperToPGText(tc.mockWrapper)
		assert.Equal(t, tc.expectedPGText, pgText)
	}
}

func Test_NullInt32WrapperToPGInt4(t *testing.T) {
	testCases := []struct {
		testname       string
		mockWrapper    *wrapperspb.Int32Value
		expectedPGInt4 pgtype.Int4
	}{
		{
			testname:    "wrapper is nil",
			mockWrapper: nil,
			expectedPGInt4: pgtype.Int4{
				Valid: false,
			},
		},
		{
			testname: "wrapper is not nil",
			mockWrapper: &wrapperspb.Int32Value{
				Value: 99,
			},
			expectedPGInt4: pgtype.Int4{
				Int32: 99,
				Valid: true,
			},
		},
	}

	for _, tc := range testCases {
		pgInt4 := utils.NullInt32WrapperToPGInt4(tc.mockWrapper)
		assert.Equal(t, tc.expectedPGInt4, pgInt4)
	}
}
