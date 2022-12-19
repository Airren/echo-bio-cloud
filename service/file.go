package service

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"go.uber.org/zap"
	"io"
	"mime/multipart"
	"strconv"
	"time"

	"github.com/minio/minio-go/v7"

	"github.com/airren/echo-bio-backend/dal"
	"github.com/airren/echo-bio-backend/global"
	minioClient "github.com/airren/echo-bio-backend/minio"
	"github.com/airren/echo-bio-backend/model"
	"github.com/airren/echo-bio-backend/model/req"
	"github.com/airren/echo-bio-backend/model/vo"
	"github.com/airren/echo-bio-backend/utils"
)

func UploadFile(c context.Context, req *model.File, file *multipart.FileHeader) (*vo.FileVO, error) {
	if req.Id == 0 {
		req.Id = utils.GenerateId()
	}
	req.CreatedAt = time.Now()
	userId, err := utils.GetUserId(c)
	if err != nil && err.Error() != "userId not exist" {
		return nil, err
	}
	org := utils.GetOrgFromCtx(c)
	req.Org = org
	exist, err := minioClient.BucketExist(c, org)
	if err != nil {
		return nil, err
	}
	if !exist {
		global.Logger.Info(fmt.Sprintf("Bucket:%s does not exist,trying to create", org))
		err := minioClient.CreateBucket(c, org)
		if err != nil {
			return nil, err
		}
	}
	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()
	Md5 := fileToMd5(src)
	req.MD5 = Md5
	total, err := dal.CheckIfFileExist(c, Md5)
	if err != nil {
		global.Logger.Error("check file md5 failed", zap.Error(err))
		return nil, err
	}
	if total == 0 {
		req, err = minioClient.UploadFileToMinio(c, req, file)
		if err != nil {
			global.Logger.Error("upload file failed", zap.Error(err))
			return nil, err
		}
	}
	req.AccountId = userId
	_, err = dal.InsertFileMetaInfo(c, req)
	if err != nil {
		return nil, err
	}
	return FileToVO(*req), err
}

func DownloadFileById(c context.Context, fileId uint64) (*minio.Object, error) {
	fileInfo, err := dal.QueryFileById(c, fileId)
	if err != nil {
		return nil, err
	}
	file, err := minioClient.DownloadObjectFromMinio(c, fileInfo.Org, fileInfo.Name)
	return file, err
}

func FileToEntity(req req.FileReq) *model.File {
	var Id int64
	if req.Id != "" {
		Id, _ = strconv.ParseInt(req.Id, 10, 64)
	}
	return &model.File{
		RecordMeta:  model.RecordMeta{Id: uint64(Id)},
		Name:        req.Name,
		Description: req.Description,
		URLPath:     req.URLPath,
	}
}

func FileToVO(file model.File) *vo.FileVO {
	return &vo.FileVO{
		Id:          fmt.Sprint(file.Id),
		Name:        file.Name,
		Description: file.Description,
		URLPath:     file.URLPath,
	}
}

func fileToMd5(file multipart.File) string {
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		global.Logger.Error("failed to convert file to Md5", zap.Error(err))
	}
	return hex.EncodeToString(hash.Sum(nil))
}
