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

func SignUp(db *mongo.Database, col string, respw http.ResponseWriter, req *http.Request) {
	var user model.User

	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		helper.ErrorResponse(respw, req, http.StatusBadRequest, "Bad Request", "error parsing request body "+err.Error())
		return
	}

	if user.NamaLengkap == "" || user.Email == "" || user.Password == "" || user.Confirmpassword == "" {
		helper.ErrorResponse(respw, req, http.StatusBadRequest, "Bad Request", "mohon untuk melengkapi data")
		return
	}
	if err := checkmail.ValidateFormat(user.Email); err != nil {
		helper.ErrorResponse(respw, req, http.StatusBadRequest, "Bad Request", "email tidak valid")
		return
	}
	userExists, _ := helper.GetUserFromEmail(user.Email, db)
	if user.Email == userExists.Email {
		helper.ErrorResponse(respw, req, http.StatusBadRequest, "Bad Request", "email sudah terdaftar")
		return
	}
	if strings.Contains(user.Password, " ") {
		helper.ErrorResponse(respw, req, http.StatusBadRequest, "Bad Request", "password tidak boleh mengandung spasi")
		return
	}
	if len(user.Password) < 8 {
		helper.ErrorResponse(respw, req, http.StatusBadRequest, "Bad Request", "password minimal 8 karakter")
		return
	}
	salt := make([]byte, 16)
	_, err = rand.Read(salt)
	if err != nil {
		helper.ErrorResponse(respw, req, http.StatusInternalServerError, "Internal Server Error", "kesalahan server : salt")
		return
	}
	hashedPassword := argon2.IDKey([]byte(user.Password), salt, 1, 64*1024, 4, 32)
	userData := bson.M{
		"namalengkap": user.NamaLengkap,
		"email":       user.Email,
		"password":    hex.EncodeToString(hashedPassword),
		"salt":        hex.EncodeToString(salt),
	}
	insertedID, err := helper.InsertOneDoc(db, col, userData)
	if err != nil {
		helper.ErrorResponse(respw, req, http.StatusInternalServerError, "Internal Server Error", "kesalahan server : insert data, "+err.Error())
		return
	}
	resp := map[string]any{
		"message":    "berhasil mendaftar",
		"insertedID": insertedID,
		"data": map[string]string{
			"email": user.Email,
		},
	}
	helper.WriteJSON(respw, http.StatusCreated, resp)
}
