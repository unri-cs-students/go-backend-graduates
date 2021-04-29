package modules

import (
	"github.com/SemmiDev/go-backend/controllers"
	"github.com/SemmiDev/go-backend/entities"
	"github.com/SemmiDev/go-backend/repositories"
	"github.com/SemmiDev/go-backend/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MajorModule struct {
	controller 	entities.MajorController
	services    entities.MajorUseCase
	repo       	entities.MajorRepository
}

func NewMajorModule(db *gorm.DB, g *gin.Engine) MajorModule {
	jurusanRepository := repositories.NewMajorRepository(db)
	jurusanUsecase := services.NewMajorUsecase(jurusanRepository)
	jurusanController := controllers.NewMajorController(g, jurusanUsecase)

	if db != nil {
		_ = db.AutoMigrate(&entities.Major{})
	}

	return MajorModule{
		controller: 	jurusanController,
		services:    	jurusanUsecase,
		repo:       	jurusanRepository,
	}
}

func ResetMajor(db *gorm.DB) {
	if db != nil {
		db.Migrator().DropTable(&entities.Major{})
	}
}

