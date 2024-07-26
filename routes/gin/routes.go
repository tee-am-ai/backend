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
	r.POST("/chat/:id", hgin.Chat)
	r.POST("/chat", hgin.Chat)
	r.GET("/chat", hgin.GetChat)
	r.GET("/chat/:id", hgin.GetChat)
	r.DELETE("/chat/:id", hgin.DeleteChat)
	r.GET("/user", hgin.GetUser)
	r.POST("/ulasan", hgin.AddUlasan)
	r.GET("/ulasan", hgin.GetAllUlasan)

	return r
}
