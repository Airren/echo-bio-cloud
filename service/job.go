package service

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"gorm.io/gorm"

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
	job.Status = model.PENDING
	err = CheckJobIsValid(c, job)
	if err != nil {
		return err
	}
	err = job.TransferParametersToJson()
	if err != nil {
		return err
	}
	_, err = dal.CreateJob(c, job)
	PublishJob(c, job)
	return err
}

func CheckJobIsValid(ctx context.Context, job *model.AnalysisJob) error {
	algo, err := dal.GetAlgorithmByName(ctx, job.Algorithm)
	if err != nil {
		return err
	} else if err == gorm.ErrRecordNotFound {
		return errors.New("please use an existed algorithm")
	}

	parameters, err := dal.QueryParameter(ctx, &model.AlgoParameter{AlgorithmId: algo.Id})
	if err != nil {
		return err
	}
	for _, p := range parameters {
		if _, ok := job.Parameters[p.Name]; p.Required && !ok {
			if p.Name == "name" {
				continue
			}
			return errors.New(fmt.Sprintf("for job: %v parameters: %s is not provided", job.Name, p.Name))
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

func JobToEntity(req req.JobReq) *model.AnalysisJob {
	var Id int64
	if req.Id != "" {
		Id, _ = strconv.ParseInt(req.Id, 10, 64)
	}

	return &model.AnalysisJob{
		RecordMeta:  model.RecordMeta{Id: uint64(Id)},
		Name:        req.Name,
		Algorithm:   req.Algorithm,
		Parameters:  req.Parameters,
		Status:      req.Status,
		Description: req.Description,
	}
}

func JobToVO(job model.AnalysisJob) *vo.JobVO {
	return &vo.JobVO{
		RecordMeta:  RecordMetaToVO(job.RecordMeta),
		Name:        job.Name,
		Algorithm:   job.Algorithm,
		Parameters:  job.Parameters,
		Outputs:     job.Outputs,
		Status:      job.Status,
		Result:      fmt.Sprintf("/api/v1/file/download/%v", job.Result),
		Description: job.Description,
	}
}

func GetImageForAnalysisJob(j *model.AnalysisJob) (string, error) {
	algo, err := dal.GetAlgorithmByName(context.TODO(), j.Algorithm)
	return algo.DockerImage, err
}
