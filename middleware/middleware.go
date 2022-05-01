package middleware

import (
	"github.com/casdoor/casdoor-go-sdk/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware(c *gin.Context) {

	token := c.GetHeader("token")

	claims, err := auth.ParseJwtToken(token)
	if err != nil {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}
	c.Set("user-id", claims.Id)
	//c.SetCookie()
	c.Next()

}
