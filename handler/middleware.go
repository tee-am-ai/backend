package handler

import (
	"net/http"

	"github.com/tee-am-ai/backend/config"
)

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if config.SetAccessControlHeaders(w, r) {
				return
			}


			next.ServeHTTP(w, r)
		},
	)
}
