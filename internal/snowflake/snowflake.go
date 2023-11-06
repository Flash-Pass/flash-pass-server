package snowflake

import "github.com/bwmarrin/snowflake"

type Handle struct {
	node *snowflake.Node
}

type IHandle interface {
	GetId() snowflake.ID
}

func NewHandle(node int64) *Handle {
	n, err := snowflake.NewNode(node)
	if err != nil {
		panic(err)
	}
	return &Handle{
		node: n,
	}
}

func (h *Handle) GetId() snowflake.ID {
	return h.node.Generate()
}

var _ IHandle = (*Handle)(nil)
