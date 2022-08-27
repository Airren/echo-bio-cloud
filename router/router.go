package router

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/airren/echo-bio-backend/controller"
)

func ApiV_1(r *gin.Engine) {
}

func StaticAPI(r *gin.Engine) {
	static := r.Group("api/static")
	static.StaticFS("/data", http.Dir("./static"))
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

func JobAPI(r *gin.Engine) {
	job := r.Group("api/v1/job")
	job.GET("/list", controller.QueryJob)
	job.POST("/create", controller.CreateJob)
}
