package snowflake

import (
	"github.com/bwmarrin/snowflake"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSnowflake(t *testing.T) {
	t.Run("test for new snowflake handle", func(t *testing.T) {
		_, err := snowflake.NewNode(1)
		require.NoError(t, err)
	})

	t.Run("get a snowflake id", func(t *testing.T) {
		handle := NewHandle(1)
		handle.GetId()
	})
}
