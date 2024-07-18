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
	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/signup", signUp).Methods("POST")
	r.HandleFunc("/login", logIn).Methods("POST")
	r.HandleFunc("/chat", chat).Methods("POST")
	return r
}

func permission(w http.ResponseWriter, r *http.Request) {
	if config.SetAccessControlHeaders(w, r) {
		return // If it's a preflight request, return early.
	}

	if config.ErrorMongoconn != nil {
		helper.ErrorResponse(w, r, http.StatusInternalServerError, "Internal Server Error", "kesalahan server : database, " + config.ErrorMongoconn.Error())
		return
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	resp := map[string]string{
		"github_repo": "https://github.com/tee-am-ai/backend",
		"message":     "Ampun puh sepuh, aku mah masih pemula 🙏",
	}
	helper.WriteJSON(w, http.StatusOK, resp)
}

func signUp(w http.ResponseWriter, r *http.Request) {
	controller.SignUp(config.Mongoconn, "users", w, r)
}

func logIn(w http.ResponseWriter, r *http.Request) {
	controller.LogIn(config.Mongoconn, w, r, config.GetEnv("PASETOPRIVATEKEY"))
}

func chat(w http.ResponseWriter, r *http.Request) {
	controller.Chat(w, r, config.GetEnv("HUGGINGFACE_API_KEY"))
}