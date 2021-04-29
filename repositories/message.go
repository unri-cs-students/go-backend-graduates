package repositories

import (
	"github.com/SemmiDev/go-backend/entities"
	"gorm.io/gorm"
)

type messageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) entities.MessageRepository {
	return &messageRepository{db: db}
}

func (repo *messageRepository) AddOne(message, sender, graduatesID string) error {
	messageEntity := entities.Message{
		ReceiverID: graduatesID,
		Message:    message,
		Sender:     sender,
	}
	if err := repo.db.Create(&messageEntity).Error; err != nil {
		return err
	}
	return nil
}

func (repo *messageRepository) DeleteOne(idMessage string) error {
	if err := repo.db.First(&entities.Message{}, "id = ?", idMessage).Error; err != nil {
		return err
	}
	if err := repo.db.Where("id = ?", idMessage).Delete(&entities.Message{}).Error; err != nil {
		return err
	}
	return nil
}

func (repo *messageRepository) GetMessage(graduatesID string) ([]entities.Message, error) {
	var results []entities.Message
	if err := repo.db.Find(&results, "receiver_id = ?", graduatesID).Error; err != nil {
		return make([]entities.Message, 0), err
	}
	return results, nil
}