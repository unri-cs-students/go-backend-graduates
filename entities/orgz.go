package entities

import (
	"github.com/SemmiDev/go-backend/commons/domain"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type Orgz struct {
	domain.EntityBase
	Name             	string `gorm:"type:VARCHAR(255);not null" json:"name"`
	Slug             	string `gorm:"type:VARCHAR(128);not null;unique" json:"slug"`
	Category         	string `gorm:"type:VARCHAR(64);not null" json:"category"`
	Logo             	string `gorm:"type:VARCHAR(255);not null" json:"logo"`
	PosterAppreciation  string `gorm:"type:VARCHAR(255);" json:"poster_appreciation"`
	WritingAppreciation string `gorm:"type:text;" json:"writing_appreciation"`
	VideoAppreciation   string `gorm:"type:VARCHAR(255);" json:"video_appreciation"`
	FacultyShort    	string `gorm:"type:VARCHAR(5)" json:"faculty_appreciation"`
}

func (Orgz) TableName() string {
	return "organization"
}

type CreateOrgzSerializer struct {
	Name             	string `json:"name" wispril:"required" binding:"lte=255"`
	Slug             	string `json:"slug" wispril:"required" binding:"lte=255"`
	Category         	string `json:"category" wispril:"required" binding:"lte=64"`
	Logo             	string `json:"logo" binding:"lte=255"`
	PosterAppreciation  string `json:"poster_appreciation" binding:"lte=255"`
	WritingAppreciation string `json:"writing_appreciation"`
	VideoAppreciation   string `json:"video_appreciation" binding:"lte=255"`
	FacultyShort    	string `json:"faculty_appreciation" binding:"lte=5"`
}

type UpdateOrgzSerializer struct {
	IdOrgz           	string `json:"id_organization" wispril:"required"`
	Slug             	string `json:"slug" binding:"lte=255"`
	Name             	string `json:"name" binding:"lte=255"`
	Category         	string `json:"category" binding:"lte=64"`
	Logo             	string `json:"logo" binding:"lte=255"`
	PosterAppreciation  string `json:"poster_appreciation" binding:"lte=255"`
	WritingAppreciation string `json:"writing_appreciation"`
	VideoAppreciation   string `json:"video_appreciation" binding:"lte=255"`
	FacultyShort    	string `json:"faculty_appreciation" binding:"lte=5"`
}

type OrgzController interface {
	CreateOrgz(ctx *gin.Context)
	UpdateOrgz(ctx *gin.Context)
	DeleteOrgz(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	GetBySlug(ctx *gin.Context)
	GetAll(ctx *gin.Context)
}

type OrgzUseCase interface {
	CreateOrgz(item CreateOrgzSerializer) error
	DeleteOrgz(idOrgz uuid.UUID) error
	UpdateOrgz(item UpdateOrgzSerializer) error
	GetOrgz(idOrgz uuid.UUID) (Orgz, error)
	GetAll() ([]Orgz, error)
	GetBySlug(slug string) (Orgz, error)
}

type OrgzRepository interface {
	GetOne(idOrgz string) (Orgz, error)
	AddOne(name, slug, category, logo, posterAppreciation, writingAppreciation, videoAppreciation, facultyAppreciation string) error
	UpdateOne(idOrgz, name, slug, category, logo, posterAppreciation, writingAppreciation, videoAppreciation, facultyAppreciation string) error
	DeleteOne(idOrgz string) error
	GetAll() ([]Orgz, error)
	GetBySlug(slug string) (Orgz, error)
}