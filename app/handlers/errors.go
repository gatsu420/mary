package handlers

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ErrAlreadyRemoved() error {
	return status.Error(codes.Internal, "the record is already removed")
}
