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

var ErrorUnableToCreateConnPool = errors.New("unable to create connection pool")
var ErrorUnableToConnectToEthClient = errors.New("unable to connect to Ethereum client")
var ErrorUnableToCreateUIPListener = errors.New("unable to create UIP listener")
var ErrorLoadTransactor = errors.New("unable to load transactor")
var ErrorSACSubmission = errors.New("error during sac submsission")
var ErrorSubscribeToLogs = errors.New("unable to subscribe to logs")
var ErrorSubscribtion = errors.New("subscription error")
var ErrorParseWordRequested = errors.New("failed to parse WordRequested event")
var ErrorParseSIDEncryptionRequested = errors.New("failed to parse SIDEncryptionRequested event")
var ErrorWordEncryption = errors.New("failed to encrypt word")
var ErrorRedundantWordEncryption = errors.New("failed to encrypt redundant word")
var ErrorPIDEncryption = errors.New("failed to encrypt PID")
var ErrorWordSubmission = errors.New("failed to submit encrypted word")
var ErrorRetrievingSIDRecord = errors.New("failed to retrieve SID record from blockchain")

var ErrorUserWithPIDNotFound = errors.New("no user found with the provided PID")
var ErrorInvalidPublicKey = errors.New("invalid public key")
var ErrorInvalidPublicKeyHeader = errors.New("invalid public key header")
var ErrorInvalidSymK = errors.New("invalid symmetric key")
var ErrorPKAlreadyAssociated = errors.New("this public key has been already associated with a PID")
var ErrorPublicKeyDecoding = errors.New("error while decoding public key")
var ErrorInvalidSignatureEncoding = errors.New("error while encoding signature")
var ErrorPIDSignatureVerification = errors.New("the provided signature is not valid for the given PID")
var ErrorSIDSignatureVerification = errors.New("the provided signature is not valid for the given SID")
var ErrorPACNotValid = errors.New("the provided PAC is expired or not valid for the given PID")
