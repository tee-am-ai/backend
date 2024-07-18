package config

import (
	"net/http"
)

// Mendeklarasikan slice Origins yang berisi daftar origin yang diizinkan
var Origins = []string{
	"http://localhost:8080",
	"http://127.0.0.1:8080",
	"http://127.0.0.1:5500",
	"http://127.0.0.1:5501",
	"http://127.0.0.1:5503",
	"https://tee-am-ai.github.io",
	"https://tee-am-ai.vercel.app",
}

// Fungsi isAllowedOrigin memeriksa apakah origin yang diberikan diizinkan
func isAllowedOrigin(origin string) bool {
	for _, o := range Origins {
		if o == origin {
			return true
		}
	}
	return false
}

// Fungsi SetAccessControlHeaders mengatur header CORS untuk HTTP response
// SetAccessControlHeaders sets CORS (Cross-Origin Resource Sharing) headers on the HTTP response.
// It takes the HTTP response writer (`w http.ResponseWriter`) and the HTTP request (`r *http.Request`).
// It checks if the origin (`r.Header.Get("Origin")`) is allowed by calling `isAllowedOrigin`.
// If the origin is allowed:
// - For OPTIONS requests (`r.Method == http.MethodOptions`), it sets specific CORS headers:
//   - Access-Control-Allow-Credentials: true
//   - Access-Control-Allow-Headers: Content-Type, Login
//   - Access-Control-Allow-Methods: POST, GET, DELETE, PUT
//   - Access-Control-Allow-Origin: the origin received in the request
//   - Access-Control-Max-Age: 3600 (1 hour) to cache preflight response
//   - It writes a StatusNoContent (204) response header and returns true.
//
// - For other HTTP methods, it sets Access-Control-Allow-Credentials and Access-Control-Allow-Origin headers.
// - If the origin is not allowed, it returns false.
func SetAccessControlHeaders(w http.ResponseWriter, r *http.Request) bool {
	origin := r.Header.Get("Origin")

	// Check if the origin is allowed by calling isAllowedOrigin function.
	if isAllowedOrigin(origin) {
		if r.Method == http.MethodOptions {
			// Handle OPTIONS requests with CORS headers.
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Login")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, DELETE, PUT")
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Max-Age", "3600") // 1 hour
			w.WriteHeader(http.StatusNoContent)
			return true
		}

		// For other HTTP methods, set CORS headers without writing a response body.
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Origin", origin)
		return false
	}

	// If the origin is not allowed, return false.
	return false
}
