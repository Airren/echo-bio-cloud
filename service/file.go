package service

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"io"
	"mime/multipart"
	"strconv"
	"time"

	"go.uber.org/zap"

	"github.com/airren/echo-bio-backend/dal"
	"github.com/airren/echo-bio-backend/global"
	minioClient "github.com/airren/echo-bio-backend/minio"
	"github.com/airren/echo-bio-backend/model"
	"github.com/airren/echo-bio-backend/model/req"
	"github.com/airren/echo-bio-backend/model/vo"
	"github.com/airren/echo-bio-backend/utils"
)

func UploadFile(c context.Context, fh *multipart.FileHeader, visibility int) (*vo.FileVO, error) {

	src, err := fh.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	Md5 := fileToMd5(src)
	fileInfo, err := dal.QueryFileByMd5(c, Md5)
	if err != nil && err.Error() != "record not found" {
		global.Logger.Error("check file md5 failed", zap.Error(err))
		return nil, err
	}
	if fileInfo.Id != 0 {
		return FileToVO(fileInfo), err
	}

	// if file not exist, create a new file info and put to the oss
	userId, err := utils.GetUserId(c)
	if err != nil {
		return nil, err
	}
	org := utils.GetOrgFromCtx(c)
	fileInfo.RecordMeta = model.RecordMeta{
		Id:        utils.GenerateId(),
		AccountId: userId,
		Org:       org,
		CreatedAt: time.Now(),
		CreatedBy: userId,
	}
	fileInfo.MD5 = Md5
	fileInfo.Name, fileInfo.FileType = minioClient.GetFileNameType(fh.Filename)
	fileInfo.Visibility = visibility

	err = minioClient.UploadFileToMinio(c, fileInfo.Org, fileInfo.MD5, fh)
	if err != nil {
		global.Logger.Error("upload file failed", zap.Error(err))
		return nil, err
	}

	fileInfo.SetURLPath()
	fileInfo, err = dal.InsertFileInfo(c, fileInfo)
	if err != nil {
		return nil, err
	}
	return FileToVO(fileInfo), err
}

func UpdateFileInfo(c context.Context, fileReq *req.FileReq) (*vo.FileVO, error) {
	fileInfo := FileToEntity(fileReq)
	fileInfo.SetURLPath()

	file, err := dal.UpdateFileInfo(c, fileInfo)
	return FileToVO(file), err
}

func DownloadFileById(c context.Context, fileId uint64) (fileInfo *model.File, bytes []byte, err error) {
	// make user is valid, have access to the file
	fileInfo, err = dal.QueryFileById(c, fileId)
	if err != nil {
		return
	}
	object, err := minioClient.DownloadObjectFromMinio(c, fileInfo.Org, fileInfo.MD5)
	bytes, err = io.ReadAll(object)
	return
}

func ListFileInfo(c context.Context, req *req.FileReq) (fileVOs []*vo.FileVO, pageInfo *model.PageInfo, err error) {
	file := FileToEntity(req)
	files, err := dal.ListFiles(c, file, &req.PageInfo)
	for _, v := range files {
		fileVOs = append(fileVOs, FileToVO(v))
	}
	return fileVOs, &req.PageInfo, err
}

func QueryFileByIds(c context.Context, req *req.IdsReq) (fileVOs []*vo.FileVO, err error) {
	files, err := dal.QueryFileByIds(c, req.IdsToInt())
	if err != nil {
		return nil, err
	}
	for _, v := range files {
		fileVOs = append(fileVOs, FileToVO(v))
	}
	return fileVOs, err
}

func DeleteFileByIds(c context.Context, idsReq *req.IdsReq) (err error) {
	return dal.DeleteFileByIds(c, idsReq.IdsToInt())
}

func FileToEntity(req *req.FileReq) *model.File {
	var Id int64
	if req.Id != "" {
		Id, _ = strconv.ParseInt(req.Id, 10, 64)
	}
	return &model.File{
		RecordMeta:  model.RecordMeta{Id: uint64(Id)},
		Name:        req.Name,
		FileType:    req.FileType,
		Visibility:  req.Visibility,
		Description: req.Description,
	}
}

func FileToVO(file *model.File) *vo.FileVO {
	return &vo.FileVO{
		RecordMeta:  RecordMetaToVO(file.RecordMeta),
		Name:        file.Name,
		FileType:    file.FileType,
		Visibility:  file.Visibility,
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
