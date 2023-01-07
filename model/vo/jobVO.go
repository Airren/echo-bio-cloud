package vo

import "github.com/airren/echo-bio-backend/model"

type JobVO struct {
	*RecordMeta
	Name        string                 `json:"name"`
	Algorithm   string                 `json:"algorithm"`
	Parameters  map[string]interface{} `json:"parameters"`
	Outputs     string                 `json:"outputs"`
	Status      model.JobStatus        `json:"status"`
	Description string                 `json:"description"`
}
