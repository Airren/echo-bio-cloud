package vo

type Group struct {
	*RecordMeta
	Name  string `json:"name" gorm:"type:varchar(64); not null"`
	Label string `json:"label" gorm:"type:varchar(64); not null"`
}
