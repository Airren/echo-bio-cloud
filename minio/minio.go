package minio

import (
	"context"
	"fmt"
	"strings"

	"github.com/minio/minio-go/v7"
	"go.uber.org/zap"
	"mime/multipart"

	"github.com/airren/echo-bio-backend/global"
)

func UploadFileToMinio(ctx context.Context, bucket, objectName string, fh *multipart.FileHeader) (
	err error) {
	// check if bucket exist
	exist, err := BucketExist(ctx, bucket)
	if err != nil {
		return err
	}
	if !exist {
		global.Logger.Info(fmt.Sprintf("Bucket:%s does not exist,trying to create", bucket))
		err := CreateBucket(ctx, bucket)
		if err != nil {
			return err
		}
	}

	// put file to the bucket
	src, err := fh.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	contentType := GetContentType(GetFileType(fh.Filename))
	_, err = global.MinioClient.PutObject(ctx, bucket, objectName,
		src, fh.Size, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		global.Logger.Error("Upload fileInfo failed", zap.Error(err))
		return err
	}
	global.Logger.Info("Successfully uploaded file: ", zap.String("name", fh.Filename))
	return err
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

func GetFileType(name string) (fileType string) {
	elems := strings.Split(name, ".")
	if len(elems) > 1 {
		return elems[len(elems)-1]
	}
	return ""
}

func GetContentType(filetype string) string {
	switch filetype {
	case "jpeg":
		fallthrough
	case "jpg":
		return "image/jpeg"
	case "png":
		return "image/png"
	case "csv":
		return "text/plain"
	default:
		return "application/octet-stream"
	}
}
