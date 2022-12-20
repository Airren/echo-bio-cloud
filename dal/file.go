package dal

import (
	"context"
	"errors"
	"fmt"

	"github.com/airren/echo-bio-backend/model"
	"github.com/airren/echo-bio-backend/utils"
)

func InsertFileInfo(ctx context.Context, file *model.File) (*model.File, error) {
	err := db.Save(file).Error
	return file, err
}
func UpdateFileInfo(ctx context.Context, file *model.File) (*model.File, error) {
	err := db.Model(file).Omit("id").Updates(file).Error
	return file, err
}

func ListFiles(c context.Context, file *model.File, pageInfo *model.PageInfo) (
	files []*model.File, err error) {
	query := db.Model(&model.File{})
	query = queryByUserId(c, query)
	if file.Name != "" {
		query = query.Where("name like ?", fmt.Sprintf("%%%s%%", file.Name))
	}
	if pageInfo != nil {
		query.Count(&pageInfo.Total)
	}

	query = queryByPage(query, pageInfo)
	err = query.Find(&files).Error
	return
}

//func UpdateFileWithAccountId(c context.Context, file *model.File) (*model.File, error) {
//	err := db.Model(&model.File{}).Where("id = ?", file.Id).
//		Update("account_id", file.AccountId).Error
//	return file, err
//}

func QueryFileById(ctx context.Context, id uint64) (file *model.File, err error) {
	query := db.Model(&model.File{})
	query = query.Where("id = ?", id)
	err = query.Find(&file).Error
	if file.IsPublic {
		return
	}
	userId, _ := utils.GetUserId(ctx)
	if userId != file.AccountId {
		return nil, errors.New("invalid user")
	}
	return
}

func QueryFileByIds(ctx context.Context, id []int64) (file []*model.File, err error) {
	query := db.Model(&model.File{})
	query = query.Where("id IN ?", id)
	query = queryByUserId(ctx, query)
	err = query.Find(&file).Error
	return
}

func CheckIfFileExist(ctx context.Context, Md5 string) (exist bool, err error) {
	var total int64
	err = db.Model(&model.File{}).
		Where("MD5 = ? AND org = ?", Md5, utils.GetOrgFromCtx(ctx)).
		Count(&total).Error
	return total > 0, err
}

func QueryFileByMd5(ctx context.Context, Md5 string) (file *model.File, err error) {
	err = db.Model(&model.File{}).Where("MD5 = ? AND org = ?", Md5, utils.GetOrgFromCtx(ctx)).First(&file).Error
	return file, err
}
