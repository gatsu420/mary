package mathint_test

import (
	"testing"

	"github.com/gatsu420/mary/common/mathint"
	"github.com/stretchr/testify/assert"
)

func Test_Abs(t *testing.T) {
	testCases := []struct {
		testName       string
		x              int
		expectedResult int
	}{
		{
			testName:       "negative int",
			x:              99,
			expectedResult: 99,
		},
		{
			testName:       "positive int",
			x:              -99,
			expectedResult: 99,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			res := mathint.Abs(tc.x)
			assert.Equal(t, tc.expectedResult, res)
		})
	}
}
