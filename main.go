package main

import (
	"github.com/airren/echo-bio-backend/service"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/airren/echo-bio-backend/config"
	"github.com/airren/echo-bio-backend/dal"
	"github.com/airren/echo-bio-backend/docs"
	"github.com/airren/echo-bio-backend/middleware"
	"github.com/airren/echo-bio-backend/minio"
	"github.com/airren/echo-bio-backend/router"
)

// @title Echo-Bio-Cloud
// @version 0.0.1
// @description Order Manager
// @contact.name Airren, Peto
// @contact.email renqiqiang@outlook.com, peto1
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host http://echo-bio.cn
// @BasePath /api/v1/
//
//go:generate swag init
func main() {

	config.InitConfig()
	config.AuthInit()

	log.Infof("echo-bio start successfully...")
	//go executor.Run(context.TODO())
	//executor.KubeInitializer()

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	r.MaxMultipartMemory = 8 << 20

	r.Use(middleware.CORSMiddleware, middleware.AuthMiddleware)
	apiV1 := r.Group("api/v1/")
	router.HealthCheck(apiV1)
	router.FileAPI(apiV1)
	router.InternalAPI(apiV1)
	router.StaticAPI(r)
	router.UserAPI(apiV1)

	router.AlgorithmAPI(apiV1)
	router.JobAPI(apiV1)

	err := dal.InitMySQL()
	if err != nil {
		panic(err)
	}
	err = minio.InitMinio()
	if err != nil {
		panic(err)
	}

	go service.ConsumerJob()

	docs.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	_ = r.Run()
}
