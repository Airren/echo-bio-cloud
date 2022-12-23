package req

import "github.com/airren/echo-bio-backend/model"

type FileReq struct {
	Id             string `json:"id"`
	Name           string `json:"name" gorm:"type:varchar(256); not null"`
	Visibility     int    `json:"visibility"`
	Description    string `json:"description" gorm:"type:varchar(256);"`
	FileType       string `json:"fileType"`
	model.PageInfo `json:",inline"`
}
