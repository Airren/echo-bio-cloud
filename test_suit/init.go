package test_suit

import "github.com/airren/echo-bio-backend/dal"

func TestInit() {
	err := dal.InitMySQL()
	if err != nil {
		panic(err)
	}
}
