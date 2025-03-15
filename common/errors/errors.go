package errors

const (
	AuthError           = "auth_error"
	ForbiddenError      = "forbidden_error"
	InternalServerError = "internal_server_error"
	BadRequestError     = "bad_request_error"
	NotFoundError       = "not_found_error"
)

func Auth(msg string) *Error {
	return &Error{
		message:  msg,
		category: AuthError,
	}
}

func Forbidden(msg string) *Error {
	return &Error{
		message:  msg,
		category: ForbiddenError,
	}
}

func InternalServer(msg string) *Error {
	return &Error{
		message:  msg,
		category: InternalServerError,
	}
}

func BadRequest(msg string) *Error {
	return &Error{
		message:  msg,
		category: BadRequestError,
	}
}

func NotFound(msg string) *Error {
	return &Error{
		message:  msg,
		category: NotFoundError,
	}
}
