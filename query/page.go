package query

type Page struct {
	limit int
	skip  int
}

func Limit(limit int) Page {
	return Page{
		limit: limit,
		skip:  0,
	}
}

func (page Page) Skip(skip int) Page {
	page.skip = skip
	return page
}
