package middlewares

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORS(inDev bool) gin.HandlerFunc {
	if inDev {
		return cors.New(cors.Config{
			AllowHeaders: []string{"*"},
			//AllowCredentials: true,
			//MaxAge:           12 * time.Hour,
			AllowAllOrigins: true,
		})
	}
	return func(c *gin.Context) {

	}
}
