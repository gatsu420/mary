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
		wrapper        *wrapperspb.StringValue
		expectedPGText pgtype.Text
	}{
		{
			testName: "wrapper is nil",
			wrapper:  nil,
			expectedPGText: pgtype.Text{
				Valid: false,
			},
		},
		{
			testName: "wrapper is not nil",
			wrapper: &wrapperspb.StringValue{
				Value: "anna & elsa",
			},
			expectedPGText: pgtype.Text{
				String: "anna & elsa",
				Valid:  true,
			},
		},
	}

	for _, tc := range testCases {
		pgText := utils.NullStringWrapperToPGText(tc.wrapper)
		assert.Equal(t, tc.expectedPGText, pgText)
	}
}

func Test_NullInt32WrapperToPGInt4(t *testing.T) {
	testCases := []struct {
		testname       string
		wrapper        *wrapperspb.Int32Value
		expectedPGInt4 pgtype.Int4
	}{
		{
			testname: "wrapper is nil",
			wrapper:  nil,
			expectedPGInt4: pgtype.Int4{
				Valid: false,
			},
		},
		{
			testname: "wrapper is not nil",
			wrapper: &wrapperspb.Int32Value{
				Value: 99,
			},
			expectedPGInt4: pgtype.Int4{
				Int32: 99,
				Valid: true,
			},
		},
	}

	for _, tc := range testCases {
		pgInt4 := utils.NullInt32WrapperToPGInt4(tc.wrapper)
		assert.Equal(t, tc.expectedPGInt4, pgInt4)
	}
}
