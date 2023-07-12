package dal

import (
	"context"
	"fmt"

	"github.com/airren/echo-bio-backend/model"
)

func CreateJob(ctx context.Context, job *model.AnalysisJob) (*model.AnalysisJob, error) {
	err := db.Create(job).Error
	return job, err
}

func UpdateJobStatus(ctx context.Context, jobId uint64, status model.JobStatus) error {
	_ = db.Model(&model.AnalysisJob{}).Where("id = ?", jobId).Update("status", status)
	return nil
}
func UpdateJobResult(ctx context.Context, jobId uint64, result uint64) error {
	_ = db.Model(&model.AnalysisJob{}).Where("id = ?", jobId).Update("result", result)
	return nil
}

func ListJob(ctx context.Context, job *model.AnalysisJob, pageInfo *model.PageInfo) (
	jobs []*model.AnalysisJob, err error) {
	query := db.Model(&model.AnalysisJob{})
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
