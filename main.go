package main

import (
	"github.com/airren/echo-bio-backend/config"
	"github.com/airren/echo-bio-backend/dal"
	"github.com/airren/echo-bio-backend/middleware"
	"github.com/airren/echo-bio-backend/router"
	"github.com/gin-gonic/gin"
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
	r.MaxMultipartMemory = 8 << 20

	r.Use(middleware.CORSMiddleware, middleware.AuthMiddleware)
	apiV1 := r.Group("api/v1/")
	router.FileAPI(apiV1)
	router.StaticAPI(r)
	router.UserAPI(apiV1)

	router.AlgorithmAPI(apiV1)
	router.JobAPI(apiV1)

	err := dal.InitMySQL()
	if err != nil {
		panic(err)
	}

	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.ApiV_1(r)
	_ = r.Run()

}
