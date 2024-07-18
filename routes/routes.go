package routes

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/tee-am-ai/backend/config"
	"github.com/tee-am-ai/backend/controller"
	"github.com/tee-am-ai/backend/helper"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.Use(permission)

	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/signup", signUp).Methods("POST")
	r.HandleFunc("/login", logIn).Methods("POST")
	r.HandleFunc("/chat", chat).Methods("POST")

	return r
}

func permission(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if config.SetAccessControlHeaders(w, r) {
				return
			}

			if config.ErrorMongoconn != nil {
				helper.ErrorResponse(w, r, http.StatusInternalServerError, "Internal Server Error", "kesalahan server : database, "+config.ErrorMongoconn.Error())
				return
			}
			next.ServeHTTP(w, r)
		},
	)
}




func logIn(w http.ResponseWriter, r *http.Request) {
	controller.LogIn(config.Mongoconn, w, r, os.Getenv("PASETOPRIVATEKEY"))
}

func chat(w http.ResponseWriter, r *http.Request) {
	controller.Chat(w, r, os.Getenv("HUGGINGFACE_API_KEY"))
}
