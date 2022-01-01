package controller

import (
	"github.com/airren/echo-bio-backend/model/vo"
	"net/http"

	"github.com/casdoor/casdoor-go-sdk/auth"
	"github.com/gin-gonic/gin"
)

func UserLogin(c *gin.Context) {
	code := c.Query("code")
	state := c.Query("state")

	token, err := auth.GetOAuthToken(code, state)
	data := &vo.TokenVO{}

	if err != nil {
		bindRespWithStatus(c, http.StatusUnauthorized, data, err)
		return
	}
	data.Token = token
	bindResp(c, data, nil)
}

func UserInfo(c *gin.Context) {
	token := c.GetHeader("token")
	claims, err := auth.ParseJwtToken(token)
	data := &vo.User{}
	if err != nil {
		bindRespWithStatus(c, http.StatusUnauthorized, data, err)
		return
	}
	data.User = claims.User
	data.Access = "user"
	bindResp(c, data, nil)
}

func UserLogout(c *gin.Context) {

}
