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

// user
func LogIn(db *mongo.Database, respw http.ResponseWriter, req *http.Request, privatekey string) {
	var user model.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		helper.ErrorResponse(respw, req, http.StatusBadRequest, "Bad Request", "error parsing request body "+err.Error())
		return
	}
	if user.Email == "" || user.Password == "" {
		helper.ErrorResponse(respw, req, http.StatusBadRequest, "Bad Request", "mohon untuk melengkapi data")
		return
	}
	if err = checkmail.ValidateFormat(user.Email); err != nil {
		helper.ErrorResponse(respw, req, http.StatusBadRequest, "Bad Request", "email tidak valid")
		return
	}
	existsDoc, err := helper.GetUserFromEmail(user.Email, db)
	if err != nil {
		helper.ErrorResponse(respw, req, http.StatusInternalServerError, "Internal Server Error", "kesalahan server : get email "+err.Error())
		return
	}
	salt, err := hex.DecodeString(existsDoc.Salt)
	if err != nil {
		helper.ErrorResponse(respw, req, http.StatusInternalServerError, "Internal Server Error", "kesalahan server : salt")
		return
	}
	hash := argon2.IDKey([]byte(user.Password), salt, 1, 64*1024, 4, 32)
	if hex.EncodeToString(hash) != existsDoc.Password {
		helper.ErrorResponse(respw, req, http.StatusUnauthorized, "Unauthorized", "password salah")
		return
	}
	tokenstring, err := helper.Encode(user.ID, user.Email, privatekey)
	if err != nil {
		helper.ErrorResponse(respw, req, http.StatusInternalServerError, "Internal Server Error", "kesalahan server : token")
		return
	}
	resp := map[string]any{
		"status":  "success",
		"message": "login berhasil",
		"token":   tokenstring,
		"data": map[string]string{
			"email":       existsDoc.Email,
			"namalengkap": existsDoc.NamaLengkap,
		},
	}
	helper.WriteJSON(respw, http.StatusOK, resp)
}
