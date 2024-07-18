package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tee-am-ai/backend/config"
	"github.com/tee-am-ai/backend/controller"
	"github.com/tee-am-ai/backend/handler"
	"github.com/tee-am-ai/backend/helper"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.Use(handler.CorsMiddleware)

	r.HandleFunc("/", handler.Home).Methods("GET")
	r.HandleFunc("/signup", handler.SignUp).Methods("POST")
	r.HandleFunc("/login", handler.LogIn).Methods("POST")
	r.HandleFunc("/chat", handler.Chat).Methods("POST")

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
