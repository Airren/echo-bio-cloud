package model

import (
	"encoding/json"
	"github.com/airren/echo-bio-backend/global"
)

type JobStatus string

const (
	PENDING     JobStatus = "Pending"
	PROGRESSING JobStatus = "Progressing"
	COMPLETED   JobStatus = "Completed"
	FAILED      JobStatus = "Failed"
	CANCELED    JobStatus = "Canceled"
)

type Job struct {
	RecordMeta
	Name          string                 `gorm:"type:varchar(128); not null"`
	Algorithm     string                 `gorm:"type:varchar(64); not null"`
	Parameters    map[string]interface{} `gorm:"-"`
	ParametersStr string                 `gorm:"type:varchar(256); not null"`
	Outputs       string                 `gorm:"type:varchar(256); not null"`
	Status        JobStatus              `gorm:"type:varchar(32); not null"`
	Description   string                 `gorm:"type:text"`
}

func (j *Job) TransferParameter() error {
	b, err := json.Marshal(j.Parameters)
	if err != nil {
		global.Logger.Error("parameter marshal failed")
		return err
	}
	j.ParametersStr = string(b)
	return nil

}
