package gin

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/tee-am-ai/backend/config"
	"github.com/tee-am-ai/backend/controller"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.GET("/", Home)
	r.POST("/signup", SignUp)
	r.POST("/login", LogIn)
	r.POST("/chat", Chat)

	return r
}

func Home(c *gin.Context) {
	c.JSON(200, gin.H{
		"github_repo": "https://github.com/tee-am-ai/backend",
		"message":     "Ampun puh sepuh, aku mah masih pemula 🙏",
	})
}

func SignUp(c *gin.Context) {
	controller.SignUp(config.Mongoconn, "users", c.Writer, c.Request)
}

