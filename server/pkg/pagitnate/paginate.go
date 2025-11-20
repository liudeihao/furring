package pagitnate

type Pagination struct {
    Page     int `json:"page"`
    PageSize int `json:"page_size"`
}

func (p *Pagination) Limit() int {
    return p.PageSize
}

func (p *Pagination) Offset() int {
    return (p.Page - 1) * p.PageSize
}
