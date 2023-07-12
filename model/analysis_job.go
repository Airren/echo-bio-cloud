package model

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

type JobStatus string

const (
	PENDING     JobStatus = "Pending"
	PROGRESSING JobStatus = "Progressing"
	COMPLETED   JobStatus = "Completed"
	FAILED      JobStatus = "Failed"
	CANCELED    JobStatus = "Canceled"
)

type AnalysisJob struct {
	RecordMeta
	Name          string                 `gorm:"type:varchar(128); not null"`
	Algorithm     string                 `gorm:"type:varchar(64); not null"`
	Parameters    map[string]interface{} `gorm:"-"`
	ParametersStr string                 `gorm:"type:varchar(256); not null"`
	Outputs       string                 `gorm:"type:varchar(256); not null"`
	Result        uint64                 `gorm:"type:bigint(32)"`
	Status        JobStatus              `gorm:"type:varchar(32); not null"`
	Description   string                 `gorm:"type:text"`
}

func (j *AnalysisJob) TransferParametersToJson() error {
	b, err := json.Marshal(j.Parameters)
	if err != nil {
		log.Infof("parameter marshal failed")
		return err
	}
	j.ParametersStr = string(b)
	return nil
}

func (j *AnalysisJob) TransferJsonToParameters() error {
	j.Parameters = make(map[string]interface{})
	err := json.Unmarshal([]byte(j.ParametersStr), &j.Parameters)
	if err != nil {
		log.Errorf("analysis_job parameter_str unmarshal failed: %v", err)
	}
	return err
}
