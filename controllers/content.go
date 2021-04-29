package controllers

import (
	"errors"
	contenttype "github.com/SemmiDev/go-backend/commons/content_type"
	serializer "github.com/SemmiDev/go-backend/commons/responses"
	"github.com/SemmiDev/go-backend/commons/statuscode"
	"github.com/SemmiDev/go-backend/commons/validator"
	"github.com/SemmiDev/go-backend/controllers/middlewares"
	"github.com/SemmiDev/go-backend/entities"
	"github.com/SemmiDev/go-backend/helpers"
	"github.com/SemmiDev/go-backend/services"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type contentController struct {
	usecase entities.ContentUseCase
}

func NewContentController(router *gin.Engine, cu entities.ContentUseCase) entities.ContentController {
	cont := contentController{usecase: cu}
	contentGroup := router.Group("/content")
	{
		contentGroup.POST("/", middlewares.Auth, cont.CreateContent)
		contentGroup.PUT("/", middlewares.Auth, cont.UpdateContent)
		contentGroup.DELETE("/:id", middlewares.Auth, cont.DeleteContent)
		contentGroup.GET("/id/:id", cont.GetContent)
		contentGroup.GET("/wisudawan/:id", cont.GetContentByGraduates)
	}
	return &cont
}

func (controller *contentController) CreateContent(ctx *gin.Context) {
	var j entities.CreateContentSerializer
	if err := ctx.ShouldBindJSON(&j); err != nil {
		helpers.ForceResponse(ctx, http.StatusBadRequest, statuscode.UncompatibleJSON.String())
		return
	}
	if err := validator.IsValid(j); err != nil {
		helpers.ForceResponse(ctx, http.StatusBadRequest, statuscode.UncompatibleJSON.String())
		return
	}
	enum, enumErr := contenttype.GetEnum(j.ContentType)
	if enumErr != nil {
		helpers.ForceResponse(ctx, http.StatusBadRequest, statuscode.UnknownType.String())
		return
	}
	j.ContentType = enum

	if err := controller.usecase.CreateContent(j); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.ForceResponse(ctx, http.StatusNotFound, statuscode.NotFound.String())
			return
		}
		helpers.ForceResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, serializer.RESPONSE_OK)
}

func (controller *contentController) UpdateContent(ctx *gin.Context) {
	var j entities.UpdateContentSerializer
	if err := ctx.ShouldBindJSON(&j); err != nil {
		helpers.ForceResponse(ctx, http.StatusBadRequest, statuscode.UncompatibleJSON.String())
		return
	}

	if j.ContentType != "" {
		enum, enumErr := contenttype.GetEnum(j.ContentType)
		if enumErr != nil {
			helpers.ForceResponse(ctx, http.StatusBadRequest, statuscode.UnknownType.String())
			return
		}
		j.ContentType = enum
	}

	if err := controller.usecase.UpdateContent(j); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.ForceResponse(ctx, http.StatusNotFound, statuscode.NotFound.String())
			return
		}
		helpers.ForceResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, serializer.RESPONSE_OK)
}

func (controller *contentController) DeleteContent(ctx *gin.Context) {
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

	if err := controller.usecase.DeleteContent(idToUuid); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.ForceResponse(ctx, http.StatusNotFound, statuscode.NotFound.String())
			return
		}
		helpers.ForceResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, serializer.RESPONSE_OK)
}

func (controller *contentController) GetContent(ctx *gin.Context) {
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

	result, err := controller.usecase.GetContent(idToUuid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.ForceResponse(ctx, http.StatusNotFound, statuscode.NotFound.String())
			return
		}
		helpers.ForceResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK,
		serializer.ResponseData{
			ResponseBase: serializer.RESPONSE_OK,
			Data:         result,
		},
	)
}

func (controller *contentController) GetContentByGraduates(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		helpers.ForceResponse(ctx, http.StatusNotFound, statuscode.EmptyParam.String())
		return
	}
	u, convertErr := strconv.ParseUint(id, 10, 32)
	if convertErr != nil {
		helpers.ForceResponse(ctx, http.StatusBadRequest, convertErr.Error())
		return
	}

	result, err := controller.usecase.GetByGraduates(uint32(u))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.ForceResponse(ctx, http.StatusNotFound, statuscode.NotFound.String())
			return
		}
		helpers.ForceResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	parsedResult := services.ConvertEntityContentsToSerializer(result)

	ctx.JSON(http.StatusOK,
		serializer.ResponseData{
			ResponseBase: serializer.RESPONSE_OK,
			Data:         parsedResult,
		},
	)
}