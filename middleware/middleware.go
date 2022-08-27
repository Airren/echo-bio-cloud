package middleware

import (
	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware(c *gin.Context) {

	token := c.GetHeader("token")

	claims, err := casdoorsdk.ParseJwtToken(token)
	if err != nil {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}
	c.Set("user-id", claims.Id)
	//c.SetCookie()
	c.Next()

}
