package dal

import (
	"context"
	"fmt"

	"github.com/airren/echo-bio-backend/model"
)

func CreateJob(ctx context.Context, job *model.Job) (*model.Job, error) {
	err := db.Create(job).Error
	return job, err
}

func ListJob(ctx context.Context, job *model.Job, pageInfo *model.PageInfo) (
	jobs []*model.Job, err error) {
	query := db.Model(&model.Job{})
	query = queryByUserId(ctx, query)
	if job.Id != 0 {
		query = query.Where("id = ?", job.Id)
	}
	if job.Algorithm != "" {
		query = query.Where("algorithm  like ?", fmt.Sprintf("%%%s%%", job.Algorithm))
	}
	if job.Name != "" {
		query = query.Where("name  like ?", fmt.Sprintf("%%%s%%", job.Name))
	}
	if job.Status != "" {
		query = query.Where("status = ?", job.Status)
	}
	if pageInfo != nil {
		query.Count(&pageInfo.Total)
	}
	query = queryByPage(query, pageInfo)
	err = query.Find(&jobs).Error
	return
}
