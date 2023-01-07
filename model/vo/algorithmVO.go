package vo

import "github.com/airren/echo-bio-backend/model"

type AlgorithmVO struct {
	*RecordMeta
	Name        string                 `json:"name" `
	Label       string                 `json:"label" `
	Image       string                 `json:"image" `
	Description string                 `json:"description" `
	Point       int64                  `json:"point" `
	Favourite   int64                  `json:"favourite" `
	Parameters  []*model.AlgoParameter `json:"parameters"`
	Command     string                 `json:"command"`
	DockerImage string                 `json:"docker_image" `
	Document    string                 `json:"document"`
	GroupId     string                 `json:"group_id"`
}
