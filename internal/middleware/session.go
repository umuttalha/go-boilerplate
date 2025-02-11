package middleware

import (
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

var (
	// Store will hold all session data
	Store *sessions.CookieStore
)

func InitSession() {
	// In production, you should use an environment variable for the secret key
	Store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
	Store.Options = &sessions.Options{
		Path:     "/",
		HttpOnly: true,
	}
}

// SessionMiddleware checks if user is authenticated
func SessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := Store.Get(r, "session-name")

		// Skip auth check for certain paths
		if r.URL.Path == "/login" || r.URL.Path == "/register" || r.URL.Path == "/health" {
			next.ServeHTTP(w, r)
			return
		}

		// Check if user is authenticated
		auth, ok := session.Values["authenticated"].(bool)
		if !ok || !auth {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
