package gin

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/tee-am-ai/backend/config"
	"github.com/tee-am-ai/backend/controller"
)

func Home(c *gin.Context) {
	c.JSON(200, gin.H{
		"github_repo": "https://github.com/tee-am-ai/backend",
		"message":     "Ampun puh sepuh, aku mah masih pemula 🙏",
	})
}

func SignUp(c *gin.Context) {
	controller.SignUp(config.Mongoconn, "users", c.Writer, c.Request)
}

func LogIn(c *gin.Context) {
	controller.LogIn(config.Mongoconn, c.Writer, c.Request, os.Getenv("PASETOPRIVATEKEY"))
}

func Chat(c *gin.Context) {
	controller.Chat(c.Writer, c.Request, os.Getenv("TOKENMODEL"))
}