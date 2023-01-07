package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/airren/echo-bio-backend/model/req"
	"github.com/airren/echo-bio-backend/service"
	"github.com/airren/echo-bio-backend/utils"
)

// CreateJob
// @Tags Job
// @Summary update task
// @Description update task
// @Accept  json
// @Produce  json
// @Router /task/update [put]
// @Success 200 {string} vo.BaseVO
// @Param task body model.Job true "task"
func CreateJob(c *gin.Context) {
	var jobReq req.JobReq

	if err := c.Bind(&jobReq); err != nil {
		bindRespWithStatus(c, http.StatusBadRequest, nil, err)
		return
	}
	ctx := utils.GetCtx(c)
	err := service.CreateJob(ctx, jobReq)
	bindResp(c, nil, err)

}

func ListJob(c *gin.Context) {
	var jobReq req.JobReq
	if err := c.Bind(&jobReq); err != nil {
		bindRespWithStatus(c, http.StatusBadRequest, nil, err)
		return
	}
	ctx := utils.GetCtx(c)
	jobs, pageInfo, err := service.ListJob(ctx, jobReq)
	bindRespWithPageInfo(c, jobs, pageInfo, err)
}
