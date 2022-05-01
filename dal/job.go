package dal

import (
	"context"
	"github.com/airren/echo-bio-backend/model"
	"github.com/airren/echo-bio-backend/utils"
)

func CreateJob(ctx context.Context, job *model.Job) (*model.Job, error) {
	org := utils.GetOrgFromCtx(ctx)
	job.Org = org
	err := db.Create(job).Error
	return job, err
}

func QueryJobs(ctx context.Context, job *model.Job, info *model.PageInfo) (
	jobs []*model.Job, err error) {

	userId, err := utils.GetUserId(ctx)
	if err != nil {
		return
	}
	query := db.Model(&model.Job{})
	query = queryByPage(query, info)
	query = queryByUserId(query, userId)
	if job.Id != 0 {
		query = db.Where("id = ?", job.Id)
	}
	if job.Algorithm != "" {
		query = db.Where("algorithm = ?", job.Algorithm)
	}
	err = query.Find(&jobs).Error
	return
}
