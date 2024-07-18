package routes

import (
	"net/http" // Package untuk melakukan operasi-operasi terkait HTTP seperti membuat server, mengirim permintaan, dan menerima respons HTTP

	"github.com/tee-am-ai/backend/config"                // Package yang mungkin berisi konfigurasi aplikasi
	controller "github.com/tee-am-ai/backend/controller" // Package yang mungkin berisi definisi-definisi controller atau handler untuk mengelola permintaan HTTP
	"github.com/tee-am-ai/backend/helper"                // Package yang mungkin berisi fungsi-fungsi bantuan (helper functions)
)

func URL(w http.ResponseWriter, r *http.Request) {
	// SetAccessControlHeaders handles CORS preflight requests. If true, return early.
	if config.SetAccessControlHeaders(w, r) {
		return
	}

	// Check for MongoDB connection error. If not nil, return Internal Server Error response.
	if config.ErrorMongoconn != nil {
		helper.ErrorResponse(w, r, http.StatusInternalServerError, "Internal Server Error", "kesalahan server : database, "+config.ErrorMongoconn.Error())
		return
	}

	// Extract HTTP method and request path from the incoming request.
	var method, path string = r.Method, r.URL.Path

	// Route the request based on HTTP method and path.
	switch {
	case method == "GET" && path == "/":
		Home(w, r) // Handle GET request to root path ("/") by calling Home function.
	case method == "POST" && path == "/signup":
		controller.SignUp(config.Mongoconn, "users", w, r) // Handle POST request to "/signup" path.
	case method == "POST" && path == "/login":
		controller.LogIn(config.Mongoconn, w, r, config.GetEnv("PASETOPRIVATEKEY")) // Handle POST request to "/login" path.
	case method == "POST" && path == "/chat":
		controller.Chat(w, r, config.GetEnv("TOKENMODEL")) // Handle POST request to "/chat" path.
	default:
		// If no matching route is found, return a Not Found response.
		helper.ErrorResponse(w, r, http.StatusNotFound, "Not Found", "The requested resource was not found")
	}
}

func Home(respw http.ResponseWriter, req *http.Request) {
	// Mendefinisikan sebuah map yang berisi informasi repository GitHub dan pesan.
	resp := map[string]string{
		"github_repo": "https://github.com/tee-am-ai/backend",
		"message":     "Ampun puh sepuh, aku mah masih pemula 🙏",
	}

	// Memanggil fungsi WriteJSON dari package helper untuk menulis respons JSON.
	helper.WriteJSON(respw, http.StatusOK, resp)
}
