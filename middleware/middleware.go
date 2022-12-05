package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
)

func AuthMiddleware(c *gin.Context) {
	// Debug log
	log.Println(c.Request.URL.String())

	if strings.HasPrefix(c.Request.URL.String(), "/api/v1/file/upload") ||
		strings.HasPrefix(c.Request.URL.String(), "/api/static") ||
		strings.HasPrefix(c.Request.URL.String(), "/api/v1/ping") ||
		strings.HasPrefix(c.Request.URL.String(), "/swagger") ||
		strings.HasPrefix(c.Request.URL.String(), "/api/v1/user/login") {
		c.Next()
		// Debug log
	        log.Println("pass directly")
		return
	}
	// Debug log
	log.Println("need authentication")

	token := c.GetHeader("token")
	if token == "" {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	claims, err := casdoorsdk.ParseJwtToken(token)
	if err != nil {
		log.Println("parse token failed:", err)
		c.AbortWithStatus(http.StatusForbidden)
		return
	}
	c.Set("user-id", claims.Id)
	//c.SetCookie()
	c.Next()
	return
}

func CORSMiddleware(c *gin.Context) {

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Headers", "Content-Type,x-jwt-token,token, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}

	c.Next()
	return
}
