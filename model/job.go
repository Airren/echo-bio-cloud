package model

type Job struct {
	RecordMeta
	Algorithm  string `gorm:"type:varchar(64); not null"`
	InputFile  string `gorm:"type:varchar(256); not null"`
	OutPutFile string `gorm:"type:varchar(256); not null"`
	Parameter  string `gorm:"type:varchar(256); not null"`
}
