package handler

import (
	"net/http"
	"os"

	"github.com/tee-am-ai/backend/config"
	"github.com/tee-am-ai/backend/controller"
	"github.com/tee-am-ai/backend/helper"
)

func Home(w http.ResponseWriter, r *http.Request) {
	resp := map[string]string{
		"github_repo": "https://github.com/tee-am-ai/backend",
		"message":     "Ampun puh sepuh, aku mah masih pemula 🙏",
	}
	
	helper.WriteJSON(w, http.StatusOK, resp)
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	controller.SignUp(config.Mongoconn, "users", w, r)
}

func LogIn(w http.ResponseWriter, r *http.Request) {
	controller.LogIn(config.Mongoconn, w, r, "users", os.Getenv("GO_PASETO_PRIVATE_KEY"))
}

func Chat(w http.ResponseWriter, r *http.Request) {
	controller.Chat(w, r, os.Getenv("GO_TOKEN_MODEL"))
}
