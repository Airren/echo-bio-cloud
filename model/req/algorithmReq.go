package req

import "github.com/airren/echo-bio-backend/model"

type AlgorithmReq struct {
	Id             string                 `json:"id"`
	Name           string                 `json:"name" `
	Label          string                 `json:"label" `
	Image          string                 `json:"image" `
	Description    string                 `json:"description" `
	Point          int64                  `json:"point" `
	Favourite      int64                  `json:"favourite"`
	Parameters     []*model.AlgoParameter `json:"parameters"`
	Command        string                 `json:"command"`
	DockerImage    string                 `json:"docker_image" `
	Document       string                 `json:"document"`
	Group          string                 `json:"group"`
	model.PageInfo `json:",inline"`
}
