package http

import (
	h "net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var mockHandler = func(w h.ResponseWriter, r *h.Request) {
	w.Write([]byte("OK"))
}

func TestHandler_AuthMiddleware(t *testing.T) {
	req, err := h.NewRequest("GET", "/products", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler := Handler{SearchService: nil}
	th := handler.AuthMiddleware(h.HandlerFunc(mockHandler))

	rr := httptest.NewRecorder()

	th.ServeHTTP(rr, req)
	cookies, ok := rr.HeaderMap["Set-Cookie"]
	if !ok {
		t.Fatal("couldn't get cookie AuthMiddleware")
	}

	for _, c := range cookies {
		if strings.Contains(c, authCookieName) {
			return
		}
	}

	t.Fatal("auth cookie was not set AuthMiddleware")
}
