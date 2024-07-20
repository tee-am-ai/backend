package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tee-am-ai/backend/config"
	"github.com/tee-am-ai/backend/routes"
)

func main() {
	config.LoadEnv()
	r := routes.Router()

	port := ":8080"
	local := "http://localhost" + port
	fmt.Println("Server started at: ", local)

	err := http.ListenAndServe(port, r)
	if err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
