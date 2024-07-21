package controller

import (
	"net/http"

	model "github.com/tee-am-ai/backend/model"
	"go.mongodb.org/mongo-driver/mongo"
)

// Fungsi untuk menambahkan ulasan
func AddUlasan(db *mongo.Database, col string, respw http.ResponseWriter, req *http.Request) {
	var ulasan model.Ulasan

	
	
}

