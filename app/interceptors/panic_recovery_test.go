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

func Test_RecoverPanic(t *testing.T) {
	ctx := context.Background()
	testCases := []struct {
		testName    string
		panicErr    any
		expectedErr error
	}{
		{
			testName:    "panic occured",
			panicErr:    errors.New(errors.InternalServerError, "some runtime error"),
			expectedErr: status.Errorf(codes.Internal, "internal server error: some runtime error"),
		},
		{
			testName:    "no panic occured",
			panicErr:    nil,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			handler := func(ctx context.Context, req any) (any, error) {
				if tc.panicErr != nil {
					panic(tc.panicErr)
				}

				return nil, nil
			}
			interceptor := interceptors.RecoverPanic()
			_, err := interceptor(ctx, nil, nil, handler)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}
