package service

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/airren/echo-bio-backend/mq"
	"strconv"
	"text/template"
	"time"

	log "github.com/sirupsen/logrus"
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
	err = CheckJobIsValid(c, job)
	if err != nil {
		return err
	}
	// TODO put this job to the work queue

	err = job.TransferParametersToJson()
	if err != nil {
		return err
	}
	_, err = dal.CreateJob(c, job)
	mq.PublishJob(c, job)
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
		if _, ok := job.Parameters[p.Name]; !ok {
			if p.Name == "name" {
				continue
			}
			return errors.New(fmt.Sprintf("parameters %s is not provided", p.Name))
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
		Description: job.Description,
	}
}

func GetImageForAnalysisJob(j *model.AnalysisJob) (string, error) {
	algo, err := dal.GetAlgorithmByName(context.TODO(), j.Algorithm)
	return algo.DockerImage, err
}

func GetCommandForAnalysisJob(j *model.AnalysisJob) (command string, err error) {

	if len(j.Parameters) == 0 {
		err := j.TransferJsonToParameters()
		if err != nil {
			return "", err
		}
	}
	//1. download the file to the AnalysisJob container
	ctx := context.TODO()
	// wget with
	//shell := "mkdir -p /tmp/job && wget www.echo-bio.cn:8088/api/v1/file/public/download/${file_id} "
	//2. create analysis command

	// get algorithm detail
	algo, err := dal.GetAlgorithmByName(context.TODO(), j.Algorithm)
	if err != nil {
		log.Infof("get algo failed", j)
		return
	}
	// get all the need file, and download to the docker image
	parameters, _ := dal.QueryParameter(ctx, &model.AlgoParameter{
		AlgorithmId: algo.Id,
	})
	fileParameterMap := make(map[string]string)
	for _, f := range parameters {
		if f.Type == model.ParamFile {
			fileParameterMap[f.Name] = j.Parameters[f.Name].(string)
		}
	}

	tmpl, err := template.New("test").Parse(algo.Command)
	if err != nil {
		log.Fatal(err)
	}

	commandstr := bytes.Buffer{}
	err = tmpl.Execute(&commandstr, fileParameterMap)
	if err != nil {
		log.Fatal(err)
	}
	// use go template render the command

	return commandstr.String(), err

}
