package vo

import "github.com/airren/echo-bio-backend/model"

type AlgorithmVO struct {
	model.RecordMeta
	Name        string             `json:"name" gorm:"type:varchar(64);not null"`
	Label       string             `json:"label" gorm:"type:varchar(64);not null"`
	Image       string             `json:"image" gorm:"type:varchar(255);"`
	Description string             `json:"description" gorm:"type:text"`
	Price       int64              `json:"price" gorm:"type:int(11)"`
	Favourite   int64              `json:"favourite" gorm:"type:int(11)"`
	Parameters  []*model.Parameter `json:"parameter"`
}
