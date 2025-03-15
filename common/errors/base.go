package errors

import "google.golang.org/grpc/codes"

type Error struct {
	message  string
	category string
}

func (e *Error) Error() string {
	if e == nil {
		return ""
	}

	return e.message
}

func (e *Error) Category() string {
	if e == nil {
		return ""
	}

	return e.category
}

var grpcErrors = map[string]codes.Code{
	AuthError:           codes.Unauthenticated,
	ForbiddenError:      codes.PermissionDenied,
	InternalServerError: codes.Internal,
	BadRequestError:     codes.InvalidArgument,
	NotFoundError:       codes.NotFound,
}

func (e *Error) GRPCCode() codes.Code {
	if e == nil {
		return codes.OK
	}

	c, exists := grpcErrors[e.category]
	if !exists {
		return codes.Unknown
	}

	return c
}
