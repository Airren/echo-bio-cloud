package dal

import (
	"context"
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

func QueryAlgorithms(ctx context.Context, algo *model.Algorithm, info *model.PageInfo) (
	algorithms []*model.Algorithm, err error) {
	query := db.Model(&model.Algorithm{})
	query = queryByPage(query, info)
	if algo.Label != "" {
		query = db.Where("label = ?", algo.Label)
	}
	err = query.Find(&algorithms).Error
	//if info != nil {
	//	query.Count(info.Total)
	//}

	return algorithms, err
}

func QueryAlgorithmsByName(ctx context.Context, name string) ([]*model.Algorithm, error) {
	var algorithms []*model.Algorithm
	query := db.Where("name = ?", name)
	err := query.Find(&algorithms).Error
	return algorithms, err
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

func CreateParameter(ctx context.Context, param *model.Parameter) (*model.Parameter, error) {
	err := db.Create(param).Error
	return param, err
}

func QueryParameter(ctx context.Context, param *model.Parameter) (params []*model.Parameter, err error) {
	query := db.Model(&model.Parameter{})
	if param.AlgorithmId != 0 {
		query = query.Where("algorithm_id = ?", param.AlgorithmId)
	}
	err = query.Find(&params).Error
	return params, err
}
