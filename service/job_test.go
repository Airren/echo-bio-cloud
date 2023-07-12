package service

import (
	"context"
	"fmt"
	"github.com/airren/echo-bio-backend/model"
	"github.com/airren/echo-bio-backend/test_suit"
	"testing"
)

func TestGenerateAnalysisCommand(t *testing.T) {
	job := &model.AnalysisJob{
		RecordMeta: model.RecordMeta{},
		Name:       "",
		Algorithm:  "",
		Parameters: nil,
		//ParametersStr: "{\"data\":\"1637481894842994688\",\"group\":\"1637481894842994688\"}",
		ParametersStr: "{\"data\":\"1637481894842994688\"}",
		//ParametersStr: "{\"group\":\"1637481894842994688\"}",
		Outputs:     "",
		Status:      "",
		Description: "",
	}
	algo := &model.Algorithm{
		RecordMeta:  model.RecordMeta{},
		Name:        "",
		Label:       "",
		Image:       "",
		Description: "",
		Point:       0,
		Favourite:   0,
		Parameters:  nil,
		Command: `
export Bin=/opt2/Scripts/16S && 
export PATH=/opt2/software/Miniconda3/envs/micro/bin:$PATH &&
export LD_LIBRARY_PATH=/opt2/software/Miniconda3/envs/micro/lib:$LD_LIBRARY_PATH &&
export PYTHONPATH=/opt2/software/Miniconda3/envs/micro/lib/python3.10/site-packages/ &&
perl $Bin/rank_abund.pl {{ if index .Parameters "data"  }} -i {{index .Parameters "data"}} {{ end }}  {{ if index .Parameters "group"  }} -g {{index .Parameters "group" }} {{ end }} -o /tmp/analysis_job
`,
		DockerImage: "",
		Document:    "",
		GroupId:     "",
	}
	err := job.TransferJsonToParameters()
	if err != nil {
		t.Fatal(err)
	}
	commandStr, err := GenerateAlgorithmCommand(job, algo)
	fmt.Println(commandStr)

}

func TestGetCommandForAnalysisJob(t *testing.T) {
	test_suit.TestInit()
	job := &model.AnalysisJob{
		RecordMeta: model.RecordMeta{Id: 1651245206588100608},
		Name:       "",
		Algorithm:  "Rank abundance curve",
		Parameters: nil,
		//ParametersStr: "{\"data\":\"1637481894842994688\",\"group\":\"1637481894842994688\"}",
		ParametersStr: "{\"data\":\"1651246311095144448\"}",
		//ParametersStr: "{\"group\":\"1637481894842994688\"}",
		Outputs:     "",
		Status:      "",
		Description: "",
	}
	commandStr, err := GetCommandForAnalysisJob(context.TODO(), job)
	t.Logf("command: %v, err: %v", commandStr, err)
}
