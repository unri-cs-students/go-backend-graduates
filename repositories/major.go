package repositories

import (
	"github.com/SemmiDev/go-backend/entities"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type majorRepository struct {
	db *gorm.DB
}

func NewMajorRepository(db *gorm.DB) entities.MajorRepository {
	return &majorRepository{db: db}
}

func (repo *majorRepository) GetOne(id uuid.UUID) (jurusan entities.Major, err error) {
	if err := repo.db.First(&jurusan, "id = ?", id).Error; err != nil {
		return jurusan, err
	}
	return
}

func (repo *majorRepository) AddOne(jurusan, fakultas, fakultasShort, jurusanShort string) error {
	jurusans := entities.Major{Major: jurusan, Faculty: fakultas, FacultyShort: fakultasShort, MajorShort: jurusanShort}
	if err := repo.db.Create(&jurusans).Error; err != nil {
		return err
	}
	return nil
}

func (repo *majorRepository) UpdateOne(id uuid.UUID, jurusan, fakultas, fakultasShort, jurusanShort string) error {
	var target entities.Major
	majorUpdate := map[string]interface{}{}
	if jurusan != "" {
		majorUpdate["major"] = jurusan
	}
	if fakultas != "" {
		majorUpdate["faculty"] = fakultas
	}
	if jurusanShort != "" {
		majorUpdate["major_short"] = jurusanShort
	}
	if fakultasShort != "" {
		majorUpdate["faculty_short"] = fakultasShort
	}
	if err := repo.db.First(&entities.Major{}, "id = ?", id).Error; err != nil {
		return err
	}
	if err := repo.db.Model(&target).Where("id = ?", id.String()).Updates(majorUpdate).Error; err != nil {
		return err
	}
	return nil
}

func (repo *majorRepository) DeleteOne(id uuid.UUID) error {
	if err := repo.db.First(&entities.Major{}, "id = ?", id).Error; err != nil {
		return err
	}
	if err := repo.db.Where("id = ?", id.String()).Delete(&entities.Major{}).Error; err != nil {
		return err
	}
	return nil
}

func (repo *majorRepository) GetAll() ([]entities.Major, error) {
	var result []entities.Major
	if err := repo.db.Model(&entities.Major{}).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}