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
// LogIn adalah handler fungsi untuk menangani permintaan login.
func LogIn(db *mongo.Database, respw http.ResponseWriter, req *http.Request, privatekey string) {
	// Mendeklarasikan variabel user sebagai tipe model.User
	var user model.User

	// Mendekode body permintaan JSON ke dalam variabel user
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		// Jika terjadi kesalahan dalam parsing data, kirim respons dengan status Bad Request dan pesan kesalahan
		helper.ErrorResponse(respw, req, http.StatusBadRequest, "Bad Request", "error parsing request body "+err.Error())
		return
	}

	// Memastikan bahwa email dan password tidak kosong
	if user.Email == "" || user.Password == "" {
		// Jika email atau password kosong, kirim respons dengan status Bad Request dan pesan kesalahan
		helper.ErrorResponse(respw, req, http.StatusBadRequest, "Bad Request", "mohon untuk melengkapi data")
		return
	}

	// Validasi format email menggunakan package checkmail
	if err = checkmail.ValidateFormat(user.Email); err != nil {
		// Jika format email tidak valid, kirim respons dengan status Bad Request dan pesan kesalahan
		helper.ErrorResponse(respw, req, http.StatusBadRequest, "Bad Request", "email tidak valid")
		return
	}

	// Mendapatkan data pengguna berdasarkan email dari database
	existsDoc, err := helper.GetUserFromEmail(user.Email, db)
	if err != nil {
		// Jika terjadi kesalahan saat mendapatkan email dari database, kirim respons dengan status Internal Server Error dan pesan kesalahan
		helper.ErrorResponse(respw, req, http.StatusInternalServerError, "Internal Server Error", "kesalahan server : get email "+err.Error())
		return
	}

	// Mendekode salt dari string heksadesimal menjadi slice byte
	salt, err := hex.DecodeString(existsDoc.Salt)
	if err != nil {
		// Jika terjadi kesalahan saat mendekode salt, kirim respons dengan status Internal Server Error dan pesan kesalahan
		helper.ErrorResponse(respw, req, http.StatusInternalServerError, "Internal Server Error", "kesalahan server : salt")
		return
	}

	// Menghasilkan hash dari password yang diberikan oleh pengguna menggunakan algoritma Argon2
	hash := argon2.IDKey([]byte(user.Password), salt, 1, 64*1024, 4, 32)

	// Membandingkan hasil hash dengan hash yang tersimpan di database
	if hex.EncodeToString(hash) != existsDoc.Password {
		// Jika password tidak cocok, kirim respons dengan status Unauthorized dan pesan kesalahan
		helper.ErrorResponse(respw, req, http.StatusUnauthorized, "Unauthorized", "password salah")
		return
	}

	// Jika autentikasi berhasil, menghasilkan token menggunakan helper.Encode
	tokenstring, err := helper.Encode(user.ID, user.Email, privatekey)
	if err != nil {
		// Jika terjadi kesalahan saat menghasilkan token, kirim respons dengan status Internal Server Error dan pesan kesalahan
		helper.ErrorResponse(respw, req, http.StatusInternalServerError, "Internal Server Error", "kesalahan server : token")
		return
	}

	// Menyusun respons sukses dengan token yang dihasilkan dan data pengguna
	resp := map[string]interface{}{
		"status":  "success",
		"message": "login berhasil",
		"token":   tokenstring,
		"data": map[string]string{
			"email":       existsDoc.Email,
			"namalengkap": existsDoc.NamaLengkap,
		},
	}

	// Mengirimkan respons dalam format JSON dengan status OK
	helper.WriteJSON(respw, http.StatusOK, resp)
}
