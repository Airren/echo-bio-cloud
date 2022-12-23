package vo

type FileVO struct {
	*RecordMeta
	Name        string `json:"name" gorm:"type:varchar(256); not null"`
	Description string `json:"description" gorm:"type:varchar(256);"`
	FileType    string `json:"file_type"`
	Visibility  int    `json:"visibility"`
	URLPath     string `json:"URLPath" gorm:"type:varchar(256); not null"`
}
