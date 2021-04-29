package controllers

import (
	"errors"
	"github.com/SemmiDev/go-backend/commons/orgztype"
	serializer "github.com/SemmiDev/go-backend/commons/responses"
	"github.com/SemmiDev/go-backend/commons/statuscode"
	"github.com/SemmiDev/go-backend/commons/validator"
	"github.com/SemmiDev/go-backend/controllers/middlewares"
	"github.com/SemmiDev/go-backend/entities"
	"github.com/SemmiDev/go-backend/helpers"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"net/http"
)

type orgzController struct {
	usecase entities.OrgzUseCase
}

func NewOrgzController(router *gin.Engine, ou entities.OrgzUseCase) entities.OrgzController {
	cont := orgzController{usecase: ou}
	orgzGroup := router.Group("/orgz")
	{
		orgzGroup.POST("/", middlewares.Auth, cont.CreateOrgz)
		orgzGroup.PUT("/", middlewares.Auth, cont.UpdateOrgz)
		orgzGroup.DELETE("/:id", middlewares.Auth, cont.DeleteOrgz)
		orgzGroup.GET("/id/:id", cont.GetByID) //TODO ganti jadi
		orgzGroup.GET("/slug/:slug", cont.GetBySlug)
		orgzGroup.GET("/all", cont.GetAll)
	}
	return &cont
}


func (controller *orgzController) CreateOrgz(ctx *gin.Context) {
	var j entities.CreateOrgzSerializer
	if err := ctx.ShouldBindJSON(&j); err != nil {
		helpers.ForceResponse(ctx, http.StatusBadRequest, statuscode.UncompatibleJSON.String())
		return
	}
	if err := validator.IsValid(j); err != nil {
		helpers.ForceResponse(ctx, http.StatusBadRequest, statuscode.UncompatibleJSON.String())
		return
	}
	enum, enumErr := orgztype.GetEnum(j.Category)
	if enumErr != nil {
		helpers.ForceResponse(ctx, http.StatusBadRequest, statuscode.UnknownType.String())
		return
	}
	j.Category = enum
	if err := controller.usecase.CreateOrgz(j); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.ForceResponse(ctx, http.StatusNotFound, statuscode.NotFound.String())
			return
		}

		helpers.ForceResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, serializer.RESPONSE_OK)
}

func (controller *orgzController) UpdateOrgz(ctx *gin.Context) {
	var j entities.UpdateOrgzSerializer
	if err := ctx.ShouldBindJSON(&j); err != nil {
		helpers.ForceResponse(ctx, http.StatusBadRequest, statuscode.UncompatibleJSON.String())
		return
	}
	if j.Category != "" {
		enum, enumErr := orgztype.GetEnum(j.Category)
		if enumErr != nil {
			helpers.ForceResponse(ctx, http.StatusBadRequest, statuscode.UnknownType.String())
			return
		}
		j.Category = enum
	}
	if err := controller.usecase.UpdateOrgz(j); err != nil {
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

func (controller *orgzController) DeleteOrgz(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		helpers.ForceResponse(ctx, http.StatusNotFound, statuscode.EmptyParam.String())
		return
	}

	idToUuid := uuid.FromStringOrNil(id)
	if uuid.Equal(idToUuid, uuid.Nil) {
		helpers.ForceResponse(ctx, http.StatusBadRequest, statuscode.UnknownUUID.String())
		return
	}

	if err := controller.usecase.DeleteOrgz(idToUuid); err != nil {
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

func (controller *orgzController) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		helpers.ForceResponse(ctx, http.StatusNotFound, statuscode.EmptyParam.String())
		return
	}

	idToUuid := uuid.FromStringOrNil(id)
	if uuid.Equal(idToUuid, uuid.Nil) {
		helpers.ForceResponse(ctx, http.StatusBadRequest, statuscode.UnknownUUID.String())
		return
	}

	result, err := controller.usecase.GetOrgz(idToUuid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.ForceResponse(ctx, http.StatusNotFound, statuscode.NotFound.String())
			return
		}
		helpers.ForceResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	if result.Logo == "" {
		result.Logo = "Logo/default-orgz.png"
	}

	ctx.JSON(http.StatusOK, serializer.ResponseData{
		ResponseBase: serializer.RESPONSE_OK,
		Data:         result})
	return
}

func (controller *orgzController) GetBySlug(ctx *gin.Context) {
	slug := ctx.Param("slug")
	if slug == "" {
		helpers.ForceResponse(ctx, http.StatusNotFound, statuscode.EmptyParam.String())
		return
	}

	result, err := controller.usecase.GetBySlug(slug)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.ForceResponse(ctx, http.StatusNotFound, statuscode.NotFound.String())
			return
		}
		helpers.ForceResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	if result.Logo == "" {
		result.Logo = "Logo/default-orgz.png"
	}

	ctx.JSON(http.StatusOK, serializer.ResponseData{
		ResponseBase: serializer.RESPONSE_OK,
		Data:         result})
	return
}

func (controller *orgzController) GetAll(ctx *gin.Context) {
	result, err := controller.usecase.GetAll()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.ForceResponse(ctx, http.StatusNotFound, statuscode.NotFound.String())
			return
		}
		helpers.ForceResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	if len(result) == 0 {
		result = make([]entities.Orgz, 0)
	} else {
		for i := range result {
			if result[i].Logo == "" {
				result[i].Logo = "Logo/default-orgz.png"
			}
		}
	}
	ctx.JSON(http.StatusOK, serializer.ResponseData{
		ResponseBase: serializer.RESPONSE_OK,
		Data:         result})
	return
}
