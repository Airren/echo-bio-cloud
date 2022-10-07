package service

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/airren/echo-bio-backend/dal"
	"github.com/airren/echo-bio-backend/model"
	"github.com/airren/echo-bio-backend/model/req"
	"github.com/airren/echo-bio-backend/model/vo"
	"github.com/airren/echo-bio-backend/utils"
)

func UploadFile(c context.Context, req model.File) (*model.File, error) {
	if req.Id == 0 {
		req.Id = utils.GenerateId()
	}
	req.CreatedAt = time.Now()
	userId, err := utils.GetUserId(c)
	if err != nil && err.Error() != "userId not exist" {
		return nil, err
	}
	req.AccountId = userId
	org := utils.GetOrgFromCtx(c)
	req.Org = org
	file, err := dal.CreateFile(c, &req)
	return file, err
}

func UpdateFile(c context.Context, req req.FileReq) (*vo.FileVO, error) {
	file := FileToEntity(req)
	oldFile, err := dal.QueryFileById(c, file.Id)
	if err != nil {
		return nil, err
	}
	userId, err := utils.GetUserId(c)
	if err != nil {
		return nil, err
	}
	oldFile.UpdatedAt = time.Now()
	oldFile.AccountId = userId
	org := utils.GetOrgFromCtx(c)
	oldFile.Org = org
	if file.Name != "" {
		oldFile.Name = file.Name
	}
	if file.Description != "" {
		oldFile.Description = file.Description
	}
	f, err := dal.UpdateFile(c, oldFile)
	return FileToVO(*f), err
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
