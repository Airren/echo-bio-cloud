package dal

import (
	"gorm.io/gorm"

	"github.com/airren/echo-bio-backend/model"
)

func queryByPage(db *gorm.DB, info *model.PageInfo) *gorm.DB {
	if info != nil {
		return db.Offset(info.Page * info.PageSize).Limit(info.PageSize)
	}
	return db
}

func queryByUserId(db *gorm.DB, userId string) *gorm.DB {
	if userId != "" {
		return db.Where("account_id = ?", userId)
	}
	return db
}
