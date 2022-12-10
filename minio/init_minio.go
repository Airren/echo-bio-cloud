package minio

import (
	"fmt"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

	"github.com/airren/echo-bio-backend/config"
	"github.com/airren/echo-bio-backend/global"
)

func InitMinio() error {
	var (
		err error
	)
	fmt.Print(config.Conf)
	// Initialize minio client object
	minioClient, err := minio.New(config.Conf.MinioConf.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.Conf.MinioConf.AccessKey, config.Conf.MinioConf.AccessSecret, ""),
		Secure: config.Conf.MinioConf.UseSSL,
	})

	global.MinioClient = minioClient
	log.Printf("successfully establish minio client")
	log.Printf("%#v\n", minioClient) // minioClient is now set up
	return err
}
