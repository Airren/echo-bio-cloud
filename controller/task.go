package controller

import (
	"github.com/airren/echo-bio-backend/dal"
	"github.com/airren/echo-bio-backend/model"
	"github.com/airren/echo-bio-backend/model/req"
	"github.com/airren/echo-bio-backend/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// GetOrderById
// @Tags Order
// @Summary get task by id
// @Description Get details of task by id
// @Accept  json
// @Produce  json
// @Router /task/{id} [get]
// @Success 200 {object} model.Order
// @Param id path uint true "task id"
func GetOrderById(c *gin.Context) {
	idStr := c.Param("id")
	ctx := util.GetOrgCtx(c)

	id, _ := strconv.Atoi(idStr)
	task, _ := dal.GetOrderById(ctx, int64(id))

	c.JSON(http.StatusOK, task)
}

// CreateOrder
// @Tags Order
// @Summary create task
// @Description create task
// @Accept  json
// @Produce  json
// @Router /task/create [post]
// @Success 200 {object} model.Order
// @Param task body model.Order true "task"
func CreateOrder(c *gin.Context) {
	var req model.Order
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	ctx := util.GetOrgCtx(c)
	req.CreatedAt = time.Now().Local()
	req.CreatedBy = "re"

	task, _ := dal.CreateOrder(ctx, &req)
	c.JSON(http.StatusOK, task)

}

// UpdateOrder
// @Tags Order
// @Summary update task
// @Description update task
// @Accept  json
// @Produce  json
// @Router /task/update [put]
// @Success 200 {object} model.Order
// @Param task body model.Order true "task"
func UpdateOrder(c *gin.Context) {
	var req model.Order
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	ctx := util.GetOrgCtx(c)
	req.UpdatedAt = time.Now().Local()
	req.UpdatedBy = "up"

	task, _ := dal.UpdateOrder(ctx, &req)
	c.JSON(http.StatusOK, task)

}


// QueryOrders
// @Tags Order
// @Summary query task
// @Description query task
// @Accept  json
// @Produce  json
// @Router /task/list [post]
// @Success 200 {array} model.Order
// @Param task body req.OrderReq true "task req"
func QueryOrders(c *gin.Context)  {

	var taskReq req.OrderReq
	if err := c.Bind(&taskReq); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	ctx := util.GetOrgCtx(c)

	task, _ := dal.QueryOrders(ctx, &taskReq)
	c.JSON(http.StatusOK, task)

}