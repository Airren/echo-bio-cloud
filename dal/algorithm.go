package dal

import (
	"context"
	"fmt"
	"time"

	"github.com/airren/echo-bio-backend/model"
	"github.com/airren/echo-bio-backend/utils"
)

func CreateAlgorithm(ctx context.Context, task *model.Algorithm) (*model.Algorithm, error) {
	org := utils.GetOrgFromCtx(ctx)
	task.Org = org
	err := db.Create(task).Error
	return task, err
}

func UpdateAlgorithm(ctx context.Context, task *model.Algorithm) (*model.Algorithm, error) {
	org := utils.GetOrgFromCtx(ctx)
	task.Org = org
	err := db.Save(task).Error
	return task, err
}

func ListAlgorithms(ctx context.Context, algo *model.Algorithm, info *model.PageInfo) (
	algorithms []*model.Algorithm, err error) {
	query := db.Model(&model.Algorithm{})

	if algo.Label != "" {
		query = query.Where("label like ?", fmt.Sprintf("%%%s%%", algo.Label))
	}
	if info != nil {
		query.Count(&info.Total)
	}
	query = queryByPage(query, info)

	err = query.Find(&algorithms).Error

	return algorithms, err
}

func GetAlgorithmByName(ctx context.Context, name string) (algorithm *model.Algorithm, err error) {
	query := db.Where("name = ?", name)
	err = query.First(&algorithm).Error
	return algorithm, err
}

func GetAlgorithmById(ctx context.Context, id int64) (*model.Algorithm, error) {
	org := utils.GetOrgFromCtx(ctx)
	task := &model.Algorithm{}
	query := db.Where("id=? and org=?", id, org)
	err := query.Find(task).Error
	return task, err

}

func QueryAlgorithmByPriority(ctx context.Context, priority string) ([]*model.Algorithm, error) {
	org := utils.GetOrgFromCtx(ctx)
	tasks := make([]*model.Algorithm, 0)
	query := db.Where("priority = ?and org=?", priority, org)
	err := query.Find(&tasks).Error
	return tasks, err
}

func QueryAlgorithmByTime(ctx context.Context, startAt time.Time, endAt time.Time) ([]*model.Algorithm, error) {
	org := utils.GetOrgFromCtx(ctx)
	tasks := make([]*model.Algorithm, 0)
	query := db.Where("created_at >? and created_at <=? and org=?", startAt, endAt, org)
	err := query.Find(&tasks).Error
	return tasks, err
}

func CreateParameter(ctx context.Context, param *model.AlgoParameter) (*model.AlgoParameter, error) {
	err := db.Create(param).Error
	return param, err
}

func QueryParameter(ctx context.Context, param *model.AlgoParameter) (params []*model.AlgoParameter, err error) {
	query := db.Model(&model.AlgoParameter{})
	if param.AlgorithmId != 0 {
		query = query.Where("algorithm_id = ?", param.AlgorithmId)
	}
	err = query.Find(&params).Error
	return params, err
}
