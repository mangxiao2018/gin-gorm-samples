package page

type Page struct {
	Page     int
	PageSize int
}

func (p *Page) GetPage() *Page {
	if p.Page <= 0 {
		p.Page = 1
	}

	switch {
	case p.PageSize > 100:
		p.PageSize = 100
	case p.PageSize <= 0:
		p.PageSize = 10
	}
	offset := (p.Page - 1) * p.PageSize
	return &Page{Page: offset, PageSize: p.PageSize}
}
