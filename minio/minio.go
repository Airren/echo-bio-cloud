package minio

import (
	"context"
	"fmt"
	"mime/multipart"
	"strings"

	"github.com/minio/minio-go/v7"
	log "github.com/sirupsen/logrus"
)

func UploadFileToMinio(ctx context.Context, bucket, objectName string, fh *multipart.FileHeader) (
	err error) {
	// check if bucket exist
	exist, err := BucketExist(ctx, bucket)
	if err != nil {
		return err
	}
	if !exist {
		log.Infof("Bucket:%s does not exist,trying to create", bucket)
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
	_, fileType := GetFileNameType(fh.Filename)
	contentType := GetContentType(fileType)
	_, err = MinioClient.PutObject(ctx, bucket, objectName,
		src, fh.Size, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Errorf("Upload fileInfo failed")
		return err
	}
	log.Infof("Successfully uploaded file: ")
	return err
}

//func GetFileUrl(ctx context.Context, bucketName string, fileName string, expires time.Duration) string {
//	//URL can have a maximum expiry of upto 7days
//	reqParams := make(url.Values)
//	fileUrl, err := MinioClient.PresignedGetObject(ctx, bucketName, fileName, expires, reqParams)
//	if err != nil {
//		Logger.Error(err.Error())
//		return ""
//	}
//	return fmt.Sprintf("%s", fileUrl)
//}

func DownloadObjectFromMinio(ctx context.Context, bucketName string, objectName string) (*minio.Object, error) {
	object, err := MinioClient.GetObject(ctx, bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		log.Errorf("Failed to download Object for bucketName: %s, objectName: %s", bucketName, objectName)
		return nil, err
	}
	return object, err
}

func BucketExist(ctx context.Context, bucketName string) (bool, error) {
	found, err := MinioClient.BucketExists(ctx, bucketName)
	if err != nil {
		log.Errorf("check bucket existence failed")
	}
	return found, err
}

func CreateBucket(ctx context.Context, bucketName string) error {
	location := "cn-east-1"
	err := MinioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := BucketExist(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Infof("We already own %s\n", bucketName)
		} else {
			log.Errorf("Create Bucket %v failed", bucketName)
			return err
		}
	} else {
		log.Infof("Successfully created %s\n", bucketName)
	}
	return err
}

func GetFileNameType(nameType string) (name, fileType string) {
	elems := strings.Split(nameType, ".")
	if len(elems) > 1 {
		fileType = elems[len(elems)-1]
		return strings.TrimRight(nameType, fmt.Sprintf(".%s", fileType)), fileType
	}
	return nameType, ""
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
