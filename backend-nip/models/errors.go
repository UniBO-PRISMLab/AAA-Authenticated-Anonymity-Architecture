package models

import "errors"

type ErrorResponseModel struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var ErrorInternalServerErrorResponseModel = ErrorResponseModel{
	Code:    500,
	Message: "internal server error",
}

var ErrorBadRequestResponseModel = ErrorResponseModel{
	Code:    400,
	Message: "bad request",
}

var ErrorNotFoundResponseModel = ErrorResponseModel{
	Code:    404,
	Message: "not found",
}

var ErrorResponseModelWithMsg = func(code int, message string) ErrorResponseModel {
	return ErrorResponseModel{
		Code:    code,
		Message: message,
	}
}

var ErrorUserWithPIDNotFound = errors.New("no user found with the provided PID")

var ErrorInvalidPublicKey = errors.New("invalid public key")
var ErrorInvalidPublicKeyHeader = errors.New("invalid public key header")
var ErrorPKAlreadyAssociated = errors.New("this public key has been already associated with a PID")
var ErrorPublicKeyDecoding = errors.New("error while decoding public key")
var ErrorInvalidSignatureEncoding = errors.New("error while encoding signature")
var ErrorPIDSignatureVerification = errors.New("error while verifying PID signature")
var ErrorPACNotValid = errors.New("the provided PAC is expired or not valid for the given PID")
