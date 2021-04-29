package middlewares

import (
	serializer "github.com/SemmiDev/go-backend/commons/responses"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func Auth(c *gin.Context) {
	if values := c.Request.Header.Get("Authorization"); len(values) > 0 {
		// Cek token
		token := os.Getenv("AUTH_TOKEN")
		if values == token {
			c.Next()
			return
		}

	}

	c.AbortWithStatusJSON(http.StatusForbidden, serializer.RESPONSE_FORBIDDEN)
}

func ResetAuth(c *gin.Context) {
	if values := c.Request.Header.Get("Authorization"); len(values) > 0 {
		// Cek token
		token := os.Getenv("RESET_TOKEN")
		if values == token {
			c.Next()
			return
		}

	}

	c.AbortWithStatusJSON(http.StatusForbidden, serializer.RESPONSE_FORBIDDEN)
}
