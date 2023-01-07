package req

import "github.com/airren/echo-bio-backend/model"

type FileReq struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	Visibility     int    `json:"visibility"`
	Description    string `json:"description"`
	FileType       string `json:"fileType"`
	model.PageInfo `json:",inline"`
}
