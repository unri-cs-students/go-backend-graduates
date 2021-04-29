package repositories

import (
	"github.com/SemmiDev/go-backend/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type contentRepository struct {
	db *gorm.DB
}

func NewContentRepository(db *gorm.DB) entities.ContentRepository {
	return &contentRepository{db: db}
}

func (repo *contentRepository) GetOne(id string) (result entities.Content, err error) {
	if err := repo.db.Model(&entities.Content{}).Where("id = ?", id).First(&result).Error; err != nil {
		return entities.Content{}, err
	}
	return
}

func (repo *contentRepository) GetByGraduates(identifierGraduates uint32) (result []entities.Content, err error) {
	var graduates entities.Graduates
	if err := repo.db.Model(&entities.Graduates{}).Find(&graduates, "identifier = ?", identifierGraduates).Error; err != nil {
		return nil, err
	}

	if err := repo.db.Preload(clause.Associations).
		Where("graduates_id = ?", graduates.ID).
		Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (repo *contentRepository) AddOne(identifierGraduates uint32, idOrgz, contenttype, headings, details, image string) error {
	var graduates entities.Graduates
	if err := repo.db.Model(&entities.Graduates{}).Where("identifier = ?", identifierGraduates).First(&graduates).Error; err != nil {
		return err
	}

	contentEntity := entities.Content{
		GraduatesID:    graduates.ID,
		OrganizationID: idOrgz,
		Type:           contenttype,
		Headings:       headings,
		Details:        details,
		Image:          image,
	}
	if err := repo.db.Create(&contentEntity).Error; err != nil {
		return err
	}
	return nil
}

func (repo *contentRepository) UpdateOne(idContent string, identifierGraduates uint32, idOrgz, contenttype, headings, details, image string) error {
	var content entities.Content
	contentUpdate := map[string]interface{}{}
	if identifierGraduates != 0 {
		var graduates entities.Graduates
		if err := repo.db.Model(&entities.Graduates{}).Find(&graduates, "identifier = ?", identifierGraduates).Error; err != nil {
			return err
		}
		contentUpdate["graduates_id"] = graduates.ID
	}

	if idOrgz != "" {
		contentUpdate["organization_id"] = idOrgz
	}

	if contenttype != "" {
		contentUpdate["content_type"] = contenttype
	}
	if headings != "" {
		contentUpdate["headings"] = headings
	}
	if details != "" {
		contentUpdate["details"] = details
	}
	if image != "" {
		contentUpdate["image"] = image
	}

	if err := repo.db.Model(&content).Where("id = ?", idContent).Updates(contentUpdate).Error; err != nil {
		return err
	}
	return nil
}

func (repo *contentRepository) DeleteOne(idContent string) error {
	if err := repo.db.First(&entities.Content{}, "id = ?", idContent).Error; err != nil {
		return err
	}
	if err := repo.db.Where("id = ?", idContent).Delete(&entities.Content{}).Error; err != nil {
		return err
	}
	return nil
}