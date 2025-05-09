package interceptors_test

import (
	"context"
	"testing"

	"github.com/gatsu420/mary/app/auth"
	"github.com/gatsu420/mary/app/interceptors"
	"github.com/gatsu420/mary/common/config"
	"github.com/gatsu420/mary/common/errors"
	mockrepository "github.com/gatsu420/mary/mocks/app/repository"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func Test_ValidateToken(t *testing.T) {
	config, _ := config.New("../../.env.example")
	query := mockrepository.NewMockQuerier(t)
	auth := auth.NewAuth(config, query)
	signedToken, _ := auth.IssueToken("testuser")

	testCases := []struct {
		testName    string
		ctx         context.Context
		info        *grpc.UnaryServerInfo
		expectedErr error
	}{
		{
			testName: "called public method",
			ctx:      context.Background(),
			info: &grpc.UnaryServerInfo{
				FullMethod: "/auth.v1.AuthService/IssueToken",
			},
			expectedErr: nil,
		},
		{
			testName: "token is not provided",
			ctx: metadata.NewIncomingContext(context.Background(), metadata.MD{
				"authorization": []string{""},
			}),
			info: &grpc.UnaryServerInfo{
				FullMethod: "/some/method/behind/auth/layer",
			},
			expectedErr: errors.New(errors.InternalServerError, "token is not provided"),
		},
		{
			testName: "token doesn't pass validation",
			ctx: metadata.NewIncomingContext(context.Background(), metadata.MD{
				"authorization": []string{"some.dummy.token"},
			}),
			info: &grpc.UnaryServerInfo{
				FullMethod: "/some/method/behind/auth/layer",
			},
			expectedErr: errors.New(errors.AuthError, "unable to parse or verify token"),
		},
		{
			testName: "token passes validation",
			ctx: metadata.NewIncomingContext(context.Background(), metadata.MD{
				"authorization": []string{signedToken},
			}),
			info: &grpc.UnaryServerInfo{
				FullMethod: "/some/method/behind/auth/layer",
			},
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			handler := func(ctx context.Context, req any) (any, error) {
				return nil, nil
			}
			interceptor := interceptors.ValidateToken(auth)
			_, err := interceptor(tc.ctx, nil, tc.info, handler)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}
