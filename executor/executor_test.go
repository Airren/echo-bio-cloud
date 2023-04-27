package executor

import (
	"context"
	"fmt"
	"github.com/airren/echo-bio-backend/model"
	"github.com/airren/echo-bio-backend/test_suit"
	"testing"
)

func init() {
	test_suit.TestInit()
	KubeInitializer()
}

func TestExecutor_CreateJob(t *testing.T) {
	fmt.Println("hello")
	e := Executor{}
	err := e.CreateJob(context.TODO(), &model.AnalysisJob{
		RecordMeta: model.RecordMeta{
			Id:        1637482017367003136,
			AccountId: "0c6d1c37-286f-4648-9472-dd658874677f",
			Org:       "",
		},
		Name:          "分析任务1",
		Algorithm:     "Rank abundance curve",
		Parameters:    nil,
		ParametersStr: "{\"data\":\"1637481894842994688\",\"group\":\"1637481894842994688\"}",
		Outputs:       "",
		Status:        "",
		Description:   "",
	})
	if err != nil {
		return
	}
}
