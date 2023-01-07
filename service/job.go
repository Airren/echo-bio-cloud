package service

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strconv"
	"time"

	"github.com/airren/echo-bio-backend/dal"
	"github.com/airren/echo-bio-backend/model"
	"github.com/airren/echo-bio-backend/model/req"
	"github.com/airren/echo-bio-backend/model/vo"
	"github.com/airren/echo-bio-backend/utils"
)

func CreateJob(c context.Context, req req.JobReq) (err error) {
	job := JobToEntity(req)
	userId, err := utils.GetUserId(c)
	if err != nil {
		return err
	}
	org := utils.GetOrgFromCtx(c)

	job.RecordMeta = model.RecordMeta{
		Id:        utils.GenerateId(),
		AccountId: userId,
		Org:       org,
		CreatedAt: time.Now(),
		CreatedBy: userId,
	}
	err = CheckJobIsValid(c, job)
	if err != nil {
		return err
	}
	// TODO put this job to the work queue

	err = job.TransferParameter()
	if err != nil {
		return err
	}
	_, err = dal.CreateJob(c, job)
	return err
}

func CheckJobIsValid(ctx context.Context, job *model.Job) error {
	algo, err := dal.QueryAlgorithmsByName(ctx, job.Algorithm)
	if err != nil {
		return err
	} else if err == gorm.ErrRecordNotFound {
		return errors.New("please use an existed algorithm")
	}

	parametes, err := dal.QueryParameter(ctx, &model.AlgoParameter{AlgorithmId: algo.Id})
	if err != nil {
		return err
	}
	for _, p := range parametes {
		if _, ok := job.Parameters[p.Name]; !ok {
			if p.Name == "name" {
				continue
			}
			return errors.New(fmt.Sprintf("parametes %s is not provided", p.Name))
		}
	}
	return nil
}

func ListJob(c context.Context, req req.JobReq) (jobVO []*vo.JobVO, pageInfo *model.PageInfo, err error) {
	job := JobToEntity(req)
	jobs, err := dal.ListJob(c, job, &req.PageInfo)
	if err != nil {
		return
	}
	for _, j := range jobs {
		jobVO = append(jobVO, JobToVO(*j))
	}
	return jobVO, &req.PageInfo, err
}

func JobToEntity(req req.JobReq) *model.Job {
	var Id int64
	if req.Id != "" {
		Id, _ = strconv.ParseInt(req.Id, 10, 64)
	}

	return &model.Job{
		RecordMeta:  model.RecordMeta{Id: uint64(Id)},
		Name:        req.Name,
		Algorithm:   req.Algorithm,
		Parameters:  req.Parameters,
		Status:      req.Status,
		Description: req.Description,
	}
}

func JobToVO(job model.Job) *vo.JobVO {
	return &vo.JobVO{
		RecordMeta:  RecordMetaToVO(job.RecordMeta),
		Name:        job.Name,
		Algorithm:   job.Algorithm,
		Parameters:  job.Parameters,
		Outputs:     job.Outputs,
		Status:      job.Status,
		Description: job.Description,
	}
}
