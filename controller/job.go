package controller

import (
	"fmt"
	"github.com/airren/echo-bio-backend/service"
	"net/http"
	"os"
	"path"

	"github.com/gin-gonic/gin"

	"github.com/airren/echo-bio-backend/actuator"
	"github.com/airren/echo-bio-backend/model/req"
	"github.com/airren/echo-bio-backend/utils"
)

// CreateJob
// @Tags Job
// @Summary update task
// @Description update task
// @Accept  json
// @Produce  json
// @Router /task/update [put]
// @Success 200 string
// @Param task body model.Job true "task"
func CreateJob(c *gin.Context) {
	var jobReq req.JobReq

	algo := c.Query("algorithm")

	file, _ := c.FormFile("file")

	ctx := utils.GetCtx(c)
	// Create folder for user
	userId, err := utils.GetUserId(ctx)
	if err != nil {
		return
	}

	folderForUser := path.Join(actuator.LocalDataBase, userId)
	err = os.MkdirAll(folderForUser, 0777)
	if err != nil {
		return
	}

	// Upload the file to specific dst.
	jobReq.Id = utils.GenerateId()
	newFileName := fmt.Sprintf("%v-%v", jobReq.Id, file.Filename)
	err = c.SaveUploadedFile(file, path.Join(folderForUser, newFileName))
	if err != nil {
		return
	}

	jobReq.Algorithm = algo
	jobReq.InputFile = newFileName
	err = service.CreateJob(ctx, jobReq)
	bindResp(c, nil, err)

}

func QueryJob(c *gin.Context) {
	var jobReq req.JobReq
	if err := c.Bind(&jobReq); err != nil {
		bindRespWithStatus(c, http.StatusBadRequest, nil, err)
		return
	}
	ctx := utils.GetCtx(c)
	jobs, err := service.QueryJob(ctx, jobReq)
	bindRespWithPageInfo(c, jobs, &jobReq.PageInfo, err)
}
