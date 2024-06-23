package controller

import (
	"encoding/hex"
	"encoding/json"
	"net/http"

	"github.com/badoux/checkmail"
	"github.com/tee-am-ai/backend/helper"
	model "github.com/tee-am-ai/backend/model"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/argon2"
)

// user
func LogIn(db *mongo.Database, respw http.ResponseWriter, req *http.Request, privatekey string) {
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
	if user.Email == "" || user.Password == "" {
		resp := map[string]string{
			"error":   "Bad Request",
			"message": "mohon untuk melengkapi data",
		}
		helper.WriteJSON(respw, http.StatusBadRequest, resp)
		return
	}
	if err = checkmail.ValidateFormat(user.Email); err != nil {
		resp := map[string]string{
			"error":   "Bad Request",
			"message": "email tidak valid",
		}
		helper.WriteJSON(respw, http.StatusBadRequest, resp)
		return
	}
	existsDoc, err := helper.GetUserFromEmail(user.Email, db)
	if err != nil {
		return
	}
	salt, err := hex.DecodeString(existsDoc.Salt)
	if err != nil {
		resp := map[string]string{
			"error":   "Internal Server Error",
			"message": "kesalahan server : salt",
		}
		helper.WriteJSON(respw, http.StatusInternalServerError, resp)
		return
	}
	hash := argon2.IDKey([]byte(user.Password), salt, 1, 64*1024, 4, 32)
	if hex.EncodeToString(hash) != existsDoc.Password {
		resp := map[string]string{
			"error":   "Bad Request",
			"message": "password salah",
		}
		helper.WriteJSON(respw, http.StatusBadRequest, resp)
		return
	}
	tokenstring, err := helper.Encode(user.ID, user.Email, privatekey)
	if err != nil {
		resp := map[string]string{
			"error":   "Internal Server Error",
			"message": "kesalahan server : token",
		}
		helper.WriteJSON(respw, http.StatusInternalServerError, resp)
		return
	}
	resp := map[string]string{
		"status":  "success",
		"message": "login berhasil",
		"token":   tokenstring,
	}
	helper.WriteJSON(respw, http.StatusOK, resp)
}
