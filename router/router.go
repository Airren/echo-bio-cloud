package router

import (
	"github.com/airren/echo-bio-backend/controller"
	"github.com/gin-gonic/gin"
)

func ApiV_1(r *gin.Engine) {

	task := r.Group("api/task")

	task.GET("/:id", controller.GetOrderById)
	task.POST("/create", controller.CreateOrder)
	task.PUT("/update", controller.UpdateOrder)
	task.POST("/list", controller.QueryOrders)

}

func UserAPI(r *gin.Engine) {
	user := r.Group("api/user")
	user.GET("/login", controller.UserLogin)
	user.GET("/info", controller.UserInfo)
	user.PUT("/logout", controller.UserLogout)
}
