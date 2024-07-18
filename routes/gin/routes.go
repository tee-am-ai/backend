package gin

import (
	"github.com/gin-gonic/gin"
	hgin "github.com/tee-am-ai/backend/handler/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(hgin.CorsMiddleware())

	r.GET("/", hgin.Home)
	r.POST("/signup", hgin.SignUp)
	r.POST("/login", hgin.LogIn)
	r.POST("/chat", hgin.Chat)

	return r
}
