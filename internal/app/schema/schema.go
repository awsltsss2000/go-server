package schema

// type ListResult struct {
// 	List       interface{}       `json:"list"`
// 	Pagination *PaginationResult `json:"pagination,omitempty"`
// }
//
// type PaginationResult struct {
// 	Total    int64 `json:"total"`
// 	Current  int   `json:"current"`
// 	PageSize int   `json:"page_size"`
// }

type PaginationParam struct {
	Pagination bool `form:"-"`
	OnlyCount  bool `form:"-"`
	Current    int  `form:"current,default=1"`
	PageSize   int  `form:"page_size,default=10" binding:"max=100"`
}

func (p *PaginationParam) GetCurrent() int {
	return p.Current
}

func (p *PaginationParam) GetPageSize() int {
	pageSize := p.PageSize
	if p.PageSize <= 0 {
		pageSize = 100
	}
	return pageSize
}

type OrderDirection int

const (
	OrderByASC OrderDirection = iota + 1
	OrderByDESC
)

func NewOrderField(key string, d OrderDirection) *OrderField {
	return &OrderField{
		Key:       key,
		Direction: d,
	}
}

type OrderField struct {
	Key       string
	Direction OrderDirection
}

type IDResult struct {
	ID uint64 `json:"id,string"`
}

type IDSResult struct {
	IDs []uint64 `json:"ids"`
}
