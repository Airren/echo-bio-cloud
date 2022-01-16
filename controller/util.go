package controller

import (
	"context"
	"github.com/airren/echo-bio-backend/model"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

func bindResp(c *gin.Context, data interface{}, err error) {
	bindRespWithStatus(c, 200, data, err)
}

func bindRespWithStatus(c *gin.Context, code int, data interface{}, err error) {
	bindRespWithStatusAndPageInfo(c, code, data, nil, err)
}

func bindRespWithPageInfo(c *gin.Context, data interface{}, pageInfo *model.PageInfo, err error) {
	bindRespWithStatusAndPageInfo(c, 200, data, pageInfo, err)
}

func bindRespWithStatusAndPageInfo(c *gin.Context, code int, data interface{}, pageInfo *model.PageInfo, err error) {
	//errCode := 0
	//errMsg := ""
	funcName := ""
	names := strings.Split(c.HandlerName(), "/")
	if len(names) != 0 {
		funcName = names[len(names)-1]
	}

	tagKv := map[string]string{
		"method":      c.Request.Method,
		"func_name":   funcName,
		"is_error":    "0",
		"status_code": strconv.Itoa(code),
	}

	if err != nil {
		//errCode = -1
		//errMsg = err.Error()
		//logs.CtxError(c, "org: %v path: %v method: %v, err: %v", getOrgByGinCtx(c), c.Request.URL.Path, c.Request.Method, err)
		tagKv["is_error"] = "1"
	}

	//helpers.ApiThroughput.Inc(getCtxByGinCtx(c), tagKv)

	//data.ErrCode = errCode
	//data.PageInfo = vo.NewRespPageInfo(pageInfo)
	//data.ErrMsg = errMsg
	c.JSON(code, data)
}

//func bindResp(c *gin.Context, code int, resp *vo.BaseVO) {
//funcName := ""
//names := strings.Split(c.HandlerName(), "/")
//if len(names) != 0 {
//	funcName = names[len(names)-1]
//}

//tagKv := map[string]string{
//	"method":      c.Request.Method,
//	"func_name":   funcName,
//	"is_error":    "0",
//	"status_code": strconv.Itoa(code),
//}

//if proxyResp.ErrMsg != "" {
//	logs.CtxError(c, "org: %v path: %v method: %v, err: %v", getOrgByGinCtx(c), c.Request.URL.Path, c.Request.Method, proxyResp.ErrMsg)
//	tagKv["is_error"] = "1"
//}
//
//helpers.ApiThroughput.Inc(getCtxByGinCtx(c), tagKv)

//	c.JSON(code, proxyResp)
//}

func getOrgByGinCtx(c *gin.Context) string {
	org := c.Request.Header.Get("org")
	if org == "" {
		org = c.Request.Header.Get("X-Org")
		if org == "" {
			org = c.Query("org")
		}
	}

	return org
}

func getCtxByGinCtx(c *gin.Context) context.Context {
	ctx := context.Background()
	//ctx = models.WithValueOrg(ctx, getOrgByGinCtx(c))
	//
	//if logId, ok := c.Get(utils.LogIdKey); ok {
	//	ctx = context.WithValue(ctx, utils.LogIdKey, logId)
	//}

	return ctx
}

func getRemoteAddr(c *gin.Context) string {
	remoteAddr := strings.TrimSpace(c.Request.Header.Get("X-Real-IP"))
	if remoteAddr != "" {
		return remoteAddr
	}
	remoteAddr = strings.TrimSpace(c.Request.Header.Get("X-Forwarded-For"))
	if remoteAddr != "" {
		lst := strings.Split(remoteAddr, ",")
		if len(lst) != 0 {
			return strings.TrimSpace(lst[len(lst)-1])
		}
	}
	return ""
}

func isValidOrg(c *gin.Context) error {
	org := c.Request.Header.Get("org")
	if org == "" {
		c.Request.Header.Set("org", "ms")
	}
	return nil
}
