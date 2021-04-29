package services

import (
	"github.com/SemmiDev/go-backend/entities"
	uuid "github.com/satori/go.uuid"
)

type orgzUseCase struct {
	orgzrepo entities.OrgzRepository
}

func NewOrgzUsecase(a entities.OrgzRepository) entities.OrgzUseCase {
	return &orgzUseCase{orgzrepo: a}
}

func (service *orgzUseCase) CreateOrgz(item entities.CreateOrgzSerializer) error {
	if err := service.orgzrepo.AddOne(
		item.Name,
		item.Slug,
		item.Category,
		item.Logo,
		item.PosterAppreciation,
		item.WritingAppreciation,
		item.VideoAppreciation,
		item.FacultyShort,
	); err != nil {
		return err
	}
	return nil
}

func (service *orgzUseCase) DeleteOrgz(idOrgz uuid.UUID) error {
	err := service.orgzrepo.DeleteOne(idOrgz.String())
	if err == nil {
		return nil
	}
	return err
}

func (service *orgzUseCase) UpdateOrgz(item entities.UpdateOrgzSerializer) error {
	err := service.orgzrepo.UpdateOne(
		item.IdOrgz,
		item.Name,
		item.Slug,
		item.Category,
		item.Logo,
		item.PosterAppreciation,
		item.WritingAppreciation,
		item.VideoAppreciation,
		item.FacultyShort,
	)
	if err != nil {
		return err
	}
	return nil
}

func (service *orgzUseCase) GetOrgz(idOrgz uuid.UUID) (entities.Orgz, error) {
	result, err := service.orgzrepo.GetOne(idOrgz.String())
	if err != nil {
		return result, err
	}
	return result, nil
}

func (service *orgzUseCase) GetAll() ([]entities.Orgz, error) {
	result, err := service.orgzrepo.GetAll()
	if err != nil {
		return result, err
	}
	return result, nil
}

func (service *orgzUseCase) GetBySlug(slug string) (entities.Orgz, error) {
	result, err := service.orgzrepo.GetBySlug(slug)
	if err != nil {
		return result, err
	}
	return result, nil
}