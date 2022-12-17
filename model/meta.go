package model

import (
	"gorm.io/gorm"
	"time"
)

type RecordMeta struct {
	Id        uint64         `json:"id" gorm:"primary_key;AUTO_INCREMENT;type:bigint(32)"`
	AccountId string         `json:"account_id" `
	Org       string         `json:"org" gorm:"type:varchar(20)"`
	CreatedAt time.Time      `json:"created_at" `
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" `
	UpdatedBy string         `json:"updated_by"`
	CreatedBy string         `json:"created_by"`
	DeletedBy string         `json:"deleted_by"`
}

const (
	DefaultPageSize = 10
	MaxPageSize     = 1000
)

type PageInfo struct {
	Page     int    `json:"page" form:"page" example:"1"`
	PageSize int    `json:"page_size" form:"page_size" example:"10"`
	Total    int64  `json:"total" form:"-"`
	OrderBy  string `json:"order_by" form:"-"`
	Asc      bool   `json:"asc" form:"-"`
}

func (pi *PageInfo) UpdatePageInfo() {
	// pagesize -1 means all
	if pi.PageSize == 0 || pi.PageSize < -2 {
		pi.PageSize = DefaultPageSize
	}
	if pi.PageSize > MaxPageSize {
		pi.PageSize = MaxPageSize
	}
	if pi.OrderBy == "" {
		pi.OrderBy = "created_at"
	}
	if pi.Page == 0 {
		pi.Page = 1
	}
}

func (pi *PageInfo) SetNoLimit() {
	pi.PageSize = -1
}
