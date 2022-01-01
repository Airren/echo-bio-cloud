package model

import (
	"time"
)

type RecordMeta struct {
	Id        int64      `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	AccountId int64      `json:"account_id" sql:"index"`
	Org       string     `json:"org" gorm:"type:varchar(20)"`
	CreatedAt time.Time  `json:"created_at" sql:"index;DEFAULT:'2000-01-01 00:00:00'"`
	UpdatedAt time.Time  `json:"updated_at" sql:"index;DEFAULT:'2000-01-01 00:00:00'"`
	DeletedAt *time.Time `json:"deleted_at" sql:"index"`
	UpdatedBy string     `json:"updated_by"`
	CreatedBy string     `json:"created_by"`
	DeletedBy string     `json:"deleted_by"`
}
