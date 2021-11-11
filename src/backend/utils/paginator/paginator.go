package paginator

import (
    "fmt"
    "math"
    "strings"

    "backend/kernel"
    "github.com/google/jsonapi"
)

type Paginator struct {
    Models                                       []interface{}
    Filter, uri                                  string
    Size, Limit, Offset, currentPage, totalPages int
    Sort                                         map[string]string
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
        "after":  p.CurrentPage() + 1,
    }
    sort := ""
    for c, d := range p.Sort {
        if d == DESC {
            sort += "-"
        }
        sort += c
    }
    return &jsonapi.Meta{"page": pg, "sort": sort, "filter": p.Filter}
}

func (p *Paginator) TotalPages() int {
    if p.totalPages != 0 {
        return p.totalPages
    }
    p.totalPages = int(math.Ceil(float64(p.Size) / float64(p.Limit)))
    return p.totalPages
}

func (p *Paginator) CurrentPage() int {
    if p.TotalPages() < p.currentPage {
        p.currentPage = 1
    }
    return p.currentPage
}
