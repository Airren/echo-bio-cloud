package dal

import (
	"fmt"
	"gorm.io/gorm"

	"github.com/airren/echo-bio-backend/model"
)

func queryByPage(db *gorm.DB, info *model.PageInfo) *gorm.DB {
	if info != nil {
		db = db.Offset((info.Page - 1) * info.PageSize).Limit(info.PageSize)
	}
	if info.OrderBy != "" && !info.Asc {
		db = db.Order(fmt.Sprintf("%v DESC", info.OrderBy))
	} else if info.OrderBy != "" {
		db = db.Order(fmt.Sprintf("%v ASC", info.OrderBy))
	}
	return db
}

func queryByUserId(db *gorm.DB, userId string) *gorm.DB {
	if userId != "" {
		return db.Where("account_id = ?", userId)
	}
	return db
}
