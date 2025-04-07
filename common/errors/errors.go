package errors

import "google.golang.org/grpc/codes"

type ErrCategory int

const (
	AuthError ErrCategory = iota + 1
	ForbiddenError
	InternalServerError
	BadRequestError
	NotFoundError
)

type Error struct {
	Category ErrCategory
	Message  string
}

func (e *Error) Error() string {
	return e.Message
}

var grpcErrors = map[ErrCategory]codes.Code{
	1: codes.Unauthenticated,
	2: codes.PermissionDenied,
	3: codes.Internal,
	4: codes.InvalidArgument,
	5: codes.NotFound,
}

func (e *Error) GRPCCode() codes.Code {
	return grpcErrors[e.Category]
}

func New(category ErrCategory, msg string) *Error {
	if category == 0 {
		return nil
	}

	return &Error{
		Category: category,
		Message:  msg,
	}
}
