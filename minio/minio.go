package minio

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"go.uber.org/zap"
	"mime/multipart"

	"github.com/airren/echo-bio-backend/global"
	"github.com/airren/echo-bio-backend/model"
)

func UploadFileToMinio(ctx context.Context, fileInfo *model.File, file *multipart.FileHeader) (*model.File, error) {
	var (
		err error
	)
	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()
	//use org as bucketName
	_, err = global.MinioClient.PutObject(ctx, fileInfo.Org, fileInfo.MD5, src, file.Size, minio.PutObjectOptions{ContentType: "contentType"})
	if err != nil {
		global.Logger.Error("Upload fileInfo failed", zap.Error(err))
		return nil, err
	}
	global.Logger.Info("Successfully uploaded file: ", zap.String("name", fileInfo.Name))
	return fileInfo, err
}

//func GetFileUrl(ctx context.Context, bucketName string, fileName string, expires time.Duration) string {
//	//URL can have a maximum expiry of upto 7days
//	reqParams := make(url.Values)
//	fileUrl, err := global.MinioClient.PresignedGetObject(ctx, bucketName, fileName, expires, reqParams)
//	if err != nil {
//		global.Logger.Error(err.Error())
//		return ""
//	}
//	return fmt.Sprintf("%s", fileUrl)
//}

func DownloadObjectFromMinio(ctx context.Context, bucketName string, objectName string) (*minio.Object, error) {
	object, err := global.MinioClient.GetObject(ctx, bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		global.Logger.Error(
			fmt.Sprintf("Failed to download Object for bucketName: %s, objectName: %s", bucketName, objectName))
		return nil, err
	}
	return object, err
}

func BucketExist(ctx context.Context, bucketName string) (bool, error) {
	found, err := global.MinioClient.BucketExists(ctx, bucketName)
	if err != nil {
		global.Logger.Error("check bucket existence failed", zap.Error(err))
	}
	return found, err
}

func CreateBucket(ctx context.Context, bucketName string) error {
	location := "cn-east-1"
	err := global.MinioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := BucketExist(ctx, bucketName)
		if errBucketExists == nil && exists {
			global.Logger.Info(fmt.Sprintf("We already own %s\n", bucketName))
		} else {
			global.Logger.Error(fmt.Sprintf("Create Bucket %v failed", bucketName), zap.Error(err))
			return err
		}
	} else {
		global.Logger.Info(fmt.Sprintf("Successfully created %s\n", bucketName))
	}
	return err
}
