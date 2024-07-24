package controller

import (
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

func History(db *mongo.Database, col string, respw http.ResponseWriter, req *http.Request) {
	
}