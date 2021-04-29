package modules

import (
	"github.com/SemmiDev/go-backend/controllers"
	"github.com/SemmiDev/go-backend/entities"
	"github.com/SemmiDev/go-backend/repositories"
	"github.com/SemmiDev/go-backend/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MessageModule struct {
	controller entities.MessageController
	usecase    entities.MessageUsecase
	repo       entities.MessageRepository
}

func NewMessageModule(db *gorm.DB, g *gin.Engine) MessageModule {
	messageRepository := repositories.NewMessageRepository(db)
	messageUsecase := services.NewMessageUsecase(messageRepository)
	messageController := controllers.NewMessageController(g, messageUsecase)
	if db != nil {
		_ = db.AutoMigrate(&entities.Message{})
		if (!db.Migrator().HasConstraint(&entities.Message{}, "Receiver")) {
			_ = db.Migrator().CreateConstraint(&entities.Message{}, "Receiver")
		}
	}

	return MessageModule{
		controller: messageController,
		usecase:    messageUsecase,
		repo:       messageRepository,
	}
}

func ResetMessage(db *gorm.DB) {
	if db != nil {
		db.Migrator().DropTable(&entities.Message{})
	}
}