package vo

import "github.com/airren/echo-bio-backend/model"

type VO interface {
	SetErrCode()
	SetErrMsg()
	SetPageInfo()
	SetData()
}

type BaseVO struct {
	ErrCode  int             `json:"error_code"`
	ErrMsg   string          `json:"error_message"`
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	Data     interface{}     `json:"data"`
}

func (b *BaseVO) SetErrCode() {
}
func (b *BaseVO) SetErrMsg() {
}
func (b *BaseVO) SetPageInfo() {
}
func (b *BaseVO) SetData() {
}
