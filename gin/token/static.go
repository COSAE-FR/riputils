package token

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

// StaticTokenMiddleware generates a gin middleware that check a token in headers
func StaticTokenMiddleware(token string, logger *log.Entry) gin.HandlerFunc {
	if logger == nil {
		logger = log.WithField("component", "auth_token")
	}
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		if auth == "" {
			logger.Trace("No Authorization header")
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		requestToken := strings.TrimPrefix(auth, "Bearer ")
		if requestToken == auth {
			logger.Trace("No Authorization bearer")
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		if requestToken == token {
			c.Next()
			return
		}
		c.AbortWithStatus(http.StatusForbidden)
	}
}
