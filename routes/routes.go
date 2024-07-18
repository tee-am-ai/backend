package routes

import (
	"github.com/gorilla/mux"
	"github.com/tee-am-ai/backend/handler"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.Use(handler.CorsMiddleware)

	r.HandleFunc("/", handler.Home).Methods("GET")
	r.HandleFunc("/signup", SignUp).Methods("POST")
	r.HandleFunc("/login", LogIn).Methods("POST")
	r.HandleFunc("/chat", Chat).Methods("POST")

	return r
}