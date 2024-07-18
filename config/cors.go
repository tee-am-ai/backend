package config

import (
	"net/http" // Mengimpor package net/http untuk menangani HTTP request dan response
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
	origin := r.Header.Get("Origin") // Mendapatkan nilai header Origin dari request

	if isAllowedOrigin(origin) { // Memeriksa apakah origin diizinkan
		if r.Method == http.MethodOptions { // Jika metode HTTP adalah OPTIONS (preflight request)
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Login")
			w.Header().Set("Access-Control-Allow-Methods", "POST,GET,DELETE,PUT")
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Max-Age", "3600")
			w.WriteHeader(http.StatusNoContent) // Mengirim response tanpa konten
			return true                         // Mengembalikan true, menandakan preflight request ditangani
		}
		// Mengatur header CORS untuk request utama
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Origin", origin)
		return false // Mengembalikan false, menandakan request utama masih perlu ditangani lebih lanjut
	}
	return false // Mengembalikan false jika origin tidak diizinkan
}
