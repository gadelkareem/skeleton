package paginator

import (
	"fmt"
	"math"
	"strings"

	"backend/kernel"
	"github.com/google/jsonapi"
)

type Paginator struct {
	Models                           []interface{}
	Filter, uri                      string
	Size, Limit, Offset, currentPage int
	Sort                             map[string]string
}

const (
	ASC  = "ASC"
	DESC = "DESC"
)

func NewPaginator(pg map[string]int, sort, filter, uri string) *Paginator {
	p := &Paginator{uri: uri}

	var curPg, limit int
	if v, ok := pg["after"]; ok {
		curPg = v
	}
	if v, ok := pg["before"]; ok {
		curPg = v - 1
	}
	if v, ok := pg["size"]; ok {
		limit = v
	}

	if curPg < 1 {
		curPg = 1
	}
	p.currentPage = curPg

	if limit < 1 && limit != -1 {
		limit = kernel.ListLimit
	}

	p.Limit = limit
	p.Offset = (p.currentPage - 1) * limit

	sl := strings.Split(sort, ",")
	p.Sort = make(map[string]string)
	for _, s := range sl {
		order := ASC
		if strings.HasPrefix(s, "-") {
			order = DESC
			s = s[1:]
		}
		p.Sort[s] = order
	}

	p.Filter = filter

	return p
}

func (p *Paginator) Optimize() {
	if p.Limit == -1 {
		p.Limit = len(p.Models)
		p.Size = len(p.Models)
	}
}

func (p *Paginator) Links() *jsonapi.Links {
	pg := fmt.Sprintf("%s%s?page[size]=%d", kernel.App.APIURL, p.uri, kernel.ListLimit)
	l := jsonapi.Links{
		jsonapi.KeyFirstPage: fmt.Sprintf("%s&page[after]=1", pg),
		jsonapi.KeyLastPage:  fmt.Sprintf("%s&page[after]=%d", pg, p.TotalPages()),
	}
	if p.currentPage > 1 {
		l[jsonapi.KeyPreviousPage] = fmt.Sprintf("%s&page[before]=%d", pg, p.CurrentPage())
	}
	if p.currentPage < p.TotalPages() {
		l[jsonapi.KeyNextPage] = fmt.Sprintf("%s&page[after]=%d", pg, p.CurrentPage()+1)
	}

	return &l
}

func (p *Paginator) Meta() *jsonapi.Meta {
	pg := map[string]interface{}{
		"total":  p.Size,
		"size":   p.Limit,
		"before": p.CurrentPage(),
	}
	sort := ""
	for c, d := range p.Sort {
		if d == DESC {
			sort += "-"
		}
		sort += c
	}
	if p.CurrentPage()+1 > p.TotalPages() {
		pg["after"] = 0
	} else {
		pg["after"] = p.CurrentPage() + 1
	}

	return &jsonapi.Meta{"page": pg, "sort": sort, "filter": p.Filter}
}

func (p *Paginator) TotalPages() int {
	return int(math.Ceil(float64(p.Size) / float64(p.Limit)))
}

func (p *Paginator) CurrentPage() int {
	if p.TotalPages() < p.currentPage {
		p.currentPage = 1
	}
	return p.currentPage
}

func (p *Paginator) Slice() {
	p.Size = len(p.Models)
	limit := p.Limit
	if limit < 0 || limit > p.Size {
		limit = p.Size
	}
	if p.Offset < 0 {
		p.Offset = 0
	} else if p.Offset > p.Size {
		p.Offset = p.Size
	}
	// Calculate the boundary for slicing. 'end' ensures that the slice does not exceed the array bounds,
	// preventing off-by-one errors.
	end := p.Offset + limit
	if end > p.Size {
		end = p.Size
	}
	p.Models = p.Models[p.Offset:end]
}
