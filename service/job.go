package service

import (
	"context"
	"time"

	"github.com/airren/echo-bio-backend/actuator"
	"github.com/airren/echo-bio-backend/dal"
	"github.com/airren/echo-bio-backend/model"
	"github.com/airren/echo-bio-backend/model/req"
	"github.com/airren/echo-bio-backend/model/vo"
	"github.com/airren/echo-bio-backend/utils"
)

func CreateJob(c context.Context, req req.JobReq) (err error) {
	userId, err := utils.GetUserId(c)
	if err != nil {
		return err
	}
	job := model.Job{
		RecordMeta: model.RecordMeta{
			Id:        req.Id,
			AccountId: userId,
			Org:       "",
			CreatedAt: time.Now(),
			CreatedBy: userId,
		},
		Algorithm: req.Algorithm,
		InputFile: req.InputFile,
		Parameter: req.Parameter,
	}
	outfile, err := actuator.StartPaint(job)
	if err != nil {
		return err
	}

	job.OutPutFile = outfile
	_, err = dal.CreateJob(c, &job)
	return err
}

func QueryJob(c context.Context, req req.JobReq) (jobVO []*vo.JobVO, err error) {
	job := JobToEntity(req)
	jobs, err := dal.QueryJobs(c, job, &req.PageInfo)
	if err != nil {
		return
	}
	for _, j := range jobs {
		jobVO = append(jobVO, JobToVO(*j))
	}
	return
}

func JobToEntity(req req.JobReq) *model.Job {
	return &model.Job{
		RecordMeta: model.RecordMeta{Id: req.Id},
		Algorithm:  req.Algorithm,
		InputFile:  req.InputFile,
		OutPutFile: req.OutPutFile,
		Parameter:  req.Parameter,
	}
}

func JobToVO(job model.Job) *vo.JobVO {
	return &vo.JobVO{
		Id:         job.Id,
		Algorithm:  job.Algorithm,
		InputFile:  job.InputFile,
		OutPutFile: job.OutPutFile,
		Parameter:  job.Parameter,
	}
}
