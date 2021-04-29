package modules

import (
	"github.com/SemmiDev/go-backend/helpers"
	"github.com/gin-gonic/gin"
	limit "github.com/yangxikun/gin-limit-by-key"
	"golang.org/x/time/rate"
	"net/http"
	"os"
	"time"
)

func NewLimiterModule(g *gin.Engine) {
	g.Use(limit.NewRateLimiter(func(c *gin.Context) string {
		return c.ClientIP() // limit rate by client ip
	}, func(c *gin.Context) (*rate.Limiter, time.Duration) {
		return rate.NewLimiter(rate.Every(500*time.Millisecond), 4), time.Hour
	}, func(c *gin.Context) {
		if values := c.Request.Header.Get("Authorization"); len(values) > 0 {
			token := os.Getenv("AUTH_TOKEN")
			if values == token {
				c.Next()
				return
			}
		}
		helpers.ForceResponse(c, http.StatusTooManyRequests, "too_many_requests")
	}))
}
