package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tee-am-ai/backend/config"
	"github.com/tee-am-ai/backend/helper"
)

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if config.SetAccessControlHeaders(c.Writer, c.Request) {
			return
		}

		if config.ErrorMongoconn != nil {
			helper.ErrorResponse(c.Writer, c.Request, http.StatusInternalServerError, "Internal Server Error", "kesalahan server : database, "+config.ErrorMongoconn.Error())
			return
		}

		c.Next()
	}
}
