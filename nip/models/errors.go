package models

type ErrorResponseModel struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var ErrorInternalServerErrorResponseModel = ErrorResponseModel{
	Code:    500,
	Message: "Internal Server Error",
}

var ErrorBadRequestResponseModel = ErrorResponseModel{
	Code:    400,
	Message: "Bad Request",
}
