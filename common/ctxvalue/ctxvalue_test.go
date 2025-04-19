package ctxvalue_test

import (
	"context"
	"testing"

	"github.com/gatsu420/mary/common/ctxvalue"
	"github.com/stretchr/testify/assert"
)

func Test_SetAndGetUser(t *testing.T) {
	t.Run("set and get user context successfully", func(t *testing.T) {
		fakeUser := ctxvalue.User{
			UserID: "test_user",
		}
		ctx := ctxvalue.SetUser(context.Background(), fakeUser)
		user := ctxvalue.GetUser(ctx)
		assert.Equal(t, fakeUser.UserID, user.UserID)
	})
}
