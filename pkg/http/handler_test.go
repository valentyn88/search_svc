package http

import (
	"context"
	h "net/http"
	"net/http/httptest"
	"testing"

	"github.com/valentyn88/search_svc"

	"github.com/valentyn88/search_svc/mock"
)

func TestHandler_Products(t *testing.T) {
	req, err := h.NewRequest("GET", "/products?q=sneakers&page=2&per_page=5&sort=price:asc&filter=brand:Adidas", nil)
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.WithValue(req.Context(), queryCtxKey, search_svc.QueryParam{
		Query:   "sneakers",
		Page:    2,
		PerPage: 5,
		Sort:    map[string]string{"price": "asc"},
		Filter:  map[string]string{"brand": "Adidas"},
	})

	ss := mock.SimpleSearchSvc{}
	handler := Handler{SearchService: ss}
	th := h.HandlerFunc(handler.Products)

	rr := httptest.NewRecorder()

	th.ServeHTTP(rr, req.WithContext(ctx))

	expectedResp := `{"page":2,"perPage":5,"count":3,"products":[{"title":"Adidas sneakers","brand":"Adidas","price":29.99,"stock":5},{"title":"Nike sneakers","brand":"Nike","price":39.99,"stock":10},{"title":"Asics sneakers","brand":"Asics","price":19.99,"stock":15}]}`
	if expectedResp != rr.Body.String() {
		t.Fatalf("expected %v and got %v are not equal", expectedResp, rr.Body.String())
	}
}
