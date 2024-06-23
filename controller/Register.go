package controller

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/badoux/checkmail"
	"github.com/tee-am-ai/backend/helper"
	model "github.com/tee-am-ai/backend/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/argon2"
)


func SignUp(db *mongo.Database, col string, respw http.ResponseWriter, req *http.Request) {
	var user model.User

	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		resp := map[string]string{
			"error":   "Bad Request",
			"message": "error parsing application/json: " + err.Error(),
		}
		helper.WriteJSON(respw, http.StatusBadRequest, resp)
		return
	}

	if user.NamaLengkap == "" || user.Email == "" || user.Password == "" || user.Confirmpassword == ""{
		resp := map[string]string{
			"error":   "Bad Request",
			"message": "mohon untuk melengkapi data",
		}
		helper.WriteJSON(respw, http.StatusBadRequest, resp)
		return
	}
	if err := checkmail.ValidateFormat(user.Email); err != nil {
		resp := map[string]string{
			"error":   "Bad Request",
			"message": "email tidak valid",
		}
		helper.WriteJSON(respw, http.StatusBadRequest, resp)
		return 
	}
	userExists, _ := helper.GetUserFromEmail(user.Email, db)
	if user.Email == userExists.Email {
		resp := map[string]string{
			"error":   "Bad Request",
			"message": "email sudah terdaftar",
		}
		helper.WriteJSON(respw, http.StatusBadRequest, resp)
		return
	}
	if strings.Contains(user.Password, " ") {
		resp := map[string]string{
			"error":   "Bad Request",
			"message": "password tidak boleh mengandung spasi",
		}
		helper.WriteJSON(respw, http.StatusBadRequest, resp)
		return
	}
	if len(user.Password) < 8 {
		resp := map[string]string{
			"error":   "Bad Request",
			"message": "password minimal 8 karakter",
		}
		helper.WriteJSON(respw, http.StatusBadRequest, resp)
		return
	}
	salt := make([]byte, 16)
	_, err = rand.Read(salt)
	if err != nil {
		resp := map[string]string{
			"error":   "Internal Server Error",
			"message": "kesalahan server : salt",
		}
		helper.WriteJSON(respw, http.StatusInternalServerError, resp)
		return
	}
	hashedPassword := argon2.IDKey([]byte(user.Password), salt, 1, 64*1024, 4, 32)
	userData := bson.M{
		"namalengkap": user.NamaLengkap,
		"email": user.Email,
		"password": hex.EncodeToString(hashedPassword),
		"conformasiPassword": user.Confirmpassword,
		"salt": hex.EncodeToString(salt),
	}
	insertedID, err := helper.InsertOneDoc(db, col, userData)
	if err != nil {
		resp := map[string]string{
			"error":   "Internal Server Error",
			"message": "kesalahan server : insert",
		}
		helper.WriteJSON(respw, http.StatusInternalServerError, resp)
		return
	}
	resp := map[string]any{
		"message": "berhasil mendaftar",
		"insertedID": insertedID,
		"data" : map[string]string{
			"email": user.Email,
		},
	}
	helper.WriteJSON(respw, http.StatusCreated, resp)
}

