package vo

type FileVO struct {
	*RecordMeta
	Name        string `json:"name" gorm:"type:varchar(256); not null"`
	Description string `json:"description" gorm:"type:varchar(256);"`
	FileType    string `json:"file_type"`
	IsPublic    bool   `json:"isPublic"`
	URLPath     string `json:"URLPath" gorm:"type:varchar(256); not null"`
}
