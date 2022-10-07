package dal

import (
	"context"
	"github.com/airren/echo-bio-backend/model"
	"github.com/airren/echo-bio-backend/utils"
)

func CreateFile(ctx context.Context, file *model.File) (*model.File, error) {
	err := db.Create(file).Error
	return file, err
}
func UpdateFile(ctx context.Context, file *model.File) (*model.File, error) {
	err := db.Save(file).Error
	return file, err
}

func QueryFileByUserId(c context.Context) (file []*model.File, err error) {
	query := db.Model(&model.File{})
	userId, err := utils.GetUserId(c)
	if err != nil {
		return
	}
	query = queryByUserId(query, userId)
	err = query.Find(&file).Error
	return
}

//func UpdateFileWithAccountId(c context.Context, file *model.File) (*model.File, error) {
//	err := db.Model(&model.File{}).Where("id = ?", file.Id).
//		Update("account_id", file.AccountId).Error
//	return file, err
//}

func QueryFileById(ctx context.Context, id uint64) (file *model.File, err error) {
	query := db.Model(&model.File{})
	query = db.Where("id = ?", id)
	err = query.Find(&file).Error
	return
}
