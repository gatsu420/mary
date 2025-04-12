package interceptors_test

import (
	"context"
	"testing"

	"github.com/gatsu420/mary/app/interceptors"
	"github.com/gatsu420/mary/common/errors"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Test_ResponseError(t *testing.T) {
	ctx := context.Background()
	testCases := []struct {
		testName    string
		handlerErr  error
		expectedErr error
	}{
		{
			testName:    "error occured",
			handlerErr:  errors.New(errors.ForbiddenError, "some permission error"),
			expectedErr: status.Error(codes.PermissionDenied, "some permission error"),
		},
		{
			testName:    "no error occured",
			handlerErr:  nil,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			handler := func(ctx context.Context, req any) (any, error) {
				return nil, tc.handlerErr
			}
			interceptor := interceptors.ResponseError()
			_, err := interceptor(ctx, nil, nil, handler)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}
