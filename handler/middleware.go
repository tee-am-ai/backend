package handler

import (
	"net/http"

	"github.com/tee-am-ai/backend/config"
	"github.com/tee-am-ai/backend/helper"
)

func CorsMiddleware(next http.Handler) http.Handler {
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
