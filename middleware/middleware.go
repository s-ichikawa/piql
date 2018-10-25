package middleware

import (
	"context"
	"net/http"
)

type contextKey string

const TokenContextKey contextKey = "key"

func AuthMiddleware(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-USER-TOKEN")
		ctx := context.WithValue(r.Context(), TokenContextKey, token)

		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
