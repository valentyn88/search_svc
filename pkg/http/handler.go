package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/valentyn88/search_svc"
)

// Handler - struct for handlers.
type Handler struct {
	SearchService search_svc.ProductSearcher
}

// ProductsResp - response for products.
type ProductsResp struct {
	Page     int                  `json:"page"`
	PerPage  int                  `json:"perPage"`
	Count    int64                `json:"count"`
	Products []search_svc.Product `json:"products"`
}

// Products - list of products handler.
func (h Handler) Products(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	queryParams, ok := ctx.Value(queryCtxKey).(search_svc.QueryParam)
	if !ok {
		log.Println("couldn't get query params from context")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	products, count, err := h.SearchService.Search(queryParams)
	if err != nil {
		log.Printf("couldn't search products by query params: %v error: %s\n", queryParams, err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := json.Marshal(ProductsResp{
		Page:     queryParams.Page,
		PerPage:  queryParams.PerPage,
		Count:    count,
		Products: products})
	if err != nil {
		log.Printf("couldn't marshal products response error: %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	if _, err := w.Write(resp); err != nil {
		log.Printf("couldn't write response error: %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
	}
}
