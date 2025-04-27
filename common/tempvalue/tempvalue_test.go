package tempvalue_test

import (
	"testing"

	"github.com/gatsu420/mary/common/tempvalue"
	"github.com/stretchr/testify/assert"
)

func Test_SetAndGetUserID(t *testing.T) {
	t.Run("set and get user ID successfully", func(t *testing.T) {
		fakeValue := tempvalue.NewValue()
		fakeUserID := "fakeUserID"

		fakeValue.SetUserID(fakeUserID)
		userID := fakeValue.GetUserID()
		assert.Equal(t, fakeUserID, userID)
	})
}
