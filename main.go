package main

import (
	"fmt"
	"net/http"

	"github.com/tee-am-ai/backend/routes"
)

func main() {
	r := routes.Router()
	port := ":8080"
	fmt.Println("Server started at: http://localhost" + port)
	http.ListenAndServe(port, r)
}
