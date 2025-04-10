package errors_test

import (
	"testing"

	"github.com/gatsu420/mary/common/errors"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
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

func Test_Error(t *testing.T) {
	err := errors.New(errors.AuthError, "auth error")
	expectedErrMsg := "auth error"
	t.Run("successfully get error message", func(t *testing.T) {
		errMsg := err.Error()
		assert.Equal(t, expectedErrMsg, errMsg)
	})
}

func Test_GRPCCode(t *testing.T) {
	err := errors.New(errors.AuthError, "auth error")
	expectedGRPCCode := codes.Unauthenticated
	t.Run("successfully get GRPC code", func(t *testing.T) {
		grpcCode := err.GRPCCode()
		assert.Equal(t, expectedGRPCCode, grpcCode)
	})
}
