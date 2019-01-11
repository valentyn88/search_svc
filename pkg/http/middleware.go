package http

import (
	"context"
	"net/http"
	"text/template"
	"time"

	"github.com/satori/go.uuid"

	"github.com/valentyn88/search_svc"
)

const (
	queryCtxKey    = "query"
	authCookieName = "auth"
)

// QueryParamsMiddleware - middleware for query params.
func (h Handler) QueryParamsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := template.HTMLEscapeString(r.URL.Query().Get("q"))
		page := template.HTMLEscapeString(r.URL.Query().Get("page"))
		perPage := template.HTMLEscapeString(r.URL.Query().Get("per_page"))
		filter := template.HTMLEscapeString(r.URL.Query().Get("filter"))
		sort := template.HTMLEscapeString(r.URL.Query().Get("sort"))

		queryParam := search_svc.NewQueryParams(q, page, perPage, sort, filter)
		ctx := context.WithValue(r.Context(), queryCtxKey, queryParam)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// AuthMiddleware - middleware for authentication.
func (h Handler) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := r.Cookie("auth")
		if err != nil {
			id := uuid.NewV4()
			cookie := &http.Cookie{Name: authCookieName, Value: id.String(), Expires: time.Now().Add(7 * 24 * time.Hour)}
			http.SetCookie(w, cookie)
		}

		next.ServeHTTP(w, r)
	})
}
