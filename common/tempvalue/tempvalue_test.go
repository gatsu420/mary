package tempvalue_test

import (
	"testing"

	"github.com/gatsu420/mary/common/tempvalue"
	"github.com/stretchr/testify/assert"
)

func Test_SetAndGetUserID(t *testing.T) {
	t.Run("set and get userID successfully", func(t *testing.T) {
		fakeUserID := "fakeUserID"

		tempvalue.SetUserID(fakeUserID)
		userID := tempvalue.GetUserID()
		assert.Equal(t, fakeUserID, userID)
	})
}
