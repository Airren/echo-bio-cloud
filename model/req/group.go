package req

import "github.com/airren/echo-bio-backend/model"

type Group struct {
	model.RecordMeta
	Id    string `json:"id"`
	Name  string `json:"name" gorm:"type:varchar(64); not null"`
	Label string `json:"label" gorm:"type:varchar(64); not null"`
}
