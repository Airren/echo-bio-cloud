package vo

const (
	DefaultPageSize = 10
	MaxPageSize     = 1000
)

type PageInfo struct {
	// 页数
	Page int `json:"page" form:"page" example:"1"`
	// 每页大小
	PageSize int    `json:"page_size" form:"page_size" example:"10"`
	Total    int    `json:"-" form:"-"`
	OrderBy  string `json:"-" form:"-"`
	Asc      bool   `json:"-" form:"-"`
}

func (pi *PageInfo) UpdatePageInfo() {
	// pagesize -1 means all
	if pi.PageSize == 0 || pi.PageSize < -2 {
		pi.PageSize = DefaultPageSize
	}
	if pi.PageSize > MaxPageSize {
		pi.PageSize = MaxPageSize
	}
	if pi.OrderBy == "" {
		pi.OrderBy = "created_at"
	}
	if pi.Page == 0 {
		pi.Page = 1
	}
}

func (pi *PageInfo) SetNoLimit() {
	pi.PageSize = -1
}

type RespPageInfo struct {
	// 页数
	Page int `json:"page" example:"1"`
	// 每页大小
	PageSize int `json:"page_size" form:"page_size" example:"10"`
	// 总数，超过50000条，显示50000
	Total   int    `json:"total"`
	OrderBy string `json:"order_by"`
	Asc     bool   `json:"asc"`
}

func NewRespPageInfo(pi *PageInfo) *RespPageInfo {
	if pi == nil {
		return nil
	}

	var vpi RespPageInfo
	vpi = RespPageInfo{
		Page:     pi.Page,
		PageSize: pi.PageSize,
		Total:    pi.Total,
		OrderBy:  pi.OrderBy,
		Asc:      pi.Asc,
	}
	return &vpi

}

type VO interface {
	SetErrCode()
	SetErrMsg()
	SetPageInfo()
	SetData(interface{})
}

type BaseVO struct {
	ErrCode  int           `json:"error_code"`
	ErrMsg   string        `json:"error_message"`
	PageInfo *RespPageInfo `json:"page_info,omitempty"`
	Data     interface{}   `json:"data"`
}

func (b *BaseVO) SetErrCode() {
}
func (b *BaseVO) SetErrMsg() {
}

func (b *BaseVO) SetPageInfo() {
}

func (b *BaseVO) SetData(data interface{}) {
	b.Data = data
}
