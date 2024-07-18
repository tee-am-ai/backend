package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tee-am-ai/backend/routes"
)

func main() {

	r := routes.Router()

	port := ":8080"
	fmt.Println("Server started at: http://localhost" + port)

	err = http.ListenAndServe(port, r)
	if err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
