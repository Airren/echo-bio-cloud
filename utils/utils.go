package utils

import (
	"context"
	"errors"
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
func GenerateId() uint64 {
	// Generate a snowflake ID.
	id := node.Generate()
	return uint64(id.Int64())
}

func SetUserId(ctx context.Context, userId string) context.Context {
	return context.WithValue(ctx, "account-id", userId)
}

func GetUserId(ctx context.Context) (string, error) {
	userId := ctx.Value("account-id")
	if userId != nil {
		return userId.(string), nil
	}
	return "", errors.New("userId not exist")
}
