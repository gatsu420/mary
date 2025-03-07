package usecases

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ErrInternal() error {
	return status.Error(codes.Internal, "internal error")
}

func ErrRemovedRecord() error {
	return status.Error(codes.NotFound, "record is not found")
}
