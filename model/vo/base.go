package vo

import (
	"github.com/airren/echo-bio-backend/model"
	"gorm.io/gorm"
	"time"
)

type RecordMeta struct {
	AccountId string         `json:"account_id" `
	Org       string         `json:"org" gorm:"type:varchar(20)"`
	CreatedAt time.Time      `json:"created_at" `
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" `
	UpdatedBy string         `json:"updated_by"`
	CreatedBy string         `json:"created_by"`
	DeletedBy string         `json:"deleted_by"`
}

type VO interface {
	SetErrCode()
	SetErrMsg()
	SetPageInfo()
	SetData()
}

type BaseVO struct {
	ErrCode int    `json:"error_code"`
	ErrMsg  string `json:"error_message"`
	*model.PageInfo
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`
}

func (b *BaseVO) SetErrCode() {
}
func (b *BaseVO) SetErrMsg() {
}
func (b *BaseVO) SetPageInfo() {
}
func (b *BaseVO) SetData() {
}
