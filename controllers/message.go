package controllers

import (
	"errors"
	serializer "github.com/SemmiDev/go-backend/commons/responses"
	"github.com/SemmiDev/go-backend/commons/statuscode"
	"github.com/SemmiDev/go-backend/commons/validator"
	"github.com/SemmiDev/go-backend/controllers/middlewares"
	"github.com/SemmiDev/go-backend/entities"
	"github.com/SemmiDev/go-backend/helpers"
	"github.com/SemmiDev/go-backend/services"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	limit "github.com/yangxikun/gin-limit-by-key"
	"golang.org/x/time/rate"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type messageController struct {
	usecase entities.MessageUsecase
}

func NewMessageController(router *gin.Engine, mu entities.MessageUsecase) entities.MessageController {
	cont := messageController{usecase: mu}
	messageGroup := router.Group("/message")
	{
		messageGroup.POST("/", limit.NewRateLimiter(func(c *gin.Context) string {
			return c.ClientIP() // limit rate by client ip
		}, func(c *gin.Context) (*rate.Limiter, time.Duration) {
			return rate.NewLimiter(rate.Every(5*time.Minute), 3), time.Hour * 12
		}, func(c *gin.Context) {
			helpers.ForceResponse(c, http.StatusTooManyRequests, "too_many_requests")
		}), cont.CreateMessage)
		messageGroup.DELETE("/:id", middlewares.Auth, cont.DeleteMessage)
		messageGroup.GET("/wisudawan/:id", cont.GetMessage)
	}
	return &cont
}

func (controller *messageController) CreateMessage(ctx *gin.Context) {
	var j entities.CreateMessageSerializer
	if err := ctx.ShouldBindJSON(&j); err != nil {
		// Error dari post
		helpers.ForceResponse(ctx, http.StatusBadRequest, statuscode.UncompatibleJSON.String())
		return
	}
	if err := validator.IsValid(j); err != nil {
		helpers.ForceResponse(ctx, http.StatusBadRequest, statuscode.UncompatibleJSON.String())
		return
	}

	if err := controller.usecase.CreateMessage(j); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.ForceResponse(ctx, http.StatusNotFound, statuscode.NotFound.String())
			return
		}
		helpers.ForceResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, serializer.RESPONSE_OK)
	return
}

func (controller *messageController) DeleteMessage(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		helpers.ForceResponse(ctx, http.StatusNotFound, statuscode.EmptyParam.String())
		return
	}

	idToUuid := uuid.FromStringOrNil(id)
	if uuid.Equal(idToUuid, uuid.Nil) {
		helpers.ForceResponse(ctx, http.StatusNotFound, statuscode.EmptyParam.String())
		return
	}

	if err := controller.usecase.DeleteMessage(idToUuid); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.ForceResponse(ctx, http.StatusNotFound, statuscode.NotFound.String())
			return
		}
		helpers.ForceResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, serializer.RESPONSE_OK)
	return
}

func (controller *messageController) GetMessage(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		helpers.ForceResponse(ctx, http.StatusNotFound, statuscode.EmptyParam.String())
		return
	}

	idToUuid := uuid.FromStringOrNil(id)
	if uuid.Equal(idToUuid, uuid.Nil) {
		helpers.ForceResponse(ctx, http.StatusNotFound, statuscode.EmptyParam.String())
		return
	}

	result, err := controller.usecase.GetMessage(idToUuid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.ForceResponse(ctx, http.StatusNotFound, statuscode.NotFound.String())
			return
		}
		helpers.ForceResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	var parsedResult []entities.GetMessageSerializer
	if len(result) == 0 {
		parsedResult = make([]entities.GetMessageSerializer, 0)
	} else {
		parsedResult = make([]entities.GetMessageSerializer, len(result))
		for i, x := range result {
			parsedResult[i] = services.ConvertEntityMessageToSerializer(x)
		}
	}

	ctx.JSON(http.StatusOK, serializer.ResponseData{
		ResponseBase: serializer.RESPONSE_OK,
		Data:         parsedResult})
	return
}