package auth

import (
	"net/http"
)

func OAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Placeholder for OAuth validation logic
		next.ServeHTTP(w, r)
	})
}
