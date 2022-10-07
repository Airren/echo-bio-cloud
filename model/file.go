package model

type File struct {
	RecordMeta
	Name        string `json:"name" gorm:"type:varchar(256); not null"`
	Description string `json:"description" gorm:"type:varchar(256);"`
	URLPath     string `json:"URLPath" gorm:"type:varchar(256); not null"`
}
