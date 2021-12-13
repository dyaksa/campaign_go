package helper

type Pagination struct {
	Limit      int         `form:"limit" json:"limit"`
	Page       int         `form:"page" json:"page"`
	TotalPages int         `json:"total_pages"`
	Sort       string      `form:"sort" json:"sort"`
	TotalRows  int64       `json:"total_rows"`
	Rows       interface{} `json:"rows"`
}

func (p *Pagination) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Pagination) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 5
	}
	return p.Limit
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Pagination) GetSort() string {
	if p.Sort == "" {
		p.Sort = "id desc"
	}
	return p.Sort
}
