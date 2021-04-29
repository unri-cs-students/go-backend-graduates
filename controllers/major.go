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
	"gorm.io/gorm"
	"log"
	"net/http"
)

type majorController struct {
	usecase entities.MajorUseCase
}

func NewMajorController(router *gin.Engine, ju entities.MajorUseCase) entities.MajorController {
	cont := majorController{usecase: ju}
	majorGroup := router.Group("/major")
	{
		majorGroup.POST("/", middlewares.Auth, cont.CreateMajor)
		majorGroup.PUT("/", middlewares.Auth, cont.UpdateMajor)
		majorGroup.DELETE("/:id", middlewares.Auth, cont.DeleteMajor)
		majorGroup.GET("/:id", cont.GetMajor)
		majorGroup.GET("/", cont.GetAllMajor)
	}
	return &cont
}

func (controller majorController) CreateMajor(ctx *gin.Context) {
	var j entities.CreateMajorSerializer
	if err := ctx.ShouldBindJSON(&j); err != nil {
		helpers.ForceResponse(ctx, http.StatusBadRequest, statuscode.UncompatibleJSON.String())
		return
	}
	if err := validator.IsValid(j); err != nil {
		helpers.ForceResponse(ctx, http.StatusBadRequest, statuscode.UncompatibleJSON.String())
		return
	}

	if err := controller.usecase.CreateMajor(j); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.ForceResponse(ctx, http.StatusNotFound, statuscode.NotFound.String())
			return
		}
		helpers.ForceResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, serializer.RESPONSE_OK)
}

func (controller majorController) UpdateMajor(ctx *gin.Context) {
	var j entities.UpdateMajorSerializer
	if err := ctx.ShouldBindJSON(&j); err != nil {
		helpers.ForceResponse(ctx, http.StatusBadRequest, statuscode.UncompatibleJSON.String())
		return
	}

	if err := controller.usecase.UpdateMajor(j); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.ForceResponse(ctx, http.StatusNotFound, statuscode.NotFound.String())
			return
		}
		helpers.ForceResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, serializer.RESPONSE_OK)
}

func (controller majorController) DeleteMajor(ctx *gin.Context) {
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

	if err := controller.usecase.DeleteMajor(idToUuid); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.ForceResponse(ctx, http.StatusNotFound, statuscode.NotFound.String())
			return
		}
		helpers.ForceResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, serializer.RESPONSE_OK)
}

func (controller majorController) GetMajor(ctx *gin.Context) {
	id := ctx.Param("id")
	log.Println(id)
	if id == "" {
		helpers.ForceResponse(ctx, http.StatusNotFound, statuscode.EmptyParam.String())
		return
	}

	idToUuid := uuid.FromStringOrNil(id)
	if uuid.Equal(idToUuid, uuid.Nil) {
		helpers.ForceResponse(ctx, http.StatusNotFound, statuscode.EmptyParam.String())
		return
	}

	result, err := controller.usecase.GetMajor(idToUuid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.ForceResponse(ctx, http.StatusNotFound, statuscode.NotFound.String())
			return
		}
		helpers.ForceResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	parsedResult := services.ConvertEntityMajorToSerializer(result)
	ctx.JSON(http.StatusOK,
		serializer.ResponseData{
			ResponseBase: serializer.RESPONSE_OK,
			Data:         parsedResult,
		},
	)
}

func (controller majorController) GetAllMajor(ctx *gin.Context) {
	result, err := controller.usecase.GetAllMajor()
	if err != nil {
		helpers.ForceResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	var parsedResult []entities.GetMajorSerializer
	if len(result) == 0 {
		parsedResult = make([]entities.GetMajorSerializer, 0)
	} else {
		parsedResult = make([]entities.GetMajorSerializer, len(result))
		for i, x := range result {
			parsedResult[i] = services.ConvertEntityMajorToSerializer(x)
		}
	}
	ctx.JSON(http.StatusOK,
		serializer.ResponseData{
			ResponseBase: serializer.RESPONSE_OK,
			Data:         parsedResult,
		},
	)
}