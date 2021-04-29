package middlewares

import (
	serializer "github.com/SemmiDev/go-backend/commons/responses"
	"github.com/ekyoung/gin-nice-recovery"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitErrorHandler(g *gin.Engine) {
	g.Use(nice.Recovery(CustomRecovery))
}

func CustomRecovery(c *gin.Context, recovered interface{}) {
	if err, ok := recovered.(string); ok {
		c.JSON(http.StatusBadRequest, serializer.ResponseBase{
			Code:    http.StatusBadRequest,
			Message: err,
		})
	}

	c.AbortWithStatus(http.StatusBadRequest)
}
