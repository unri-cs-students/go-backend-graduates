package modules

import (
	"github.com/SemmiDev/go-backend/controllers"
	"github.com/SemmiDev/go-backend/entities"
	"github.com/SemmiDev/go-backend/repositories"
	"github.com/SemmiDev/go-backend/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ContentModule struct {
	controllers entities.ContentController
	services    entities.ContentUseCase
	repo       entities.ContentRepository
}

func NewContentModule(db *gorm.DB, g *gin.Engine) ContentModule {
	contentRepository := repositories.NewContentRepository(db)
	contentUsecase := services.NewContentUsecase(contentRepository)
	contentController := controllers.NewContentController(g, contentUsecase)

	if db != nil {
		db.AutoMigrate(&entities.Content{})
		if (!db.Migrator().HasConstraint(&entities.Content{}, "Wisudawan")) {
			db.Migrator().CreateConstraint(&entities.Content{}, "Wisudawan")
		}
	}

	return ContentModule{
		controllers: contentController,
		services:    contentUsecase,
		repo:       contentRepository,
	}
}

func ResetContent(db *gorm.DB) {
	if db != nil {
		_ = db.Migrator().DropTable(&entities.Content{})
	}
}
