package httpserver

import "net/http"

func enableCorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Allow requests from any origin
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// Allow specified HTTP methods
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, HEAD, OPTIONS")

		// Allow specified headers
		w.Header().Set("Access-Control-Allow-Headers", "*")

		// Allow credentials
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// Continue with the next handler
		next.ServeHTTP(w, r)
	})
}
