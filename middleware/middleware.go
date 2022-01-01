package middleware

import (
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		//token := c.GetHeader("token")
		//
		//claims, err := auth.ParseJwtToken(token)
		//if err != nil {
		//	err = c.AbortWithError(405, fmt.Errorf("use err %s", err))
		//	return
		//}
		//fmt.Println(claims)
		//c.SetCookie()
		c.Next()
	}

}
