package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/airren/echo-bio-backend/controller"
)

func HealthCheck(r *gin.RouterGroup) {
	r.GET("ping", func(context *gin.Context) {
		context.String(http.StatusOK, "pong")
	})
}
func StaticAPI(r *gin.Engine) {
	static := r.Group("api/static")
	static.StaticFS("/data", http.Dir("./static"))
}
func UserAPI(r *gin.RouterGroup) {

	user := r.Group("user")
	user.GET("/login", controller.UserLogin)
	user.GET("/info", controller.UserInfo)
	user.PUT("/logout", controller.UserLogout)
}

func AlgorithmAPI(r *gin.RouterGroup) {
	algo := r.Group("algorithm")
	algo.POST("/create", controller.CreateAlgorithm)
	algo.PUT("/update", controller.UpdateAlgorithm)
	algo.POST("/list", controller.QueryAlgorithm)

	group := algo.Group("group")
	group.POST("/create", controller.CreateAlgoGroup)
	group.PUT("/update", controller.UpdateAlgoGroup)
	group.GET("/list", controller.ListAlgoGroup)
	group.DELETE("/delete", controller.DeleteAlgoGroupById)
}

func JobAPI(r *gin.RouterGroup) {
	job := r.Group("job")
	job.GET("/list", controller.QueryJob)
	job.POST("/create", controller.CreateJob)
}

func FileAPI(r *gin.RouterGroup) {
	file := r.Group("file")
	file.POST("/create", controller.CreateFile)
	file.PUT("/update", controller.UpdateFile)
	file.POST("/upload", controller.UploadFile)
	file.GET("/list", controller.ListFile)
}
