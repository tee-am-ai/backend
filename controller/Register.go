package controller

import (
	"net/http"

	model "github.com/tee-am-ai/backend/model"
	"go.mongodb.org/mongo-driver/mongo"
)

// SignUp
func SignUp(db *mongo.Database, col string, respw http.ResponseWriter, req *http.Request) {
	var user model.User

}
