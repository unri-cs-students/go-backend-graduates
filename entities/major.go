package entities

import (
	"github.com/SemmiDev/go-backend/commons/domain"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type Major struct {
	domain.EntityBase
	Major       	string `gorm:"type:VARCHAR(100);not null" json:"major"`
	Faculty      	string `gorm:"type:VARCHAR(100);not null" json:"faculty"`
	FacultyShort 	string `gorm:"type:VARCHAR(5);not null" json:"faculty_short"`
	MajorShort  	string `gorm:"type:VARCHAR(5);not null" json:"major_short"`
}

type GetMajorSerializer struct {
	Id            	string `json:"id_major"`
	Major       	string `json:"major"`
	Faculty      	string `json:"faculty"`
	FacultyShort 	string `json:"faculty_short"`
	MajorShort  	string `json:"major_short"`
}

func (Major) TableName() string {
	return "major"
}

type CreateMajorSerializer struct {
	Major       	string `json:"major" wispril:"required"`
	Faculty      	string `json:"faculty" wispril:"required"`
	FacultyShort 	string `json:"faculty_short" wispril:"required" binding:"lte=5"`
	MajorShort  	string `json:"major_short" wispril:"required" binding:"lte=5"`
}

type UpdateMajorSerializer struct {
	IdMajor     	uuid.UUID `json:"id_major" wispril:"required"`
	Major       	string    `json:"major"`
	Faculty      	string    `json:"faculty"`
	FacultyShort 	string    `json:"faculty_short" binding:"lte=5"`
	MajorShort  	string    `json:"major_short" binding:"lte=5"`
}

type MajorController interface {
	CreateMajor(ctx *gin.Context)
	UpdateMajor(ctx *gin.Context)
	DeleteMajor(ctx *gin.Context)
	GetMajor(ctx *gin.Context)
	GetAllMajor(ctx *gin.Context)
}

type MajorUseCase interface {
	CreateMajor(item CreateMajorSerializer) error
	DeleteMajor(IdMajor uuid.UUID) error
	UpdateMajor(item UpdateMajorSerializer) error
	GetMajor(IdMajor uuid.UUID) (Major, error)
	GetAllMajor() ([]Major, error)
}

type MajorRepository interface {
	GetOne(id uuid.UUID) (Major, error)
	AddOne(major, faculty, facultyShort, majorShort string) error
	UpdateOne(id uuid.UUID, major, faculty, facultyShort, majorShort string) error
	DeleteOne(id uuid.UUID) error
	GetAll() ([]Major, error)
}