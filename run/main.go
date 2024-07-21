package main

import (
	"fmt"
	"net/http"

	"github.com/tee-am-ai/backend/routes"
)

func main() {
	http.HandleFunc("/", routes.URL)
	port := ":8080"
	fmt.Println("Server started at: http://localhost" + port)
	http.ListenAndServe(port, nil)
}

func ChatPredictions(w http.ResponseWriter, r *http.Request) {
}
