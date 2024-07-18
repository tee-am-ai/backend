package routes

import (
	"net/http" // Package untuk melakukan operasi-operasi terkait HTTP seperti membuat server, mengirim permintaan, dan menerima respons HTTP

	"github.com/tee-am-ai/backend/config"                // Package yang mungkin berisi konfigurasi aplikasi
	controller "github.com/tee-am-ai/backend/controller" // Package yang mungkin berisi definisi-definisi controller atau handler untuk mengelola permintaan HTTP
	"github.com/tee-am-ai/backend/helper"                // Package yang mungkin berisi fungsi-fungsi bantuan (helper functions)
)

func URL(w http.ResponseWriter, r *http.Request) {
	if config.SetAccessControlHeaders(w, r) {
		return // If it's a preflight request, return early.
	}

	if config.ErrorMongoconn != nil {
		helper.ErrorResponse(w, r, http.StatusInternalServerError, "Internal Server Error", "kesalahan server : database, "+config.ErrorMongoconn.Error())
		return
	}

	var method, path string = r.Method, r.URL.Path
	switch {
	case method == "GET" && path == "/":
		Home(w, r)
	case method == "POST" && path == "/signup":
		controller.SignUp(config.Mongoconn, "users", w, r)
	case method == "POST" && path == "/login":
		controller.LogIn(config.Mongoconn, w, r, config.GetEnv("PASETOPRIVATEKEY"))
	case method == "POST" && path == "/chat":
		controller.Chat(w, r, config.GetEnv("TOKENMODEL"))
	default:
		helper.ErrorResponse(w, r, http.StatusNotFound, "Not Found", "The requested resource was not found")
	}
}

func Home(respw http.ResponseWriter, req *http.Request) {
	resp := map[string]string{
		"github_repo": "https://github.com/tee-am-ai/backend",
		"message":     "Ampun puh sepuh, aku mah masih pemula 🙏",
	}
	helper.WriteJSON(respw, http.StatusOK, resp)
}
