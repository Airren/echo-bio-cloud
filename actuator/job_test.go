package actuator

import (
	"github.com/airren/echo-bio-backend/model"
	"testing"
)

func TestStartPaint(t *testing.T) {

	job := model.Job{
		RecordMeta: model.RecordMeta{Id: 210582357203, AccountId: "89757"},
		Algorithm:  "pie",
		InputFile:  "pie-data.xls",
		OutPutFile: "",
		Parameter:  "",
	}
	_, err := StartPaint(job)
	if err != nil {
		t.Fatalf("job faile %v", err)
	}
}
