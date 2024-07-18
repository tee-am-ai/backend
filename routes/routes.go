package routes

import (
	"github.com/gorilla/mux"
	"github.com/tee-am-ai/backend/handler"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.Use(handler.CorsMiddleware)

}
