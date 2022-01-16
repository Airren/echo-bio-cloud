package req

import "github.com/airren/echo-bio-backend/model"

type AlgorithmReq struct {
	Label string `json:"label"`
	model.PageInfo
}
