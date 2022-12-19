package dal

import (
	"context"

	"github.com/airren/echo-bio-backend/model"
	"github.com/airren/echo-bio-backend/utils"
)

func InsertFileMetaInfo(ctx context.Context, file *model.File) (*model.File, error) {
	err := db.Save(file).Error
	return file, err
}

func QueryFileByUserId(c context.Context, pageInfo *model.PageInfo) (file []*model.File, err error) {
	query := db.Model(&[]model.File{})
	userId, err := utils.GetUserId(c)
	if err != nil {
		return
	}
	query = queryByUserId(query, userId)
	query = queryByPage(query, pageInfo)
	err = query.Find(&file).Error
	return
}

//func UpdateFileWithAccountId(c context.Context, file *model.File) (*model.File, error) {
//	err := db.Model(&model.File{}).Where("id = ?", file.Id).
//		Update("account_id", file.AccountId).Error
//	return file, err
//}

func QueryFileById(ctx context.Context, id uint64) (file *model.File, err error) {
	query := db.Model(&[]model.File{})
	query = db.Where("id = ?", id)
	err = query.Find(&file).Error
	return
}

func QueryFileByIds(ctx context.Context, id []uint64) (file []*model.File, err error) {
	query := db.Model(&[]model.File{})
	query = db.Where("id IN ?", id)
	err = query.Find(&file).Error
	return
}

func CheckIfFileExist(ctx context.Context, Md5 string) (total int64, err error) {
	_ = db.Model(&model.File{}).Where("MD5 = ?", Md5).Count(&total)
	return
}
