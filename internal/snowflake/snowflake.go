package snowflake

import "github.com/bwmarrin/snowflake"

type Handle struct {
	node *snowflake.Node
}

//go:generate mockgen -source=snowflake.go -destination=./mocks/snowflake_mock.go -package SnowflakeMocks
type IHandle interface {
	GetId() snowflake.ID
	GetUInt64Id() uint64
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

func (h *Handle) GetUInt64Id() uint64 {
	return uint64(h.GetId().Int64())
}

var _ IHandle = (*Handle)(nil)
