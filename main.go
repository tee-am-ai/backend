package main

import (
	"github.com/tee-am-ai/backend/config"
	"github.com/tee-am-ai/backend/routes/gin"
)

func main() {
	config.LoadEnv()
	r := gin.Router()

	r.Run(":8080")
}
