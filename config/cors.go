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

	if isAllowedOrigin(origin) {
		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Login")
			w.Header().Set("Access-Control-Allow-Methods", "POST,GET,DELETE,PUT")
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Max-Age", "3600")
			w.WriteHeader(http.StatusNoContent)
			return true
		}
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Origin", origin)
		return false
	}
	return false
}
