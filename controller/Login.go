package controller

import (
	"encoding/hex"
	"encoding/json"
	"net/http"

	"github.com/badoux/checkmail"
	"github.com/tee-am-ai/backend/helper"
	"github.com/tee-am-ai/backend/model"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/argon2"
)

func LogIn(db *mongo.Database, w http.ResponseWriter, r *http.Request, col, privatekey string) {
	var user model.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		helper.ErrorResponse(w, r, http.StatusBadRequest, "Bad Request", "error parsing request body " + err.Error())
		return
	}

	if user.Email == "" || user.Password == "" {
		helper.ErrorResponse(w, r, http.StatusBadRequest, "Bad Request", "mohon untuk melengkapi data")
		return
	}

	if err = checkmail.ValidateFormat(user.Email); err != nil {
		helper.ErrorResponse(w, r, http.StatusBadRequest, "Bad Request", "email tidak valid")
		return
	}

	existsDoc, err := helper.GetUserFromEmail(db, col, user.Email)
	if err != nil {
		helper.ErrorResponse(w, r, http.StatusInternalServerError, "Internal Server Error", "kesalahan server : get email " + err.Error())
		return
	}

	salt, err := hex.DecodeString(existsDoc.Salt)
	if err != nil {
		helper.ErrorResponse(w, r, http.StatusInternalServerError, "Internal Server Error", "kesalahan server : salt")
		return
	}

	hash := argon2.IDKey([]byte(user.Password), salt, 1, 64*1024, 4, 32)
	if hex.EncodeToString(hash) != existsDoc.Password {
		helper.ErrorResponse(w, r, http.StatusUnauthorized, "Unauthorized", "password salah")
		return
	}

	tokenstring, err := helper.Encode(user.ID, user.Email, privatekey)
	if err != nil {
		helper.ErrorResponse(w, r, http.StatusInternalServerError, "Internal Server Error", "kesalahan server : token")
		return
	}

	resp := map[string]any{
		"status":  "success",
		"message": "login berhasil",
		"token":   tokenstring,
		"data" : map[string]string{
			"email": existsDoc.Email,
			"namalengkap": existsDoc.NamaLengkap,
		},
	}

	helper.WriteJSON(w, http.StatusOK, resp)
}