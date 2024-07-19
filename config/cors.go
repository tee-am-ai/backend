package config

// Mengimpor paket "net/http" dari standar library Go.
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
