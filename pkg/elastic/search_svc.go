package elastic

import (
	"context"
	"reflect"

	"github.com/olivere/elastic"

	"github.com/valentyn88/search_svc"
)

const (
	productIndexKey = "product"
	productIndex    = `
{
	"settings":{
		"number_of_shards":1,
		"number_of_replicas":0
	},
	"mappings":{
		"doc":{
			"properties":{
				"title":{
					"type":"text"
				},
				"brand":{
					"type":"text"
				},
				"price":{
					"type":"float"
				},
                "stock":{
                    "type":"integer"
                }
			}
		}
	}
}`
)

// SearchService - search service.
type SearchService struct {
	ElasticClient *elastic.Client
}

// Search - search products.
func (ss SearchService) Search(qp search_svc.QueryParam) ([]search_svc.Product, int64, error) {
	var querySet = make([]elastic.Query, 0)

	if qp.Query != "" {
		querySet = append(querySet, elastic.NewTermQuery("title", qp.Query))
	}

	if len(qp.Filter) > 0 {
		for k, v := range qp.Filter {
			querySet = append(querySet, elastic.NewMatchQuery(k, v))
		}
	}

	search := ss.ElasticClient.Search().Index(productIndexKey)
	if len(querySet) > 0 {
		search.Query(elastic.NewBoolQuery().Must(querySet...))
	}

	if len(qp.Sort) > 0 {
		var sortVal bool
		if qp.Sort["price"] == "asc" {
			sortVal = true
		}
		search.Sort("price", sortVal)
	}

	search.From(qp.Page - 1).Size(qp.PerPage).Pretty(true)

	searchRes, err := search.Do(context.Background())
	if err != nil {
		return nil, 0, err
	}

	var (
		product search_svc.Product
		pp      []search_svc.Product
	)

	for _, prod := range searchRes.Each(reflect.TypeOf(product)) {
		p, ok := prod.(search_svc.Product)
		if ok {
			pp = append(pp, p)
		}
	}

	return pp, searchRes.TotalHits(), nil
}

// InitFixtures - init default products
func (ss SearchService) InitFixtures() error {
	indexExists, err := ss.ElasticClient.IndexExists(productIndexKey).Do(context.Background())
	if err != nil {
		return err
	}

	if !indexExists {
		if _, err = ss.ElasticClient.CreateIndex(productIndexKey).Body(productIndex).
			Do(context.Background()); err != nil {
			return err
		}
	}

	prod1 := search_svc.Product{Title: "Adidas sneakers", Brand: "Adidas", Price: 59.99, Stock: 5}
	if _, err = ss.ElasticClient.Index().
		Index(productIndexKey).
		Type("doc").
		Id("1").
		BodyJson(prod1).
		Do(context.Background()); err != nil {
		return err
	}

	prod2 := search_svc.Product{Title: "Nike sneakers", Brand: "Nike", Price: 79.99, Stock: 5}
	if _, err = ss.ElasticClient.Index().
		Index(productIndexKey).
		Type("doc").
		Id("2").
		BodyJson(prod2).
		Do(context.Background()); err != nil {
		return err
	}

	prod3 := search_svc.Product{Title: "Puma sneakers", Brand: "Puma", Price: 89.99, Stock: 5}
	if _, err = ss.ElasticClient.Index().
		Index(productIndexKey).
		Type("doc").
		Id("3").
		BodyJson(prod3).
		Do(context.Background()); err != nil {
		return err
	}

	prod4 := search_svc.Product{Title: "Puma tshirt", Brand: "Puma", Price: 19.99, Stock: 5}
	if _, err = ss.ElasticClient.Index().
		Index(productIndexKey).
		Type("doc").
		Id("4").
		BodyJson(prod4).
		Do(context.Background()); err != nil {
		return err
	}

	prod5 := search_svc.Product{Title: "Nike tshirt", Brand: "Nike", Price: 29.99, Stock: 5}
	if _, err = ss.ElasticClient.Index().
		Index(productIndexKey).
		Type("doc").
		Id("5").
		BodyJson(prod5).
		Do(context.Background()); err != nil {
		return err
	}

	prod6 := search_svc.Product{Title: "Asics tshirt", Brand: "Asics", Price: 9.99, Stock: 5}
	if _, err = ss.ElasticClient.Index().
		Index(productIndexKey).
		Type("doc").
		Id("6").
		BodyJson(prod6).
		Do(context.Background()); err != nil {
		return err
	}

	_, err = ss.ElasticClient.Flush().Index(productIndexKey).Do(context.Background())

	return err
}
