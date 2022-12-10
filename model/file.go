package model

type File struct {
	RecordMeta
	Name          string `json:"name" gorm:"type:varchar(256); not null"`
	Description   string `json:"description" gorm:"type:varchar(256);"`
	URLPath       string `json:"URLPath" gorm:"type:varchar(256); not null"`
	URLExpires    int64  `json:"URLExpires" gorm:"type:long"`
	AlgorithmName string `json:"algorithmName" gorm:"type:varchar(256); not null"`
}
