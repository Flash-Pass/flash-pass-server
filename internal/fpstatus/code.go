package fpstatus

const (
	SuccessCode int = iota
	UnexpectedErrorCode
	ParseTokenErrorCode
	ParseParametersErrorCode
	ParametersValidateErrorCode
	SystemErrorCode
)

var (
	Success                 = NewErrNo(SuccessCode, "success")
	UnexpectedError         = NewErrNo(UnexpectedErrorCode, "unexpected error")
	ParseTokenError         = NewErrNo(ParseTokenErrorCode, "parse token error")
	ParseParametersError    = NewErrNo(ParseParametersErrorCode, "parse parameters error")
	ParametersValidateError = NewErrNo(ParametersValidateErrorCode, "parameters validate error")
	SystemError             = NewErrNo(SystemErrorCode, "service error")
	// book
	BookExistError  = NewErrNo(10001, "book exist")
	BookDeleteError = NewErrNo(10002, "can not delete book created by others")
)
