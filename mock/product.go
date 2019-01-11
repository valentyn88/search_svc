package mock

import "github.com/valentyn88/search_svc"

func products() []search_svc.Product {
	products := []search_svc.Product{
		search_svc.Product{
			Title: "Adidas sneakers",
			Brand: "Adidas",
			Price: 29.99,
			Stock: 5,
		},
		search_svc.Product{
			Title: "Nike sneakers",
			Brand: "Nike",
			Price: 39.99,
			Stock: 10,
		},
		search_svc.Product{
			Title: "Asics sneakers",
			Brand: "Asics",
			Price: 19.99,
			Stock: 15,
		},
	}

	return products
}
