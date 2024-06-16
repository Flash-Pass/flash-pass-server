package middlewares

import (
	"github.com/Flash-Pass/flash-pass-server/internal/auth"
	"github.com/Flash-Pass/flash-pass-server/internal/constants"
	"github.com/Flash-Pass/flash-pass-server/internal/fpstatus"
	"github.com/Flash-Pass/flash-pass-server/internal/res"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

var (
	UrlWhiteList = []string{
		"login", "register",
	}
)

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		// All get requests are allowed
		//if c.Request.Method == http.MethodGet {
		//	c.Next()
		//	return
		//}

		for _, url := range UrlWhiteList {
			if strings.Contains(c.Request.URL.Path, url) {
				c.Next()
				return
			}
		}

		token := c.GetHeader("Authorization")

		claim, err := auth.ParseToken(c, token)
		if err != nil {
			res.RespondWithError(c, http.StatusUnauthorized, fpstatus.ParseTokenError, nil)
		}

		c.Set(constants.CtxUserIdKey, claim.Id)
		c.Set(constants.CtxOpenIdKey, claim.OpenId)
		c.Set(constants.CtxSessionKeyKey, claim.SessionKey)
		c.Set(constants.CtxUnionIdKey, claim.UnionId)

		c.Next()
	}
}
