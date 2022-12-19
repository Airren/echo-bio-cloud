package controller

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/airren/echo-bio-backend/dal"
	"github.com/airren/echo-bio-backend/global"
	"github.com/airren/echo-bio-backend/model"
	"github.com/airren/echo-bio-backend/model/req"
	"github.com/airren/echo-bio-backend/model/vo"
	"github.com/airren/echo-bio-backend/service"
	"github.com/airren/echo-bio-backend/utils"
)

// UploadFile godoc
//
//	@Summary		Upload a file
//	@Description	Upload a file
//	@Tags			file
//	@Accept			json
//	@Produce		json
//	@Param			file	formData	file	true	"FILE"
//	@Success		200	{object}	vo.FileVO
//	@Failure		400	{object}	vo.BaseVO
//	@Failure		500	{object}	vo.BaseVO
//	@Router			/file/update/ [put]
func UploadFile(c *gin.Context) {
	fileInfo := model.File{}

	//file, err := c.FormFile("file")
	file, err := c.FormFile("file")
	if err != nil {
		bindRespWithStatus(c, http.StatusBadRequest, nil, errors.New("file not be empty"))
		return
	}
	ctx := utils.GetCtx(c)
	global.Logger.Info(file.Filename)
	fileInfo.Id = utils.GenerateId()
	fileInfo.Name = file.Filename
	userId, err := utils.GetUserId(c)
	fileInfo.AccountId = userId
	f, err := service.UploadFile(ctx, &fileInfo, file)
	if err != nil {
		bindRespWithStatus(c, http.StatusInternalServerError, nil, err)
	}
	bindResp(c, vo.FileVO{Id: fmt.Sprint(f.Id)}, err)
}

// DownloadFileById godoc
//
//	@Summary		Download a file
//	@Description	Download by file ID
//	@Tags			file
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"FILE ID"	Format(octet-stream)
//	@Success		200	{object}	vo.BaseVO
//	@Failure		400	{object}	vo.BaseVO
//	@Failure		404	{object}	vo.BaseVO
//	@Failure		500	{object}	vo.BaseVO
//	@Router			/file/download/{id} [post]
func DownloadFileById(c *gin.Context) {
	ctx := utils.GetCtx(c)
	fileId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		bindRespWithStatus(c, http.StatusBadRequest, nil, err)
		return
	}
	fileSrc, err := service.DownloadFileById(ctx, fileId)
	if err != nil {
		bindRespWithStatus(c, http.StatusInternalServerError, nil, err)
		return
	}
	fileInfo, err := dal.QueryFileById(ctx, fileId)
	fileSrc.Stat()
	if err != nil {
		bindRespWithStatus(c, http.StatusNotFound, nil, err)
		return
	}
	//todo unable to display chinese
	fileName := fileInfo.Name
	extraHeaders := map[string]string{
		"Content-Disposition": "attachment; filename=" + fileName,
		"Cache-Control":       "no-cache",
	}
	c.DataFromReader(http.StatusOK, -1, "application/octet-stream", fileSrc, extraHeaders)
}

// ListFileInfo godoc
//
//	@Summary		List files
//	@Description	List files by user id
//	@Tags			file
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]model.File
//	@Router			/file/list [get]
func ListFileInfo(c *gin.Context) {
	pageInfo := model.PageInfo{}
	if err := c.ShouldBindJSON(&pageInfo); err != nil {
		bindRespWithStatus(c, http.StatusBadRequest, nil, err)
		return
	}
	ctx := utils.GetCtx(c)
	files, err := dal.QueryFileByUserId(ctx, &pageInfo)
	var total int64 = int64(len(files))
	pageInfo.Total = &total
	bindRespWithPageInfo(c, files, &pageInfo, err)
}

// ListFileInfoByIds godoc
//
//	@Summary		List files
//	@Description	List files by file ids
//	@Tags			file
//	@Accept			json
//	@Produce		json
//	@Param			ids body  req.ListFileByIdsReq true	"FILE ID LIST"
//	@Success		200	{array}		model.File
//	@Failure		400	{object}	vo.BaseVO
//	@Router			/file/listByIds [get]
func ListFileInfoByIds(c *gin.Context) {
	ctx := utils.GetCtx(c)
	ids := req.ListFileByIdsReq{}
	if err := c.ShouldBindJSON(&ids); err != nil {
		bindRespWithStatus(c, http.StatusBadRequest, nil, err)
		return
	}
	file, err := dal.QueryFileByIds(ctx, ids.Ids)
	bindResp(c, file, err)
}

func UpdateFile(c *gin.Context) {
	file := model.File{}
	if err := c.Bind(&file); err != nil {
		bindRespWithStatus(c, http.StatusBadRequest, nil, err)
		return
	}
}
