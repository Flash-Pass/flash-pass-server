package generator

import (
	"crypto/rand"
	"fmt"
	"io"
)

type Generator struct{}

type IGenerator interface {
	GenSalt() (string, error)
}

func New() *Generator {
	return &Generator{}
}

func (g *Generator) GetSalt() (string, error) {
	b := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, b)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", b), nil
}

func (g *Generator) GenSalt() (string, error) {
	b := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, b)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", b), nil
}

var _ IGenerator = (*Generator)(nil)
