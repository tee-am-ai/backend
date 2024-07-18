package controller

import (
	"crypto/rand"   // Package untuk penggunaan fungsi random cryptographic
	"encoding/hex"  // Package untuk encoding dan decoding dalam format hexadecimal
	"encoding/json" // Package untuk encoding dan decoding dalam format JSON
	"net/http"      // Package untuk melakukan operasi HTTP
	"strings"       // Package untuk operasi manipulasi string

	"github.com/badoux/checkmail"              // Package untuk validasi alamat email
	"github.com/tee-am-ai/backend/helper"      // Package yang mungkin berisi fungsi bantuan (helper functions)
	model "github.com/tee-am-ai/backend/model" // Package yang mungkin berisi definisi model data
	"go.mongodb.org/mongo-driver/bson"         // Package untuk encoding dan decoding data dalam format BSON
	"go.mongodb.org/mongo-driver/mongo"        // Package untuk melakukan operasi terkait MongoDB
	"golang.org/x/crypto/argon2"               // Package untuk mengimplementasikan algoritma argon2 hashing
)

// SignUp adalah fungsi untuk menangani permintaan pendaftaran pengguna baru.
// Fungsi ini memerlukan akses ke database MongoDB (db *mongo.Database),
// nama koleksi MongoDB (col string), menanggapi permintaan HTTP (respw http.ResponseWriter, req *http.Request).
func SignUp(db *mongo.Database, col string, respw http.ResponseWriter, req *http.Request) {
	// Deklarasi variabel untuk menyimpan data pengguna yang dikirimkan dalam body permintaan HTTP
	var user model.User

	// Menguraikan dan memasukkan data JSON dari body permintaan ke dalam struktur model.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		// Jika terjadi kesalahan dalam parsing data, kirim respons dengan status Bad Request dan pesan kesalahan
		helper.ErrorResponse(respw, req, http.StatusBadRequest, "Bad Request", "error parsing request body "+err.Error())
		return
	}

	// Memastikan bahwa semua field yang diperlukan untuk pendaftaran diisi dengan benar
	if user.NamaLengkap == "" || user.Email == "" || user.Password == "" || user.Confirmpassword == "" {
		helper.ErrorResponse(respw, req, http.StatusBadRequest, "Bad Request", "mohon untuk melengkapi data")
		return
	}

	// Validasi format email menggunakan package checkmail
	if err := checkmail.ValidateFormat(user.Email); err != nil {
		helper.ErrorResponse(respw, req, http.StatusBadRequest, "Bad Request", "email tidak valid")
		return
	}

	// Memeriksa apakah email sudah terdaftar di dalam database
	userExists, _ := helper.GetUserFromEmail(user.Email, db)
	if user.Email == userExists.Email {
		helper.ErrorResponse(respw, req, http.StatusBadRequest, "Bad Request", "email sudah terdaftar")
		return
	}

	// Memeriksa apakah password mengandung spasi
	if strings.Contains(user.Password, " ") {
		helper.ErrorResponse(respw, req, http.StatusBadRequest, "Bad Request", "password tidak boleh mengandung spasi")
		return
	}

	// Memeriksa panjang password minimal 8 karakter
	if len(user.Password) < 8 {
		helper.ErrorResponse(respw, req, http.StatusBadRequest, "Bad Request", "password minimal 8 karakter")
		return
	}

	// Membuat salt acak untuk digunakan dalam proses hashing password
	salt := make([]byte, 16)
	_, err = rand.Read(salt)
	if err != nil {
		helper.ErrorResponse(respw, req, http.StatusInternalServerError, "Internal Server Error", "kesalahan server : salt")
		return
	}

	// Menggunakan argon2 untuk menghasilkan hash dari password yang diberikan menggunakan salt yang dibuat
	hashedPassword := argon2.IDKey([]byte(user.Password), salt, 1, 64*1024, 4, 32)

	// Menyiapkan data pengguna yang akan disimpan di dalam database MongoDB
	userData := bson.M{
		"namalengkap": user.NamaLengkap,
		"email":       user.Email,
		"password":    hex.EncodeToString(hashedPassword),
		"salt":        hex.EncodeToString(salt),
	}

	// Memasukkan data pengguna ke dalam database menggunakan fungsi bantuan InsertOneDoc
	insertedID, err := helper.InsertOneDoc(db, col, userData)
	if err != nil {
		helper.ErrorResponse(respw, req, http.StatusInternalServerError, "Internal Server Error", "kesalahan server : insert data, "+err.Error())
		return
	}

	// Persiapan respons yang akan dikirimkan kembali kepada klien setelah pendaftaran berhasil
	resp := map[string]interface{}{
		"message":    "berhasil mendaftar",
		"insertedID": insertedID,
		"data": map[string]string{
			"email": user.Email,
		},
	}

	// Mengirimkan respons dalam format JSON dengan status Created (201)
	helper.WriteJSON(respw, http.StatusCreated, resp)
}
