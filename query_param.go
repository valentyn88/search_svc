package search_svc

import (
	"strconv"
	"strings"
)

const (
	sortFltrSeparator = ":"
	fltrSeparator     = ";"
	pageDef           = 1
	perPageDef        = 10
)

// QueryParam - params for query.
type QueryParam struct {
	Query   string
	Page    int
	PerPage int
	Sort    map[string]string
	Filter  map[string]string
}

// NewQueryParams - new instance of QueryParam.
func NewQueryParams(q, p, pp, sort, filter string) QueryParam {
	var qp QueryParam

	if q != "" {
		qp.Query = q
	}

	page, err := strconv.Atoi(p)
	if err != nil {
		page = pageDef
	}
	qp.Page = page

	perPage, err := strconv.Atoi(pp)
	if err != nil {
		perPage = perPageDef
	}
	qp.PerPage = perPage

	if sortVals := strings.Split(sort, sortFltrSeparator); len(sortVals) == 2 {
		qp.Sort = map[string]string{sortVals[0]: sortVals[1]}
	}

	if ftrVal := strings.Split(filter, fltrSeparator); len(ftrVal) >= 1 {
		filters := make(map[string]string, len(ftrVal))
		for _, fltr := range ftrVal {
			if f := strings.Split(fltr, sortFltrSeparator); len(f) == 2 {
				filters[f[0]] = f[1]
			}
		}
		qp.Filter = filters
	}

	return qp
}
