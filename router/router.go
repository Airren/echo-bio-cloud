package router

import (
	"github.com/airren/echo-bio-backend/controller"
	"github.com/gin-gonic/gin"
)

func ApiV_1(r *gin.Engine) {
}

func UserAPI(r *gin.Engine) {
	user := r.Group("api/user")
	user.GET("/login", controller.UserLogin)
	user.GET("/info", controller.UserInfo)
	user.PUT("/logout", controller.UserLogout)
}

func AlgorithmAPI(r *gin.Engine) {
	algo := r.Group("api/v1/algorithm")
	algo.POST("/create", controller.CreateAlgorithm)
	algo.PUT("/update", controller.UpdateAlgorithm)
	algo.GET("/list", controller.QueryAlgorithm)
}
