package modules

import (
	"github.com/SemmiDev/go-backend/controllers"
	"github.com/SemmiDev/go-backend/entities"
	"github.com/SemmiDev/go-backend/repositories"
	"github.com/SemmiDev/go-backend/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrgzModule struct {
	controller entities.OrgzController
	usecase    entities.OrgzUseCase
	repo       entities.OrgzRepository
}

func NewOrgzModule(db *gorm.DB, g *gin.Engine) OrgzModule {
	orgzRepository := repositories.NewOrgzRepository(db)
	orgzUsecase := services.NewOrgzUsecase(orgzRepository)
	orgzController := controllers.NewOrgzController(g, orgzUsecase)
	if db != nil {
		_ = db.AutoMigrate(&entities.Orgz{})
	}

	return OrgzModule{
		controller: orgzController,
		usecase:    orgzUsecase,
		repo:       orgzRepository,
	}
}

func ResetOrgz(db *gorm.DB) {
	if db != nil {
		_ = db.Migrator().DropTable(&entities.Orgz{})
	}
}
