package main

import (
	"net/http"

	"github.com/tee-am-ai/backend/routes"
)

func main() {
	http.HandleFunc("/", routes.URL)
	http.ListenAndServe(":8080", nil)
}