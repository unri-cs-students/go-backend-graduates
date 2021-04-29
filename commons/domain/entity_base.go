package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type EntityBase struct {
	ID        string    `gorm:"primaryKey;type:VARCHAR(50)" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (e *EntityBase) BeforeCreate(scope *gorm.DB) error {
	e.ID = uuid.NewV4().String()
	return nil
}