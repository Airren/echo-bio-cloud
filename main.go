package main

import (
	"github.com/gin-gonic/gin"

	"github.com/airren/echo-bio-backend/config"
	"github.com/airren/echo-bio-backend/dal"
	"github.com/airren/echo-bio-backend/docs"
	"github.com/airren/echo-bio-backend/middleware"
	"github.com/airren/echo-bio-backend/router"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Echo-Bio-Cloud
// @version 0.0.1
// @description Order Manager
// @contact.name Airren, Peto
// @contact.email renqiqiang@outlook.com, peto1
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host http://echo-bio.cn
// @BasePath /

//go:generate swag init
func main() {

	config.InitConfig()
	config.AuthInit()

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	r.MaxMultipartMemory = 8 << 20

	r.Use(middleware.CORSMiddleware, middleware.AuthMiddleware)
	apiV1 := r.Group("api/v1/")
	router.HealthCheck(apiV1)
	router.FileAPI(apiV1)
	router.StaticAPI(r)
	router.UserAPI(apiV1)

	router.AlgorithmAPI(apiV1)
	router.JobAPI(apiV1)

	err := dal.InitMySQL()
	if err != nil {
		panic(err)
	}

	docs.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	_ = r.Run()

}
