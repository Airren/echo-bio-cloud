package utils

import (
	"fmt"

	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node

func init() {
	var err error
	node, err = snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		return
	}

}
func GenerateId() int64 {
	// Generate a snowflake ID.
	id := node.Generate()
	return id.Int64()
}
