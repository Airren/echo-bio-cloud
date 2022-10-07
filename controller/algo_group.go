package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/airren/echo-bio-backend/model/req"
	"github.com/airren/echo-bio-backend/service"
	"github.com/airren/echo-bio-backend/utils"
)

func CreateAlgoGroup(c *gin.Context) {
	var groupReq req.Group
	if err := c.Bind(&groupReq); err != nil {
		bindRespWithStatus(c, http.StatusBadRequest, nil, err)
		return
	}
	ctx := utils.GetCtx(c)
	err := service.CreateAlgoGroup(ctx, groupReq)
	bindResp(c, nil, err)
}

func DeleteAlgoGroupById(c *gin.Context) {
	var idStrs []string
	var ids []uint64
	if err := c.Bind(&idStrs); err != nil {
		bindRespWithStatus(c, http.StatusBadRequest, nil, err)
		return
	}
	for _, idStr := range idStrs {
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err == nil {
			ids = append(ids, uint64(id))
		}
	}
	err := service.DeleteAlgoGroupById(c, ids)
	bindResp(c, nil, err)
}

func UpdateAlgoGroup(c *gin.Context) {
	var groupReq req.Group
	if err := c.Bind(&groupReq); err != nil {
		bindRespWithStatus(c, http.StatusBadRequest, nil, err)
		return
	}
	ctx := utils.GetCtx(c)
	err := service.UpdateAlgoGroup(ctx, groupReq)
	bindResp(c, nil, err)
}

func ListAlgoGroup(c *gin.Context) {
	ctx := utils.GetCtx(c)
	gs, err := service.ListAlgoGroup(ctx)
	bindResp(c, gs, err)
}
