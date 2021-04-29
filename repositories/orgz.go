package repositories

import (
	"github.com/SemmiDev/go-backend/entities"
	"gorm.io/gorm"
)

type orgzRepository struct {
	db *gorm.DB
}

func NewOrgzRepository(db *gorm.DB) entities.OrgzRepository {
	return &orgzRepository{db: db}
}

func (repo *orgzRepository) GetOne(idOrgz string) (entities.Orgz, error) {
	var result entities.Orgz
	if err := repo.db.First(&result, "id = ?", idOrgz).Error; err != nil {
		return result, err
	}
	return result, nil
}

func (repo *orgzRepository) AddOne(name, slug, category, logo, apresiasiPoster, apresiasiTulisan, apresiasiVideo, fakultasShort string) error {
	create := entities.Orgz{
		Name:                name,
		Slug:                slug,
		Category:            category,
		Logo:                logo,
		PosterAppreciation:  apresiasiPoster,
		WritingAppreciation: apresiasiTulisan,
		VideoAppreciation:   apresiasiVideo,
		FacultyShort:        fakultasShort,
	}
	if err := repo.db.Create(&create).Error; err != nil {
		return err
	}
	return nil
}

func (repo *orgzRepository) UpdateOne(idOrgz, name, slug, category, logo, apresiasiPoster, apresiasiTulisan, apresiasiVideo, fakultasShort string) error {
	var target entities.Orgz
	update := map[string]interface{}{}
	if idOrgz != "" {
		update["id"] = idOrgz
	}
	if slug != "" {
		update["slug"] = slug
	}
	if name != "" {
		update["name"] = name
	}
	if category != "" {
		update["category"] = category
	}
	if logo != "" {
		update["logo"] = logo
	}
	if apresiasiPoster != "" {
		update["apresiasi_poster"] = apresiasiPoster
	}
	if apresiasiTulisan != "" {
		update["apresiasi_tulisan"] = apresiasiTulisan
	}
	if apresiasiVideo != "" {
		update["apresiasi_video"] = apresiasiVideo
	}
	if fakultasShort != "" {
		update["fakultas_short"] = fakultasShort
	}
	if err := repo.db.First(&entities.Orgz{}, "id = ?", idOrgz).Error; err != nil {
		return err
	}
	if err := repo.db.Model(&target).Where("id = ?", idOrgz).Updates(update).Error; err != nil {
		return err
	}
	return nil
}

func (repo *orgzRepository) DeleteOne(idOrgz string) error {
	if err := repo.db.First(&entities.Orgz{}, "id = ?", idOrgz).Error; err != nil {
		return err
	}
	if err := repo.db.Where("id = ?", idOrgz).Delete(&entities.Orgz{}).Error; err != nil {
		return err
	}
	return nil
}

func (repo *orgzRepository) GetAll() ([]entities.Orgz, error) {
	var results []entities.Orgz
	if err := repo.db.Find(&results).Error; err != nil {
		return make([]entities.Orgz, 0), err
	}
	return results, nil
}

func (repo *orgzRepository) GetBySlug(slug string) (entities.Orgz, error) {
	var result entities.Orgz
	if err := repo.db.First(&result, "slug = ?", slug).Error; err != nil {
		return entities.Orgz{}, err
	}
	return result, nil
}