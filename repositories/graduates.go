package repositories

import (
	"github.com/SemmiDev/go-backend/entities"
	"log"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

)

type graduatesRepo struct {
	db *gorm.DB
}

func NewGraduatesRepository(db *gorm.DB) entities.GraduatesRepository {
	return &graduatesRepo{db: db}
}

func (repo *graduatesRepo) GetOne(graduatedID string) (result entities.Graduates, err error) {
	if err = repo.db.Preload(clause.Associations).Where("id = ?", graduatedID).First(&result).Error; err != nil {
		return entities.Graduates{}, err
	}

	return
}

func (repo *graduatesRepo) GetAll() (results []entities.Graduates, err error) {
	if err = repo.db.Find(&results).Error; err != nil {
		return make([]entities.Graduates, 0), err
	}

	return
}

func (repo *graduatesRepo) AddOne(
	identifier uint32, incoming uint16, name, nickName, thesisTitle, majorId, instagram, linkedin, twitter, pob, photo string, dob time.Time) error {
	graduates := entities.Graduates{
		Identifier:     identifier,
		Incoming:    	incoming,
		Name:        	name,
		NickName:   	nickName,
		ThesisTitle:    thesisTitle,
		Instagram:   	instagram,
		Linkedin:    	linkedin,
		Twitter:     	twitter,
		PlaceOfBirth: 	pob,
		Photo:       	photo,
		MajorID:   		majorId,
	}

	if !dob.IsZero() {
		graduates.DateOfBirth = dob
	}

	if err := repo.db.Create(&graduates).Error; err != nil {
		return err
	}

	return nil
}

func (repo *graduatesRepo) UpdateOne(identifier uint32, incoming uint16, name, nickName, thesisTitle, majorId, instagram, linkedin, twitter, pob, photo string, dob time.Time) error {
	var graduates entities.Graduates
	graduates_update := map[string]interface{}{}

	if incoming != 0 {
		graduates_update["incoming"] = incoming
	}
	if name != "" {
		graduates_update["name"] = name
	}
	if nickName != "" {
		graduates_update["nick_name"] = nickName
	}
	if thesisTitle != "" {
		graduates_update["thesis_title"] = thesisTitle
	}
	if instagram != "" {
		graduates_update["instagram"] = instagram
	}
	if linkedin != "" {
		graduates_update["linkedin"] = linkedin
	}
	if twitter != "" {
		graduates_update["twitter"] = twitter
	}
	if pob != "" {
		graduates_update["place_of_birth"] = pob
	}
	if photo != "" {
		graduates_update["photo"] = photo
	}
	if !dob.IsZero() {
		graduates_update["date_of_birth"] = dob
	}
	if majorId != "" {
		graduates_update["major_id"] = majorId
	}
	if err := repo.db.First(&entities.Graduates{}, "identifier = ?", identifier).Error; err != nil {
		log.Println(err.Error())
		return err
	}
	if err := repo.db.Model(&graduates).Where("identifier = ?", identifier).Updates(graduates_update).Error; err != nil {
		return err
	}
	return nil
}

func (repo *graduatesRepo) DeleteOne(GraduatesID string) (err error) {
	if err = repo.db.First(&entities.Graduates{}, "id = ?", GraduatesID).Error; err != nil {
		return err
	}
	if err = repo.db.Where("id = ?", GraduatesID).Delete(&entities.Graduates{}).Error; err != nil {
		return err
	}
	return nil
}

func (repo *graduatesRepo) FilterByOrgzSlug(organizationSlug string) (result []entities.Graduates, err error) {

	resultGraduatesID := repo.db.Table("graduates").
		Joins("INNER JOIN content ON content.graduates_id = graduates.id").
		Joins("INNER JOIN organization ON organization.id = content.organization_id").
		Where("slug = ?", organizationSlug).
		Distinct("graduates.id")

	if err = resultGraduatesID.Error; err != nil {
		return nil, err
	}

	if err = repo.db.Preload(clause.Associations).
		Find(&result, "id IN (?)", resultGraduatesID).
		Error; err != nil {
		return nil, err
	}
	return result, nil
}