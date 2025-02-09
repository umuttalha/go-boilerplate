package middleware

import (
	"context"
	"go-boilerplate/internal/service"
	"net/http"
	"strings"
)

func AuthMiddleware(svc *service.Service) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Skip auth check for certain paths
			if r.URL.Path == "/login" || r.URL.Path == "/register" || r.URL.Path == "/health" {
				next.ServeHTTP(w, r)
				return
			}

			// Get token from Authorization header
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "No authorization header", http.StatusUnauthorized)
				return
			}

			// Extract token
			splitToken := strings.Split(authHeader, "Bearer ")
			if len(splitToken) != 2 {
				http.Error(w, "Invalid token format", http.StatusUnauthorized)
				return
			}

			token := splitToken[1]
			session, err := svc.ValidateSession(token)
			if err != nil || session == nil {
				http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
				return
			}

			// Add session to request context
			ctx := r.Context()
			ctx = context.WithValue(ctx, "user_id", session.UserID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
