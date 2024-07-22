package routes

import (
	"net/http"

	"github.com/tee-am-ai/backend/config"
	controller "github.com/tee-am-ai/backend/controller"
	"github.com/tee-am-ai/backend/helper"
)

func URL(w http.ResponseWriter, r *http.Request) {
	if config.SetAccessControlHeaders(w, r) {
		return // If it's a preflight request, return early.
	}

	if config.ErrorMongoconn != nil {
		helper.ErrorResponse(w, r, http.StatusInternalServerError, "Internal Server Error", "kesalahan server: database, "+config.ErrorMongoconn.Error())
		return
	}

	switch r.Method {
	case "GET":
		if r.URL.Path == "/" {
			Home(w, r)
			return
		}
	case "POST":
		switch r.URL.Path {
		case "/signup":
			controller.SignUp(config.Mongoconn, "users", w, r)
			return
		case "/login":
			controller.LogIn(config.Mongoconn, w, r, config.GetEnv("PASETOPRIVATEKEY"))
			return
		case "/chat":
			controller.Chat(w, r, config.GetEnv("TOKENMODEL"))
			return
		}
	}

	helper.ErrorResponse(w, r, http.StatusNotFound, "Not Found", "The requested resource was not found")
}


func Home(respw http.ResponseWriter, req *http.Request) {
	resp := map[string]string{
		"github_repo": "https://github.com/tee-am-ai/backend",
		"message": "Ampun puh sepuh, aku mah masih pemula 🙏",
	}
	helper.WriteJSON(respw, http.StatusOK, resp)
}
