package middlewares

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func CORS(inDev bool) gin.HandlerFunc {
	if inDev {
		return cors.New(cors.Config{
			AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
			AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
			AllowCredentials: false,
			MaxAge:           12 * time.Hour,
			ExposeHeaders:    []string{"Content-Length"},
			AllowOriginFunc: func(origin string) bool {
				return true
			},
		})
	}
	return func(ctx *gin.Context) {

	}
}
