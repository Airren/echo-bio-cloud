package main

import (
	"github.com/airren/echo-bio-backend/config"
	"github.com/airren/echo-bio-backend/dal"
	middleware2 "github.com/airren/echo-bio-backend/middleware"
	"github.com/airren/echo-bio-backend/router"
	"github.com/gin-gonic/gin"

	//"github.com/swaggo/gin-swagger" // gin-swagger middleware
	//"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)

// @title Break Jail
// @version 0.0.1
// @description Order Manager
// @contact.name Airren
// @contact.email renqiang.luffy@bytedance.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {

	config.AuthInit()
	//r := gin.Default()

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	router.UserAPI(r)
	r.Use(middleware2.AuthMiddleware())
	err := dal.InitMySQL()
	if err != nil {
		panic(err)
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/health", func(c *gin.Context) {
		val := c.Query("val")
		c.String(http.StatusOK, "ok: "+val)
	})

	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.ApiV_1(r)
	_ = r.Run()

}
