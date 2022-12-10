package req

type FileReq struct {
	Id          string `json:"id"`
	Name        string `json:"name" gorm:"type:varchar(256); not null"`
	Description string `json:"description" gorm:"type:varchar(256);"`
	URLPath     string `json:"URLPath" gorm:"type:varchar(256); not null"`
}

type ListFileByIdsReq struct {
	Ids []uint64 `json:"ids"`
}
