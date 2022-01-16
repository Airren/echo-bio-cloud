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

func GetOrgCtx(c *gin.Context) context.Context {
	ctx := context.Background()
	org := c.GetHeader("org")
	return context.WithValue(ctx, "org", org)
}
