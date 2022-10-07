package dal

import (
	"context"
	"github.com/airren/echo-bio-backend/model"
)

func CreateAlgoGroup(ctx context.Context, group *model.AlgoGroup) (*model.AlgoGroup, error) {
	err := db.Create(group).Error
	return group, err
}

func UpdateAlgoGroup(ctx context.Context, group *model.AlgoGroup) (*model.AlgoGroup, error) {
	err := db.Save(group).Error
	return group, err
}

func DeleteAlgoGroupById(ctx context.Context, Ids []uint64) error {
	return db.Delete(&model.AlgoGroup{}, Ids).Error
}

func ListAlgoGroup(ctx context.Context) (groups []*model.AlgoGroup, err error) {
	query := db.Model(&model.AlgoGroup{})
	err = query.Find(&groups).Error
	return groups, err
}

func QueryAlgoGroupByName(ctx context.Context, name string) (groups []*model.AlgoGroup, err error) {
	query := db.Model(&model.AlgoGroup{})
	query = query.Where("name = ?", name)
	err = query.Find(&groups).Error
	return groups, err
}
