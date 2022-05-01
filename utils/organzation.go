package utils

import (
	"context"

	"github.com/gin-gonic/gin"
)

func GetOrgFromCtx(ctx context.Context) string {
	org := ctx.Value("org")
	if org == nil || org == "" {
		return "alarm"
	}
	return org.(string)
}

func GetCtx(c *gin.Context) context.Context {
	ctx := context.Background()
	org := c.GetHeader("org")
	userId, _ := c.Get("user-id")

	ctx = context.WithValue(ctx, "org", org)
	return context.WithValue(ctx, "user-id", userId)
}
