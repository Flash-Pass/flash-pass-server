package encryptor

import (
	"crypto/md5"
	"fmt"
	"io"
)

type Handle struct{}

type IHandle interface {
	encryptMd5(data string) (string, error)
	PasswordEncrypt(password, salt string) (string, error)
}

func New() *Handle {
	return &Handle{}
}

func (e *Handle) encryptMd5(data string) (string, error) {
	h := md5.New()
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func (e *Handle) PasswordEncrypt(password, salt string) (string, error) {
	return e.encryptMd5(password + salt)
}

var _ IHandle = (*Handle)(nil)
