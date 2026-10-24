package bizerr

type BizErr struct {
	Code    int
	Message string
}

func (be *BizErr) Error() string { return be.Message }

func NotFoundErr(msg string) *BizErr {
	return &BizErr{
		Code:    404,
		Message: msg,
	}
}

func ForbiddenErr(msg string) *BizErr {
	return &BizErr{
		Code:    403,
		Message: msg,
	}
}

func BadRequestErr(msg string) *BizErr {
	return &BizErr{
		Code:    400,
		Message: msg,
	}
}

func UnauthorizedErr(msg string) *BizErr {
	return &BizErr{
		Code:    401,
		Message: msg,
	}
}
