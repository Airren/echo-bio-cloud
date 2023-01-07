package model

type Algorithm struct {
	RecordMeta

	Name        string           `json:"name" gorm:"type:varchar(64);unique;uniqueIndex;not null"`
	Label       string           `json:"label" gorm:"type:varchar(64);not null"`
	Image       string           `json:"image" gorm:"type:text;"`
	Description string           `json:"description" gorm:"type:text"`
	Point       int64            `json:"point" gorm:"type:int(11)"`
	Favourite   int64            `json:"favourite" gorm:"type:int(11)"`
	Parameters  []*AlgoParameter `gorm:"-"`
	Command     string           `json:"command" gorm:"text"`
	DockerImage string           `json:"docker_image" gorm:"type:varchar(256);not null"`

	Document string `json:"document" gorm:"text"`
	GroupId  string `json:"group_id" gorm:"type:varchar(64)"`
}

type AlgoGroup struct {
	RecordMeta
	Name  string `json:"name" gorm:"type:varchar(64); not null"`
	Label string `json:"label" gorm:"type:varchar(64); not null"`
}

type AlgoParameter struct {
	RecordMeta
	AlgorithmId uint64    `json:"algorithmId" gorm:"type:bigint"`
	Name        string    `json:"name" gorm:"type:varchar(64)"`
	Label       string    `json:"label" gorm:"type:varchar(64)"`
	Required    bool      `json:"required" gorm:"type:tinyint"`
	Description string    `json:"description" gorm:"type:text"`
	Type        ParamType `json:"type" gorm:"type:varchar(16)"`
	ValueList   string    `json:"value_list" gorm:"type:varchar(255)"`
}

type ParamType string

const (
	ParamString ParamType = "string"
	ParamFile   ParamType = "file"
	ParamRadio  ParamType = "radio"
	ParamSelect ParamType = "select"
)
