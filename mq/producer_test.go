package mq

import (
	"context"
	"github.com/airren/echo-bio-backend/model"
	"github.com/airren/echo-bio-backend/test_suit"
	"testing"
)

func TestPublishJob(t *testing.T) {
	test_suit.TestInit()
	job := &model.AnalysisJob{
		RecordMeta:    model.RecordMeta{Id: 11999},
		Name:          "test_job1",
		Algorithm:     "Rank abundance curve",
		Parameters:    nil,
		ParametersStr: "{\"data\":\"1637481894842994688\",\"group\":\"1637481894842994688\"}",
		Status:        "Pending",
		Description:   "this is a test",
	}
	PublishJob(context.TODO(), job)
}
func TestConsumerJob(t *testing.T) {
	test_suit.TestInit()
	ConsumerJob()
}
