package vo

import (
	"golang.org/x/oauth2"
)

type User struct {
	BaseVO
	Access      string `json:"access"`
	Name        string `json:"name"`
	Avatar      string `json:"avatar"`
	UserId      string `json:"userid"`
	Email       string `json:"email"`
	Signature   string `json:"signature"`
	Title       string `json:"title"`
	Group       string `json:"group"`
	NotifyCount string `json:"notifycount"`
	UnreadCount string `json:"unreadcount"`
	Country     string `json:"country"`
	Address     string `json:"address"`
	Phone       string `json:"phone"`
}

type TokenVO struct {
	BaseVO
	*oauth2.Token
}
