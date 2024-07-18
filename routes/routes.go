package routes

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/tee-am-ai/backend/config"
	controller "github.com/tee-am-ai/backend/controller"
	"github.com/tee-am-ai/backend/helper"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", Home).Methods("GET")
	r.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		controller.SignUp(config.Mongoconn, "users", w, r)
	}).Methods("POST")
	r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		controller.LogIn(config.Mongoconn, w, r, config.GetEnv("PASETOPRIVATEKEY"))
	}).Methods("POST")
	r.HandleFunc("/chat", func(w http.ResponseWriter, r *http.Request) {
		controller.Chat(w, r, config.GetEnv("TOKENMODEL"))
	}).Methods("POST")

	return r
}

// func URL(w http.ResponseWriter, r *http.Request) {
// 	if config.SetAccessControlHeaders(w, r) {
// 		return // If it's a preflight request, return early.
// 	}

// 	if config.ErrorMongoconn != nil {
// 		helper.ErrorResponse(w, r, http.StatusInternalServerError, "Internal Server Error", "kesalahan server : database, " + config.ErrorMongoconn.Error())
// 		return
// 	}

// 	var method, path string = r.Method, r.URL.Path
// 	switch {
// 	case method == "GET" && path == "/":
// 		Home(w, r)
// 	case method == "POST" && path == "/signup":
// 		controller.SignUp(config.Mongoconn, "users", w, r)
// 	case method == "POST" && path == "/login":
// 		controller.LogIn(config.Mongoconn, w, r, config.GetEnv("PASETOPRIVATEKEY"))
// 	case method == "POST" && path == "/chat":
// 		controller.Chat(w, r, config.GetEnv("TOKENMODEL"))
// 	default:
// 		helper.ErrorResponse(w, r, http.StatusNotFound, "Not Found", "The requested resource was not found")
// 	}
// }

func Home(respw http.ResponseWriter, req *http.Request) {
	resp := map[string]string{
		"github_repo": "https://github.com/tee-am-ai/backend",
		"message":     "Ampun puh sepuh, aku mah masih pemula 🙏",
	}
	helper.WriteJSON(respw, http.StatusOK, resp)
}
