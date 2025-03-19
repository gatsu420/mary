package errors

import "google.golang.org/grpc/codes"

type errCategory int

const (
	AuthError errCategory = iota + 1
	ForbiddenError
	InternalServerError
	BadRequestError
	NotFoundError
)

type Error struct {
	category errCategory
	message  string
}

func (e *Error) Error() string {
	return e.message
}

var grpcErrors = map[errCategory]codes.Code{
	1: codes.Unauthenticated,
	2: codes.PermissionDenied,
	3: codes.Internal,
	4: codes.InvalidArgument,
	5: codes.NotFound,
}

func (e *Error) GRPCCode() codes.Code {
	return grpcErrors[e.category]
}

func New(category errCategory, msg string) *Error {
	if category == 0 {
		return nil
	}

	return &Error{
		category: category,
		message:  msg,
	}
}
