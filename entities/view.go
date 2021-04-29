package entities

import (
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type View struct {
	ID          string    `gorm:"primaryKey;type:VARCHAR(50)" json:"id"`
	GraduatesID string    `gorm:"type:VARCHAR(50);not null" json:"id_graduates"`
	IP          string    `gorm:"type:VARCHAR(255);not null" json:"ip"`
	AccessTime  time.Time `json:"time"`
}

type GetViewGraduates struct {
	Graduates Graduates `gorm:"embedded" json:"graduates"`
	Count     int64     `json:"count"`
}

func (e *View) BeforeCreate(scope *gorm.DB) error {
	e.ID = uuid.NewV4().String()
	return nil
}

func (View) TableName() string {
	return "view"
}

type ViewController interface {
	AddView(ctx *gin.Context)
}

type ViewUseCase interface {
	AddView(idGraduates uuid.UUID, clientIP string) error
	GetTop5() ([]GetViewGraduates, error)
}

type ViewRepository interface {
	AddOne(IdGraduates string, IP string, Time time.Time) error
	GetLast(IdGraduates string, IP string) (View, error)
	GetTop5() ([]GetViewGraduates, error)
}
