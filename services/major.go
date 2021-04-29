package services

import (
	"github.com/SemmiDev/go-backend/entities"
	uuid "github.com/satori/go.uuid"
)

type majorUseCase struct {
	majorrepo entities.MajorRepository
}

func NewMajorUsecase(j entities.MajorRepository) entities.MajorUseCase {
	return &majorUseCase{majorrepo: j}
}

func ConvertEntityMajorToSerializer(j entities.Major) entities.GetMajorSerializer {
	return entities.GetMajorSerializer{
		Id:            	j.ID,
		Major:       	j.Major,
		MajorShort:  	j.MajorShort,
		Faculty:      	j.Faculty,
		FacultyShort: 	j.FacultyShort,
	}
}

func (service *majorUseCase) CreateMajor(item entities.CreateMajorSerializer) error {
	if err := service.majorrepo.AddOne(item.Major, item.Faculty, item.FacultyShort, item.MajorShort); err != nil {
		return err
	}
	return nil
}

func (service *majorUseCase) DeleteMajor(IdMajor uuid.UUID) error {
	err := service.majorrepo.DeleteOne(IdMajor)
	if err == nil {
		return nil
	}
	return err
}

func (service *majorUseCase) UpdateMajor(item entities.UpdateMajorSerializer) error {
	err := service.majorrepo.UpdateOne(item.IdMajor, item.Major, item.Faculty, item.FacultyShort, item.MajorShort)
	if err != nil {
		return err
	}
	return nil
}

func (service *majorUseCase) GetMajor(IdMajor uuid.UUID) (entities.Major, error) {
	result, err := service.majorrepo.GetOne(IdMajor)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (service *majorUseCase) GetAllMajor() ([]entities.Major, error) {
	result, err := service.majorrepo.GetAll()
	if err != nil {
		return result, err
	}
	return result, nil
}