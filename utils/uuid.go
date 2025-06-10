package utils

import "github.com/bwmarrin/snowflake"

var node *snowflake.Node

func init() {
	var err error
	node, err = snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}
}

func GenShortId() string {
	id := node.Generate()
	return id.Base36()
}

func GenInt64MsgId() int64 {
	id := node.Generate()
	return id.Int64()
}
