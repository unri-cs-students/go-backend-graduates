package controllers

import (
	"errors"
	"fmt"
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

type graduatesController struct {
	usecase     entities.GraduatesUsecase
	viewUsecase entities.ViewUseCase
}

func NewGraduatesController(router *gin.Engine, wu entities.GraduatesUsecase, vu entities.ViewUseCase) entities.GraduatesController {
	cont := graduatesController{usecase: wu, viewUsecase: vu}

	graduatesGroup := router.Group("/graduates")
	{
		graduatesGroup.POST("/", middlewares.Auth, cont.CreateGraduates)
		graduatesGroup.PUT("/", middlewares.Auth, cont.UpdateGraduates)
		graduatesGroup.DELETE("/:id", middlewares.Auth, cont.DeleteGraduates)
		graduatesGroup.GET("/id/:id", cont.GetGraduates)
		graduatesGroup.GET("/org/:slug", cont.FilterGraduatesByOrgzSlug)
		graduatesGroup.GET("/trending", cont.Trending)
	}
	return &cont
}

func (controller *graduatesController) CreateGraduates(ctx *gin.Context) {
	var j entities.CreateGraduatesSerializer
	if err := ctx.ShouldBindJSON(&j); err != nil {
		log.Println(j)
		log.Println(err.Error())
		helpers.ForceResponse(ctx, http.StatusBadRequest, statuscode.UncompatibleJSON.String())
		return

	}
	if err := validator.IsValid(j); err != nil {
		helpers.ForceResponse(ctx, http.StatusBadRequest, statuscode.UncompatibleJSON.String())
		return
	}

	if err := controller.usecase.CreateGraduates(j); err != nil {
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

func (controller *graduatesController) UpdateGraduates(ctx *gin.Context) {
	var j entities.UpdateGraduatesSerializer
	if err := ctx.ShouldBindJSON(&j); err != nil {
		helpers.ForceResponse(ctx, http.StatusBadRequest, statuscode.UncompatibleJSON.String())
		return
	}

	if err := controller.usecase.UpdateGraduates(j); err != nil {
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

func (controller *graduatesController) DeleteGraduates(ctx *gin.Context) {
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

	if err := controller.usecase.DeleteGraduates(idToUuid); err != nil {
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

func (controller *graduatesController) GetGraduates(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		helpers.ForceResponse(ctx, http.StatusNotFound, statuscode.EmptyParam.String())
		return
	}

	idToUuid := uuid.FromStringOrNil(id)
	if uuid.Equal(idToUuid, uuid.Nil) {
		helpers.ForceResponse(ctx, http.StatusNotFound, statuscode.NotFound.String())
		return
	}

	result, err := controller.usecase.GetGraduates(idToUuid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.ForceResponse(ctx, http.StatusNotFound, statuscode.NotFound.String())
			return
		}
		helpers.ForceResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	parsedResult := services.ConvertEntityGraduatesToSerializer(result)
	viewErr := controller.viewUsecase.AddView(idToUuid, ctx.ClientIP())
	if viewErr != nil {
		fmt.Println(viewErr.Error())
	}

	ctx.JSON(http.StatusOK, serializer.ResponseData{
		ResponseBase: serializer.RESPONSE_OK,
		Data:         parsedResult})
	return
}

func (controller *graduatesController) FilterGraduatesByOrgzSlug(ctx *gin.Context) {
	slug := ctx.Param("slug")
	if slug == "" {
		helpers.ForceResponse(ctx, http.StatusNotFound, statuscode.EmptyParam.String())
		return
	}

	result, err := controller.usecase.FilterGraduatesByOrgzSlug(slug)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.ForceResponse(ctx, http.StatusNotFound, statuscode.NotFound.String())
			return
		}
		helpers.ForceResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	var parsedResult []entities.GetSimpleGraduatesSerializer
	if len(result) == 0 {
		parsedResult = make([]entities.GetSimpleGraduatesSerializer, 0)
	} else {
		parsedResult = make([]entities.GetSimpleGraduatesSerializer, len(result))
		for i, x := range result {
			parsedResult[i] = services.ConvertEntityGraduatesToSimpleSerializer(x)
		}
	}

	ctx.JSON(http.StatusOK, serializer.ResponseData{
		ResponseBase: serializer.RESPONSE_OK,
		Data:         parsedResult})
	return
}

func (controller *graduatesController) Trending(ctx *gin.Context) {
	result, err := controller.viewUsecase.GetTop5()
	if err != nil {
		helpers.ForceResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	var parsedResult []struct {
		Graduates entities.GetSimpleGraduatesSerializer
		Count     int64
	}
	if len(result) != 0 {
		parsedResult = make([]struct {
			Graduates entities.GetSimpleGraduatesSerializer
			Count     int64
		}, len(result))
		for i := range result {
			parsedResult[i] = struct {
				Graduates entities.GetSimpleGraduatesSerializer
				Count     int64
			}{Count: result[i].Count,
				Graduates: services.ConvertEntityGraduatesToSimpleSerializer(result[i].Graduates)}
		}
	} else {
		parsedResult = make([]struct {
			Graduates entities.GetSimpleGraduatesSerializer
			Count     int64
		}, 0)
	}

	ctx.JSON(http.StatusOK, serializer.ResponseData{
		ResponseBase: serializer.RESPONSE_OK,
		Data:         parsedResult})
	return
}