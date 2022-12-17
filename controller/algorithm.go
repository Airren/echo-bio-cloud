package controller

import (
	"github.com/airren/echo-bio-backend/dal"
	"github.com/airren/echo-bio-backend/model/req"
	"github.com/airren/echo-bio-backend/service"
	"github.com/airren/echo-bio-backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetAlgorithmById
// @Tags Algorithm
// @Summary get task by id
// @Description Get details of task by id
// @Accept  json
// @Produce  json
// @Router /task/{id} [get]
// @Success 200 {object} model.Algorithm
// @Param id path uint true "task id"
func GetAlgorithmById(c *gin.Context) {
	idStr := c.Param("id")
	ctx := utils.GetCtx(c)

	id, _ := strconv.Atoi(idStr)
	task, _ := dal.GetAlgorithmById(ctx, int64(id))

	c.JSON(http.StatusOK, task)
}

// CreateAlgorithmByFile
// @Tags Algorithm
// @Summary create task
// @Description create task
// @Accept  json
// @Produce  json
// @Router /task/create_by_file [post]
// @Success 200 {object} model.Algorithm
// @Param task body model.Algorithm true "task"
func CreateAlgorithmByFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		bindRespWithStatus(c, http.StatusBadRequest, nil, err)
		return
	}
	f, err := file.Open()
	if err != nil {
		bindRespWithStatus(c, http.StatusBadRequest, nil, err)
		return
	}
	ctx := utils.GetCtx(c)
	err = service.CreateAlgorithmByFile(ctx, f)
	bindResp(c, nil, err)
}

func CreateAlgorithm(c *gin.Context) {
	var algoReq req.AlgorithmReq
	if err := c.Bind(&algoReq); err != nil {
		bindRespWithStatus(c, http.StatusBadRequest, nil, err)
		return
	}
	ctx := utils.GetCtx(c)
	err := service.CreateAlgorithm(ctx, algoReq)
	bindResp(c, nil, err)
}

// UpdateAlgorithm
// @Tags Algorithm
// @Summary update task
// @Description update task
// @Accept  json
// @Produce  json
// @Router /task/update [put]
// @Success 200 {object} model.Algorithm
// @Param task body model.Algorithm true "task"
func UpdateAlgorithm(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		bindRespWithStatus(c, http.StatusBadRequest, nil, err)
		return
	}
	f, err := file.Open()
	if err != nil {
		bindRespWithStatus(c, http.StatusBadRequest, nil, err)
		return
	}
	ctx := utils.GetCtx(c)
	err = service.UpdateAlgorithm(ctx, f)
	bindResp(c, nil, err)

}

// QueryAlgorithm
// @Tags Algorithm
// @Summary query task
// @Description query task
// @Accept  json
// @Produce  json
// @Router /task/list [post]
// @Success 200 {array} model.Algorithm
// @Param task body req.AlgorithmReq true "task req"
func QueryAlgorithm(c *gin.Context) {

	var algoReq req.AlgorithmReq
	if err := c.Bind(&algoReq); err != nil {
		bindRespWithStatus(c, http.StatusBadRequest, nil, err)
		return
	}
	ctx := utils.GetCtx(c)

	algorithms, pageInfo, err := service.QueryAlgorithm(ctx, algoReq)
	bindRespWithPageInfo(c, algorithms, pageInfo, err)

}
