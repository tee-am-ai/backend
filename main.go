package main

import (
	hgin "github.com/tee-am-ai/backend/routes/gin"
)

func main() {
	// config.LoadEnv()
	r := hgin.Router()

	r.Run(":8080")
}
