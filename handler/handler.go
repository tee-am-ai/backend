package handler

import (
	"net/http"

	"github.com/tee-am-ai/backend/helper"
)

func Home(w http.ResponseWriter, r *http.Request) {
	resp := map[string]string{
		"github_repo": "https://github.com/tee-am-ai/backend",
		"message":     "Ampun puh sepuh, aku mah masih pemula 🙏",
	}
	helper.WriteJSON(w, http.StatusOK, resp)
}
