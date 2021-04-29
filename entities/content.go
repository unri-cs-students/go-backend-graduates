package entities


import (
	"github.com/SemmiDev/go-backend/commons/domain"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type Content struct {
	domain.EntityBase
	GraduatesID    string    `json:"id_graduates" gorm:"type:VARCHAR(50);not null"`
	Graduates      Graduates `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	OrganizationID string    `json:"id_organization" gorm:"type:VARCHAR(50);null"`
	Organization   Orgz      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Type           string    `gorm:"type:VARCHAR(16);not null" json:"content_type"`
	Headings       string    `gorm:"type:VARCHAR(255);not null" json:"headings"`
	Details        string    `gorm:"type:TEXT" json:"details"`
	Image          string    `gorm:"type:VARCHAR(255)" json:"image"`
}

type GetContentSerializer struct {
	ContentType string `json:"content_type"`
	Headings    string `json:"headings"`
	Details     string `json:"details"`
	Image       string `json:"image"`
}

type GetContentSerializer2 struct {
	GetContentSerializer
	OrganizationName string `json:"organization_name"`
	OrganizationLogo string `json:"organization_logo"`
}

type GetContentsSerializer struct {
	OrganizationalContents []GetContentSerializer2 `json:"org_data"`
	SelfContents           []GetContentSerializer  `json:"self_data"`
}

func (Content) TableName() string {
	return "content"
}

type CreateContentSerializer struct {
	Identifier   uint32 `json:"identifier" wispril:"required"`
	Organization string `json:"id_organization"`
	ContentType  string `json:"content_type" wispril:"required" binding:"lte=16"`
	Headings     string `json:"headings" wispril:"required" binding:"lte=255"`
	Details      string `json:"details"`
	Image        string `json:"image" binding:"lte=255"`
}

type UpdateContentSerializer struct {
	Content      string `json:"id_content" wispril:"required"`
	Identifier   uint32 `json:"identifier"`
	Organization string `json:"id_organization"`
	ContentType  string `json:"content_type" binding:"lte=16"`
	Headings     string `json:"headings" binding:"lte=255"`
	Details      string `json:"details"`
	Image        string `json:"image" binding:"lte=255"`
}

type ContentController interface {
	CreateContent(ctx *gin.Context)
	UpdateContent(ctx *gin.Context)
	DeleteContent(ctx *gin.Context)
	GetContent(ctx *gin.Context)
	GetContentByGraduates(ctx *gin.Context)
}

type ContentUseCase interface {
	CreateContent(item CreateContentSerializer) error
	DeleteContent(IdContent uuid.UUID) error
	UpdateContent(item UpdateContentSerializer) error
	GetContent(IdContent uuid.UUID) (Content, error)
	GetByGraduates(IdentifierGraduates uint32) ([]Content, error)
}

type ContentRepository interface {
	GetOne(id string) (Content, error)
	GetByGraduates(identifierGraduates uint32) ([]Content, error)
	AddOne(identifierGraduates uint32, idOrgz, contenttype, headings, details, image string) error
	UpdateOne(idContent string, identifierGraduates uint32, idOrgz, contenttype, headings, details, image string) error
	DeleteOne(idContent string) error
}