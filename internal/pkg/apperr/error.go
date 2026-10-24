package apperr

type AppErr struct {
	Code    int
	Message string
}

func (e *AppErr) Error() string { return e.Message }

func NotFoundErr(msg string) *AppErr {
	return &AppErr{
		Code:    404,
		Message: msg,
	}
}

func ForbiddenErr(msg string) *AppErr {
	return &AppErr{
		Code:    403,
		Message: msg,
	}
}

func BadRequestErr(msg string) *AppErr {
	return &AppErr{
		Code:    400,
		Message: msg,
	}
}

func UnauthorizedErr(msg string) *AppErr {
	return &AppErr{
		Code:    401,
		Message: msg,
	}
}

var tooManyRequestErr = AppErr {
	Code: 429,
	Message: "too many request",
}

func TooManyRequestErr() *AppErr {
	return &tooManyRequestErr
}