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



func SignUp(c *gin.Context) {
	controller.SignUp(config.Mongoconn, "users", c.Writer, c.Request)
}

func LogIn(c *gin.Context) {
	controller.LogIn(config.Mongoconn, c.Writer, c.Request, os.Getenv("PASETOPRIVATEKEY"))
}

func Chat(c *gin.Context) {
	controller.Chat(c.Writer, c.Request, os.Getenv("TOKENMODEL"))
}
