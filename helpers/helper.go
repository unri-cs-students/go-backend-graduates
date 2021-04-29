package helpers

import (
	serializer "github.com/SemmiDev/go-backend/commons/responses"
	"github.com/gin-gonic/gin"
)

func ForceResponse(ctx *gin.Context, status int, message string) {
	ctx.AbortWithStatusJSON(status,
		serializer.ResponseBase{
			Code:    status,
			Message: message,
		},
	)
}