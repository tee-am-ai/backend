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







func Chat(c *gin.Context) {
	controller.Chat(c.Writer, c.Request, os.Getenv("TOKENMODEL"))
}
