package dal

import (
	"context"
	"fmt"
	"gorm.io/gorm"

	"github.com/airren/echo-bio-backend/model"
	"github.com/airren/echo-bio-backend/utils"
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
	if info.OrderBy != "" && !info.Asc {
		db = db.Order(fmt.Sprintf("%v DESC", info.OrderBy))
	} else if info.OrderBy != "" {
		db = db.Order(fmt.Sprintf("%v ASC", info.OrderBy))
	}
	return db
}

func queryByUserId(ctx context.Context, db *gorm.DB) *gorm.DB {
	userId, _ := utils.GetUserId(ctx)
	if userId != "" {
		return db.Where("account_id = ?", userId)
	}
	return db
}
