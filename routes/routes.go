package routes

import (
	"net/http" // Package untuk melakukan operasi-operasi terkait HTTP seperti membuat server, mengirim permintaan, dan menerima respons HTTP

	"github.com/tee-am-ai/backend/config"                // Package yang mungkin berisi konfigurasi aplikasi
	controller "github.com/tee-am-ai/backend/controller" // Package yang mungkin berisi definisi-definisi controller atau handler untuk mengelola permintaan HTTP
	"github.com/tee-am-ai/backend/helper"                // Package yang mungkin berisi fungsi-fungsi bantuan (helper functions)
)

// URL adalah handler fungsi untuk menangani berbagai request berdasarkan metode dan path URL.
func URL(w http.ResponseWriter, r *http.Request) {
	// Memeriksa dan mengatur header kontrol akses. Jika diatur, kembalikan dari fungsi.
	if config.SetAccessControlHeaders(w, r) {
		return
	}

	// Memeriksa koneksi ke database. Jika terjadi kesalahan koneksi, kirimkan respons kesalahan internal server.
	if config.ErrorMongoconn != nil {
		helper.ErrorResponse(w, r, http.StatusInternalServerError, "Internal Server Error", "kesalahan server : database, "+config.ErrorMongoconn.Error())
		return
	}

	// Mendapatkan metode request (GET, POST, dll.) dan path URL.
	var method, path string = r.Method, r.URL.Path

	// Menggunakan switch-case untuk menangani berbagai kombinasi metode dan path URL.
	switch {
	// Menangani permintaan GET ke path root ("/") dengan memanggil fungsi Home.
	case method == "GET" && path == "/":
		Home(w, r)
	// Menangani permintaan POST ke path "/signup" dengan memanggil fungsi SignUp.
	case method == "POST" && path == "/signup":
		controller.SignUp(config.Mongoconn, "users", w, r)
	// Menangani permintaan POST ke path "/login" dengan memanggil fungsi LogIn.
	case method == "POST" && path == "/login":
		controller.LogIn(config.Mongoconn, w, r, config.GetEnv("PASETOPRIVATEKEY"))
	// Menangani permintaan POST ke path "/chat" dengan memanggil fungsi Chat.
	case method == "POST" && path == "/chat":
		controller.Chat(w, r, config.GetEnv("TOKENMODEL"))
	// Menangani semua permintaan lain yang tidak cocok dengan kasus di atas, mengirimkan respons 404 Not Found.
	default:
		helper.ErrorResponse(w, r, http.StatusNotFound, "Not Found", "The requested resource was not found")
	}
}

// func Home(respw http.ResponseWriter, req *http.Request) {
// 	// Mendefinisikan sebuah map yang berisi informasi repository GitHub dan pesan.
// 	resp := map[string]string{
// 		"github_repo": "https://github.com/tee-am-ai/backend",
// 		"message":     "Ampun puh sepuh, aku mah masih pemula 🙏",
// 	}

// 	// Memanggil fungsi WriteJSON dari package helper untuk menulis respons JSON.
// 	helper.WriteJSON(respw, http.StatusOK, resp)
// }
