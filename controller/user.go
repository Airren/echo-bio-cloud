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
	//token := c.GetHeader("token")
	//claims, err := auth.ParseJwtToken(token)
	//if err != nil {
	//	bindRespWithStatus(c, http.StatusUnauthorized, nil, err)
	//	return
	//}

	//userVO := &vo.User{
	//	Access:      "admin",
	//	Name:        claims.Name,
	//	Avatar:      claims.Avatar,
	//	UserId:      claims.Id,
	//	Email:       claims.Email,
	//	Signature:   "",
	//	Title:       "",
	//	Group:       "",
	//	NotifyCount: "",
	//	UnreadCount: "",
	//	Country:     "",
	//	Address:     fmt.Sprint(claims.Address),
	//	Phone:       claims.Phone,
	//}

	baseVo := vo.BaseVO{}

	bindResp(c, &baseVo, nil)
}

func UserLogout(c *gin.Context) {

}
