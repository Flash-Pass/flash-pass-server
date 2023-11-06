package fpstatus

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestErrNo(t *testing.T) {
	errNo := NewErrNo(1, "test")
	require.Equal(t, errNo.ErrCode, 1)
	require.Equal(t, errNo.ErrMsg, "test")

	errNo1 := errNo.WithMessage("another message")
	require.Equal(t, errNo.ErrCode, 1)
	require.Equal(t, errNo.ErrMsg, "test")
	require.Equal(t, errNo1.ErrMsg, "another message")
}
