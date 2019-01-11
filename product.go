package search_svc

// ProductSearcher - interface for operations with Product.
type ProductSearcher interface {
	Search(qp QueryParam) ([]Product, int64, error)
}

// Product - domain object.
type Product struct {
	Title string  `json:"title"`
	Brand string  `json:"brand"`
	Price float32 `json:"price"`
	Stock int     `json:"stock"`
}
