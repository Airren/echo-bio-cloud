package req

import "strconv"

type IdsReq struct {
	Ids []string `json:"ids"`
}

func (idsReq *IdsReq) IdsToInt() []int64 {
	idsList := make([]int64, len(idsReq.Ids))
	for i, v := range idsReq.Ids {
		id, err := strconv.ParseInt(v, 10, 64)
		if err == nil {
			idsList[i] = id
		}
	}
	return idsList
}
