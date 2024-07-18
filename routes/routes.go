package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tee-am-ai/backend/config"
	"github.com/tee-am-ai/backend/handler"
	"github.com/tee-am-ai/backend/helper"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.Use(corsMiddleware)

	r.HandleFunc("/", handler.Home).Methods("GET")
	r.HandleFunc("/signup", handler.SignUp).Methods("POST")
	r.HandleFunc("/login", handler.LogIn).Methods("POST")
	r.HandleFunc("/chat", handler.Chat).Methods("POST")

	return r
}