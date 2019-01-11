package mock

import "github.com/valentyn88/search_svc"

// SimpleSearchSvc - implementation of ProductSearcher for test.
type SimpleSearchSvc struct {
}

// Search - search products.
func (s SimpleSearchSvc) Search(qp search_svc.QueryParam) ([]search_svc.Product, int64, error) {
	pp := products()

	return pp, int64(len(pp)), nil
}
