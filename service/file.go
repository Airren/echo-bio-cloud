package service

import (
	"context"
	"fmt"
	"mime/multipart"
	"strconv"
	"time"

	minioClient "github.com/airren/echo-bio-backend/minio"
	"github.com/minio/minio-go/v7"

	"github.com/airren/echo-bio-backend/dal"
	"github.com/airren/echo-bio-backend/global"
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
	fileInfo, err := minioClient.UploadFileToMinio(c, req, file)
	fileInfo.AccountId = userId
	_, err = dal.InsertFileMetaInfo(c, fileInfo)
	if err != nil {
		return nil, err
	}
	return FileToVO(*fileInfo), err
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
