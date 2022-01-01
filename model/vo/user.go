package vo

import (
	"github.com/casdoor/casdoor-go-sdk/auth"
	"golang.org/x/oauth2"
)

type User struct {
	BaseVO
	auth.User
	Access string `json:"access"`
}

type TokenVO struct {
	BaseVO
	*oauth2.Token
}