package controller

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/airren/echo-bio-backend/minio"
	"github.com/airren/echo-bio-backend/model/req"
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
//	@Router			/file/upload/ [post]
func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		bindRespWithStatus(c, http.StatusBadRequest, nil, err)
		return
	}
	isPublic, _ := c.GetPostForm("visibility")
	visibility := 2
	if isPublic == "public" {
		visibility = 1
	}

	ctx := utils.GetCtx(c)
	fileInfo, err := service.UploadFile(ctx, file, visibility)
	if err != nil {
		bindRespWithStatus(c, http.StatusInternalServerError, nil, err)
		return
	}
	bindResp(c, fileInfo, err)
}

// UpdateFileInfo godoc
//
//	@Summary		update file info
//	@Description	update file info
//	@Tags			file
//	@Accept			json
//	@Produce		json
//	@Param			file	formData	file	true	"FILE"
//	@Success		200	{object}	vo.FileVO
//	@Router			/file/update/ [put]
func UpdateFileInfo(c *gin.Context) {
	fileReq := req.FileReq{}

	if err := c.Bind(&fileReq); err != nil {
		bindRespWithStatus(c, http.StatusBadRequest, nil, err)
		return
	}
	ctx := utils.GetCtx(c)
	fileInfo, err := service.UpdateFileInfo(ctx, &fileReq)
	bindResp(c, fileInfo, err)
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
	fileInfo, bytes, err := service.DownloadFileById(ctx, fileId)
	if err != nil {
		bindRespWithStatus(c, http.StatusInternalServerError, nil, err)
		return
	}

	contentType := minio.GetContentType(fileInfo.FileType)


	fileName := url.QueryEscape(fileInfo.Name)
	c.Header("Content-Disposition",
		fmt.Sprintf("inline; filename*=UTF-8''%s.%s", fileName, fileInfo.FileType))
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Cache-Control", "no-cache")
	c.Data(http.StatusOK, contentType, bytes)
}

// ListFileInfo godoc
//
//	@Summary	    List files
//	@Description	List files by user id
//	@Tags			file
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]model.File
//	@Router			/file/list [post]
func ListFileInfo(c *gin.Context) {
	fileReq := req.FileReq{}
	if err := c.Bind(&fileReq); err != nil {
		bindRespWithStatus(c, http.StatusBadRequest, nil, err)
		return
	}
	ctx := utils.GetCtx(c)
	fileReq.UpdatePageInfo()
	files, pageInfo, err := service.ListFileInfo(ctx, &fileReq)
	bindRespWithPageInfo(c, files, pageInfo, err)
}

// ListFileInfoByIds godoc
//
//	@Summary		List files
//	@Description	List files by file ids
//	@Tags			file
//	@Accept			json
//	@Produce		json
//	@Param			ids body  req.IdsReq true	"FILE ID LIST"
//	@Success		200	{array}		model.File
//	@Failure		400	{object}	vo.BaseVO
//	@Router			/file/listByIds [get]
func ListFileInfoByIds(c *gin.Context) {
	ctx := utils.GetCtx(c)
	ids := req.IdsReq{}
	if err := c.ShouldBindJSON(&ids); err != nil {
		bindRespWithStatus(c, http.StatusBadRequest, nil, err)
		return
	}
	file, err := service.QueryFileByIds(ctx, &ids)
	bindResp(c, file, err)
}

// DeleteFileInfoByIds godoc
//
//	@Summary		Delete files
//	@Description	Delete  files by file ids
//	@Tags			file
//	@Accept			json
//	@Produce		json
//	@Param			ids body  req.IdsReq true	"FILE ID LIST"
//	@Success		200	{array}		vo.BaseVO
//	@Failure		400	{object}	vo.BaseVO
//	@Router			/file/delete_by_ids [get]
func DeleteFileInfoByIds(c *gin.Context) {
	ctx := utils.GetCtx(c)
	ids := req.IdsReq{}
	if err := c.ShouldBindJSON(&ids); err != nil {
		bindRespWithStatus(c, http.StatusBadRequest, nil, err)
		return
	}
	err := service.DeleteFileByIds(ctx, &ids)
	bindResp(c, nil, err)
}
