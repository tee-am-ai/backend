package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tee-am-ai/backend/config"
	"github.com/tee-am-ai/backend/controller"
	"github.com/tee-am-ai/backend/helper"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", Home).Methods("GET")
	return r
}

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
	controller.LogIn(config.Mongoconn, w, r, config.GetEnv("PASETOPRIVATEKEY"))
}

func Chat(w http.ResponseWriter, r *http.Request) {
	controller.Chat(w, r, config.GetEnv("HUGGINGFACE_API_KEY"))
}