package controller

import (
	"errors"
	"fmt"
	"github.com/airren/echo-bio-backend/actuator"
	"github.com/airren/echo-bio-backend/dal"
	"github.com/airren/echo-bio-backend/model"
	"github.com/airren/echo-bio-backend/model/req"
	"github.com/airren/echo-bio-backend/model/vo"
	"github.com/airren/echo-bio-backend/service"
	"github.com/airren/echo-bio-backend/utils"
	"github.com/gin-gonic/gin"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path"
)

// UploadFile
// @Tags File
// @Summary  upload file
// @Description  upload file
// @Accept  json
// @Produce  json
// @Router /task/update [put]
// @Success 200 {string} Helooww
// @Param task body model.Job true "task"
func UploadFile(c *gin.Context) {
	fileItem := model.File{}

	//file, err := c.FormFile("file")
	form, _ := c.MultipartForm()
	files, ok := form.File["file"]
	if !ok {
		bindRespWithStatus(c, http.StatusBadRequest, nil, errors.New("file not be empty"))
		return
	}

	var file *multipart.FileHeader
	for _, t := range files {
		file = t
		log.Printf(file.Filename)

	}

	ctx := utils.GetCtx(c)

	folderForUser := path.Join(actuator.LocalDataBase, "")
	err := os.MkdirAll(folderForUser, 0777)
	if err != nil {
		return
	}

	// Upload the file to specific dst.
	fileItem.Id = utils.GenerateId()
	newFileName := fmt.Sprintf("%v-%v", fileItem.Id, file.Filename)
	filePath := path.Join(folderForUser, newFileName)
	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		bindRespWithStatus(c, http.StatusInternalServerError, nil, err)
		return
	}
	url := "http://www.echo-bio.cn/api/static/data/"
	fileItem.URLPath = fmt.Sprintf("%v%v", url, newFileName)

	f, err := service.UploadFile(ctx, fileItem)

	bindResp(c, vo.FileVO{Id: fmt.Sprint(f.Id)}, err)

}

func ListFile(c *gin.Context) {
	ctx := utils.GetCtx(c)
	files, err := dal.QueryFileByUserId(ctx)
	bindResp(c, files, err)
}

func CreateFile(c *gin.Context) {
	file := req.FileReq{}
	if err := c.Bind(&file); err != nil {
		bindRespWithStatus(c, http.StatusBadRequest, nil, err)
		return
	}
	ctx := utils.GetCtx(c)

	f, err := service.UpdateFile(ctx, file)
	bindResp(c, f, err)
}

func UpdateFile(c *gin.Context) {
	file := model.File{}
	if err := c.Bind(&file); err != nil {
		bindRespWithStatus(c, http.StatusBadRequest, nil, err)
		return
	}
}
