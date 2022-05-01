package req

import "github.com/airren/echo-bio-backend/model"

type JobReq struct {
	Id         int64
	Algorithm  string `json:"algorithm"`
	InputFile  string `json:"inputFile"`
	OutPutFile string
	Parameter  string

	model.PageInfo
}
