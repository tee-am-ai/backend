package controller

import (
	"net/http"

	"github.com/tee-am-ai/backend/helper"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUser(db *mongo.Database, w http.ResponseWriter, r *http.Request,  col, publickey string) {
	token := r.Header.Get("Authorization")
	if token == "" {
		helper.ErrorResponse(w, r, http.StatusUnauthorized, "Unauthorized", "token tidak ditemukan")
		return
	}

	userInfo, err := helper.Decode(publickey, token)
	if err != nil {
		helper.ErrorResponse(w, r, http.StatusBadRequest, "Bad Request", "token tidak valid")
		return
	}

	user, err := helper.GetUserFromEmail(db, col, userInfo.Email)
	if err != nil {
		helper.ErrorResponse(w, r, http.StatusBadRequest, "Bad Request", "user tidak ditemukan")
		return
	}

	resp := map[string]any{
		"status":  "success",
		"message": "get user berhasil",
		"data": map[string]string{
			"email":       user.Email,
			"namalengkap": user.NamaLengkap,
		},
	}

	helper.WriteJSON(w, http.StatusOK, resp)
}