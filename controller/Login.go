package controller

import (
	"encoding/hex"  // Mengimpor paket encoding/hex untuk encoding dan decoding data hexadecimal
	"encoding/json" // Mengimpor paket encoding/json untuk bekerja dengan JSON
	"net/http"      // Mengimpor paket net/http untuk menangani HTTP request dan response

	"github.com/badoux/checkmail"              // Mengimpor paket checkmail untuk memvalidasi alamat email
	"github.com/tee-am-ai/backend/helper"      // Mengimpor package helper dari aplikasi backend
	model "github.com/tee-am-ai/backend/model" // Mengimpor package model dari aplikasi backend
	"go.mongodb.org/mongo-driver/mongo"        // Mengimpor paket mongo-driver untuk bekerja dengan MongoDB
	"golang.org/x/crypto/argon2"               // Mengimpor paket argon2 untuk hashing kata sandi dengan algoritma Argon2
)

// LogIn adalah fungsi untuk menangani permintaan login pengguna.
func LogIn(db *mongo.Database, respw http.ResponseWriter, req *http.Request, privatekey string) {
	// Deklarasi variabel untuk menyimpan data pengguna yang dikirimkan dalam body permintaan HTTP
	var user model.User

	// Menguraikan dan memasukkan data JSON dari body permintaan ke dalam struktur model.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		// Jika terjadi kesalahan dalam parsing data, kirim respons dengan status Bad Request dan pesan kesalahan
		helper.ErrorResponse(respw, req, http.StatusBadRequest, "Bad Request", "error parsing request body "+err.Error())
		return
	}

	// Memastikan bahwa email dan password tidak kosong
	if user.Email == "" || user.Password == "" {
		helper.ErrorResponse(respw, req, http.StatusBadRequest, "Bad Request", "mohon untuk melengkapi data")
		return
	}

	// Validasi format email menggunakan package checkmail
	if err = checkmail.ValidateFormat(user.Email); err != nil {
		helper.ErrorResponse(respw, req, http.StatusBadRequest, "Bad Request", "email tidak valid")
		return
	}

	// Mendapatkan data pengguna berdasarkan email dari database
	existsDoc, err := helper.GetUserFromEmail(user.Email, db)
	if err != nil {
		helper.ErrorResponse(respw, req, http.StatusInternalServerError, "Internal Server Error", "kesalahan server : get email "+err.Error())
		return
	}

	// Mendekode salt dari string heksadesimal menjadi slice byte
	salt, err := hex.DecodeString(existsDoc.Salt)
	if err != nil {
		helper.ErrorResponse(respw, req, http.StatusInternalServerError, "Internal Server Error", "kesalahan server : salt")
		return
	}

	// Menggunakan argon2 untuk menghasilkan hash dari password yang diberikan
	hash := argon2.IDKey([]byte(user.Password), salt, 1, 64*1024, 4, 32)

	// Membandingkan hasil hash dengan hash yang tersimpan di database
	if hex.EncodeToString(hash) != existsDoc.Password {
		helper.ErrorResponse(respw, req, http.StatusUnauthorized, "Unauthorized", "password salah")
		return
	}

	// Jika autentikasi berhasil, menghasilkan token menggunakan helper.Encode
	tokenstring, err := helper.Encode(user.ID, user.Email, privatekey)
	if err != nil {
		helper.ErrorResponse(respw, req, http.StatusInternalServerError, "Internal Server Error", "kesalahan server : token")
		return
	}

	// Persiapan respons yang akan dikirimkan kembali kepada klien
	resp := map[string]interface{}{
		"status":  "success",
		"message": "login berhasil",
		"token":   tokenstring,
		"data": map[string]string{
			"email":       existsDoc.Email,
			"namalengkap": existsDoc.NamaLengkap,
		},
	}

	// Mengirimkan respons dalam format JSON dengan status OK (200)
	helper.WriteJSON(respw, http.StatusOK, resp)
}
