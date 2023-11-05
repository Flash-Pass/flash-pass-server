package fpstatus

import (
	"fmt"
)

type ErrNo struct {
	ErrCode int
	ErrMsg  string
}

func (e *ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

func NewErrNo(code int, msg string) *ErrNo {
	return &ErrNo{code, msg}
}

func (e *ErrNo) WithMessage(msg string) *ErrNo {
	return NewErrNo(e.ErrCode, msg)
}

func ConvertErr(err error) *ErrNo {
	return UnexpectedError.WithMessage(err.Error())
}
