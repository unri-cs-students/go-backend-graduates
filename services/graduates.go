package services

import (
	"github.com/SemmiDev/go-backend/entities"
	uuid "github.com/satori/go.uuid"
	"time"
)

type graduatesService struct {
	wisudawanrepo entities.GraduatesRepository
}

func NewGraduatesUsecase(a entities.GraduatesRepository) entities.GraduatesUsecase {
	return &graduatesService{wisudawanrepo: a}
}

func ConvertEntityGraduatesToSerializer(x entities.Graduates) entities.GetGraduatesSerializer {
	if x.Photo == "" {
		x.Photo = "PasFoto/default-wisudawan.png"
	}
	return entities.GetGraduatesSerializer{
		ID:           	x.ID,
		Identifier:   	x.Identifier,
		Name:         	x.Name,
		NickName:     	x.NickName,
		ThesisTitle:  	x.ThesisTitle,
		Incoming:     	x.Incoming,
		Major:       	x.Major.Major,
		MajorShort:  	x.Major.MajorShort,
		Faculty:      	x.Major.Faculty,
		FacultyShort: 	x.Major.FacultyShort,
		Instagram:     	x.Instagram,
		Linkedin:      	x.Linkedin,
		Twitter:       	x.Twitter,
		PlaceOfBirth:   x.PlaceOfBirth,
		DateOfBirth:  	x.DateOfBirth.Format("02-01-2006"),
		Photo:         	x.Photo,
	}
}

func ConvertEntityGraduatesToSimpleSerializer(x entities.Graduates) entities.GetSimpleGraduatesSerializer {
	if x.Photo == "" {
		x.Photo = "PasFoto/default-wisudawan.png"
	}
	return entities.GetSimpleGraduatesSerializer{
		ID:            	x.ID,
		Identifier:    	x.Identifier,
		Name:          	x.Name,
		ThesisTitle:   	x.ThesisTitle,
		Major:       	x.Major.Major,
		MajorShort:  	x.Major.MajorShort,
		Faculty:      	x.Major.Faculty,
		FacultyShort: 	x.Major.FacultyShort,
		Photo:         	x.Photo,
	}
}

func (service *graduatesService) CreateGraduates(item entities.CreateGraduatesSerializer) error {
	tglLahir, timeErr := time.Parse("01-02-2006", item.DateOfBirth)
	if timeErr != nil {
		tglLahir = time.Time{}
	}
	if err := service.wisudawanrepo.AddOne(
		item.Identifier,
		item.Incoming,
		item.Name,
		item.NickName,
		item.ThesisTitle,
		item.Major,
		item.Instagram,
		item.Linkedin,
		item.Twitter,
		item.PlaceOfBirth,
		item.Photo,
		tglLahir,
	); err != nil {
		return err
	}
	return nil
}

func (service *graduatesService) DeleteGraduates(idGraduates uuid.UUID) error {
	if err := service.wisudawanrepo.DeleteOne(idGraduates.String()); err != nil {
		return err
	}
	return nil
}

func (service *graduatesService) UpdateGraduates(item entities.UpdateGraduatesSerializer) error {
	tglLahir, timeErr := time.Parse("01-02-2006", item.DateOfBirth)

	if timeErr != nil {
		tglLahir = time.Time{}
	}

	if err := service.wisudawanrepo.UpdateOne(
		item.NIM,
		item.Incoming,
		item.Name,
		item.NickName,
		item.ThesisTitle,
		item.Major,
		item.Instagram,
		item.Linkedin,
		item.Twitter,
		item.DateOfBirth,
		item.Photo,
		tglLahir,
	); err != nil {
		return err
	}
	return nil
}

func (service *graduatesService) GetGraduates(idGraduates uuid.UUID) (entities.Graduates, error) {
	result, err := service.wisudawanrepo.GetOne(idGraduates.String())
	if err != nil {
		return result, err
	}
	return result, nil
}

func (service *graduatesService) GetAllGraduates() ([]entities.Graduates, error) {
	result, err := service.wisudawanrepo.GetAll()
	if err != nil {
		return result, err
	}
	return result, nil
}

func (service *graduatesService) FilterGraduatesByOrgzSlug(organizationSlug string) ([]entities.Graduates, error) {
	result, err := service.wisudawanrepo.FilterByOrgzSlug(organizationSlug)
	if err != nil {
		return result, err
	}
	return result, nil
}