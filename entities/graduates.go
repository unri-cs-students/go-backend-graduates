package entities

import (
	"github.com/SemmiDev/go-backend/commons/domain"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Graduates struct {
	domain.EntityBase
	Identifier  	uint32    `json:"identifier" gorm:"unique"`
	Name        	string    `json:"name" gorm:"type:VARCHAR(255);not null"`
	NickName    	string    `json:"nick_name" gorm:"type:VARCHAR(255);not null"`
	ThesisTitle 	string    `json:"thesis_title" gorm:"type:VARCHAR(1024);not null"`
	Incoming    	uint16    `json:"incoming" gorm:"type:SMALLINT;not null"`
	MajorID    		string    `json:"major_id" gorm:"type:VARCHAR(50);not null"`
	Major      		Major     `json:"major" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Instagram   	string    `json:"instagram" gorm:"type:VARCHAR(255)"`
	Linkedin    	string    `json:"linkedin" gorm:"type:VARCHAR(255)"`
	Twitter     	string    `json:"twitter" gorm:"type:VARCHAR(255)"`
	PlaceOfBirth  	string    `json:"pob" gorm:"type:VARCHAR(255)"`
	DateOfBirth 	time.Time `json:"dob" gorm:"type:DATE"`
	Photo        	string    `json:"photo" gorm:"type:VARCHAR(255)"`
}

func (Graduates) TableName() string {
	return "graduates"
}

type GetGraduatesSerializer struct {
	ID            	string `json:"id_graduates"`
	Identifier    	uint32 `json:"identifier"`
	Name          	string `json:"name"`
	NickName     	string `json:"nick_name"`
	ThesisTitle     string `json:"thesis_title"`
	Incoming      	uint16 `json:"incoming"`
	Major       	string `json:"major"`
	MajorShort  	string `json:"major_short"`
	Faculty      	string `json:"faculty"`
	FacultyShort 	string `json:"faculty_short"`
	Instagram     	string `json:"instagram"`
	Linkedin      	string `json:"linkedin"`
	Twitter       	string `json:"twitter"`
	PlaceOfBirth   	string `json:"pob"`
	DateOfBirth  	string `json:"dob"`
	Photo         	string `json:"photo"`
}

type GetSimpleGraduatesSerializer struct {
	ID            	string `json:"id_graduates"`
	Identifier      uint32 `json:"identifier"`
	Name          	string `json:"name"`
	ThesisTitle     string `json:"thesis_title"`
	Major       	string `json:"major"`
	MajorShort  	string `json:"major_short"`
	Faculty      	string `json:"faculty"`
	FacultyShort 	string `json:"faculty_short"`
	Photo         	string `json:"photo"`
}

type CreateGraduatesSerializer struct {
	Identifier      uint32 `json:"identifier" wispril:"required"`
	Name         	string `json:"name" wispril:"required" binding:"lte=255"`
	NickName    	string `json:"nick_name" wispril:"required" binding:"lte=255"`
	ThesisTitle     string `json:"thesis_title" wispril:"required" binding:"lte=1024"`
	Incoming     	uint16 `json:"incoming" wispril:"required" binding:"lte=25"`
	Major      		string `json:"major_id" wispril:"required"`
	Instagram    	string `json:"instagram" binding:"lte=255"`
	Linkedin     	string `json:"linkedin" binding:"lte=255"`
	Twitter      	string `json:"twitter" binding:"lte=255"`
	PlaceOfBirth  	string `json:"pob" binding:"lte=255"`
	DateOfBirth 	string `json:"dob" binding:"lte=10"`
	Photo        	string `json:"photo" binding:"lte=255"`
}

type UpdateGraduatesSerializer struct {
	NIM          	uint32 `json:"identifier" wispril:"required"`
	Name         	string `json:"name" binding:"lte=255"`
	NickName    	string `json:"nick_name" binding:"lte=255"`
	ThesisTitle     string `json:"thesis_title" binding:"lte=1024"`
	Incoming     	uint16 `json:"incoming" binding:"lte=25"`
	Major      		string `json:"major_id"`
	Instagram    	string `json:"instagram" binding:"lte=255"`
	Linkedin     	string `json:"linkedin" binding:"lte=255"`
	Twitter      	string `json:"twitter" binding:"lte=255"`
	PlaceOfBirth  	string `json:"pob"  binding:"lte=255"`
	DateOfBirth 	string `json:"dob" binding:"lte=10"`
	Photo        	string `json:"photo"  binding:"lte=255"`
}

type GraduatesController interface {
	CreateGraduates(ctx *gin.Context)
	UpdateGraduates(ctx *gin.Context)
	DeleteGraduates(ctx *gin.Context)
	GetGraduates(ctx *gin.Context)
	FilterGraduatesByOrgzSlug(ctx *gin.Context)
	Trending(ctx *gin.Context)
}

type GraduatesUsecase interface {
	CreateGraduates(item CreateGraduatesSerializer) error
	DeleteGraduates(idGraduates uuid.UUID) error
	UpdateGraduates(item UpdateGraduatesSerializer) error
	GetGraduates(idGraduates uuid.UUID) (Graduates, error)
	GetAllGraduates() ([]Graduates, error)
	FilterGraduatesByOrgzSlug(organizationSlug string) ([]Graduates, error)
}

type GraduatesRepository interface {
	GetOne(graduatesID string) (Graduates, error)
	GetAll() ([]Graduates, error)
	AddOne(identifier uint32, incoming uint16, name, nickName, thesisTitle, major, instagram, linkedin, twitter, pob, photo string, dob time.Time) error
	UpdateOne(identifier uint32, incoming uint16, name, nickName, thesisTitle, majorID, instagram, linkedin, twitter, pob, photo string, dob time.Time) error
	DeleteOne(GraduatesID string) error
	FilterByOrgzSlug(organizationSlug string) ([]Graduates, error)
}