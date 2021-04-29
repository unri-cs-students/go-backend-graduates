package entities

import (
	"github.com/SemmiDev/go-backend/commons/domain"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type Message struct {
	domain.EntityBase
	ReceiverID string    `gorm:"type:VARCHAR(50);not null" json:"id_graduates"`
	Receiver   Graduates `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Message    string    `gorm:"type:TEXT;not null" json:"message"`
	Sender     string    `gorm:"type:VARCHAR(255);not null" json:"sender"`
}
type GetMessageSerializer struct {
	ID      string `json:"id_message"`
	Sent    string `json:"sent"`
	Sender  string `json:"sender"`
	Message string `json:"message"`
}

func (Message) TableName() string {
	return "message"
}

type CreateMessageSerializer struct {
	IdGraduates string `json:"id_graduates" wispril:"required"`
	Message     string `json:"message" wispril:"required"`
	Sender      string `json:"sender" wispril:"required" binding:"lte=255"`
}

type MessageController interface {
	CreateMessage(ctx *gin.Context)
	DeleteMessage(ctx *gin.Context)
	GetMessage(ctx *gin.Context)
}

type MessageUsecase interface {
	CreateMessage(item CreateMessageSerializer) error
	DeleteMessage(idMessage uuid.UUID) error
	GetMessage(idGraduates uuid.UUID) ([]Message, error)
}

type MessageRepository interface {
	AddOne(message, sender, idGraduates string) error
	DeleteOne(idMessage string) error
	GetMessage(idGraduates string) ([]Message, error)
}