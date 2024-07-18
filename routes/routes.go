package routes

import (
	"net/http"

	"github.com/tee-am-ai/backend/config"
	controller "github.com/tee-am-ai/backend/controller"
	"github.com/tee-am-ai/backend/helper"
)


func Home(respw http.ResponseWriter, req *http.Request) {
	resp := map[string]string{
		"github_repo": "https://github.com/tee-am-ai/backend",
		"message": "Ampun puh sepuh, aku mah masih pemula 🙏",
	}
	helper.WriteJSON(respw, http.StatusOK, resp)
}
