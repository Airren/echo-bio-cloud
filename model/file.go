package model

import "fmt"

type File struct {
	RecordMeta
	Name        string `json:"name" gorm:"type:varchar(256); not null"`
	Description string `json:"description" gorm:"type:varchar(256);"`
	URLPath     string `json:"URLPath" gorm:"type:varchar(256); not null"`
	FileType    string `json:"file_type" gorm:"type:varchar(8)"`
	// allowed access by other user, 1: public ; 2: private
	Visibility int    `json:"visibility" gorm:"visibility;type:tinyint(1);not null"`
	MD5        string `json:"MD5" gorm:"type:varchar(256); not null"`
}

func (f *File) IsPublic() bool {
	return f.Visibility == 1
}

func (f *File) IsPrivate() bool {
	return f.Visibility == 2
}

func (f *File) SetURLPath() {
	if f.IsPublic() {
		f.URLPath = fmt.Sprintf("/api/v1/file/public/download/%d", f.Id)
	} else {
		f.URLPath = fmt.Sprintf("/api/v1/file/download/%d", f.Id)
	}
}
