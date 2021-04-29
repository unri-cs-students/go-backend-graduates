package services

import (
	"github.com/SemmiDev/go-backend/entities"
	contenttype "github.com/SemmiDev/go-backend/commons/content_type"
	uuid "github.com/satori/go.uuid"
	"strings"
)

type contentUsecase struct {
	contentrepo entities.ContentRepository
}

func NewContentUsecase(a entities.ContentRepository) entities.ContentUseCase {
	return &contentUsecase{contentrepo: a}
}

func (service *contentUsecase) CreateContent(item entities.CreateContentSerializer) error {
	if err := service.contentrepo.AddOne(
		item.Identifier,
		item.Organization,
		item.ContentType,
		item.Headings,
		item.Details,
		item.Image,
	); err != nil {
		return err
	}
	return nil
}

func (service *contentUsecase) DeleteContent(IdContent uuid.UUID) error {
	err := service.contentrepo.DeleteOne(IdContent.String())
	if err == nil {
		return nil
	}
	return err
}

func (service *contentUsecase) UpdateContent(item entities.UpdateContentSerializer) error {
	err := service.contentrepo.UpdateOne(
		item.Content,
		item.Identifier,
		item.Organization,
		item.ContentType,
		item.Headings,
		item.Details,
		item.Image,
	)
	if err != nil {
		return err
	}
	return nil
}

func (service *contentUsecase) GetContent(IdContent uuid.UUID) (entities.Content, error) {
	result, err := service.contentrepo.GetOne(IdContent.String())
	if err != nil {
		return result, err
	}
	return result, nil
}

func (service *contentUsecase) GetByGraduates(graduatesIdentifier uint32) ([]entities.Content, error) {
	result, err := service.contentrepo.GetByGraduates(graduatesIdentifier)
	if err != nil {
		return nil, err
	}
	return result, nil
}


func ConvertEntityContentsToSerializer(data []entities.Content) entities.GetContentsSerializer {
	var selfData []entities.GetContentSerializer
	var orgzData []entities.GetContentSerializer2
	for _, x := range data {
		if strings.EqualFold(x.Type, contenttype.Karya.String()) ||
			strings.EqualFold(x.Type, contenttype.Prestasi.String()) ||
			strings.EqualFold(x.Type, contenttype.Funfact.String()) ||
			strings.EqualFold(x.Type, contenttype.Tips.String()) {
			selfData = append(selfData, entities.GetContentSerializer{
				ContentType: x.Type,
				Headings:    x.Headings,
				Details:     x.Details,
				Image:       x.Image,
			})
		} else {
			orgzData = append(orgzData, entities.GetContentSerializer2{
				GetContentSerializer: entities.GetContentSerializer{
					ContentType: x.Type,
					Headings:    x.Headings,
					Details:     x.Details,
					Image:       x.Image,
				},
				OrganizationName: x.Organization.Name,
				OrganizationLogo: x.Organization.Logo,
			})
		}

	}
	if len(orgzData) == 0 {
		orgzData = make([]entities.GetContentSerializer2, 0)
	}

	if len(selfData) == 0 {
		selfData = make([]entities.GetContentSerializer, 0)
	}
	return entities.GetContentsSerializer{
		OrganizationalContents: orgzData,
		SelfContents:           selfData,
	}
}
