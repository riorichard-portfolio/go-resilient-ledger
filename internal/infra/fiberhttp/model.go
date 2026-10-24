package fiberhttp

type ErrorResponse struct {
	Message string `json:"message"`
	Errors map[string]string `json:"errors"`
}

var InternalServerError = ErrorResponse{
	Message: "internal server error",
}
