package controller

import (
	"encoding/json"
	"net/http"

	"github.com/tee-am-ai/backend/helper"
	model "github.com/tee-am-ai/backend/model"
	"go.mongodb.org/mongo-driver/mongo"
)

// SignUp
func SignUp(db *mongo.Database, col string, respw http.ResponseWriter, req *http.Request) {
	var user model.User

	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		helper.ErrorResponse(respw, req, http.StatusBadRequest, "Bad Request", "error parsing request body "+err.Error())
		return
	}

}
