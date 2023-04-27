package test_suit

import (
	"github.com/airren/echo-bio-backend/config"
	"github.com/airren/echo-bio-backend/dal"
)

func TestInit() {
	config.InitConfig()
	config.AuthInit()
	err := dal.InitMySQL()
	if err != nil {
		panic(err)
	}
}
