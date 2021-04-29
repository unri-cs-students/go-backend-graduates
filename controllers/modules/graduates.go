package modules

import (
	"github.com/SemmiDev/go-backend/controllers"
	"github.com/SemmiDev/go-backend/entities"
	"github.com/SemmiDev/go-backend/repositories"
	"github.com/SemmiDev/go-backend/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GraduatesModule struct {
	controller entities.GraduatesController
	usecase    entities.GraduatesUsecase
	repo       entities.GraduatesRepository
}

func NewGraduatesModule(db *gorm.DB, g *gin.Engine) GraduatesModule {
	graduatesRepository := repositories.NewGraduatesRepository(db)
	graduatesUsecase := services.NewGraduatesUsecase(graduatesRepository)
	viewRepo := repositories.NewViewRepository(db)
	viewUsecase := services.NewViewUsecase(viewRepo)
	graduatesController := controllers.NewGraduatesController(g, graduatesUsecase, viewUsecase)

	if db != nil {
		_ = db.AutoMigrate(&entities.Graduates{})
		_ = db.AutoMigrate(&entities.View{})
		if (!db.Migrator().HasConstraint(&entities.Graduates{}, "Major")) {
			_ = db.Migrator().CreateConstraint(&entities.Graduates{}, "Major")
		}
	}
	return GraduatesModule{
		controller: graduatesController,
		usecase:    graduatesUsecase,
		repo:       graduatesRepository,
	}
}

func ResetGraduates(db *gorm.DB) {
	if db != nil {
		db.Migrator().DropTable(&entities.Graduates{})
	}
}