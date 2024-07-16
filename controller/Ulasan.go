package controller

import (
	"encoding/json"
	"net/http"

	"github.com/tee-am-ai/backend/helper"
	model "github.com/tee-am-ai/backend/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Fungsi untuk menambahkan ulasan
func AddUlasan(db *mongo.Database, col string, respw http.ResponseWriter, req *http.Request) {
	var ulasan model.Ulasan

		// Decode request body menjadi struct Ulasan
		err := json.NewDecoder(req.Body).Decode(&ulasan)
		if err != nil {
			helper.ErrorResponse(respw, req, http.StatusBadRequest, "Bad Request", "error parsing request body "+err.Error())
			return
		}
}