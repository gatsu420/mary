package errors_test

import (
	"testing"

	"github.com/gatsu420/mary/common/errors"
	"github.com/stretchr/testify/assert"
)

func Test_New(t *testing.T) {
	category := errors.AuthError
	msg := "auth error"
	expectedErr := &errors.Error{
		Category: 1,
		Message:  "auth error",
	}
	t.Run("successfully generate new error", func(t *testing.T) {
		err := errors.New(category, msg)
		assert.Equal(t, expectedErr, err)
	})
}
