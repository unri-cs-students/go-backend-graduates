package services

import (
	"errors"
	"fmt"
	"github.com/SemmiDev/go-backend/entities"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

type viewUseCase struct {
	Viewrepo entities.ViewRepository
}

func NewViewUsecase(v entities.ViewRepository) entities.ViewUseCase {
	return &viewUseCase{Viewrepo: v}
}

func (service *viewUseCase) AddView(idWisudawan uuid.UUID, clientIP string) error {
	lastRecord, lastErr := service.Viewrepo.GetLast(idWisudawan.String(), clientIP)
	if lastErr != nil {
		if !errors.Is(lastErr, gorm.ErrRecordNotFound) {
			return lastErr
		}
	} else {
		diff := time.Now().Sub(lastRecord.AccessTime).Minutes()
		fmt.Println(diff)
		if diff < 10 {
			return nil
		}
	}

	if err := service.Viewrepo.AddOne(
		idWisudawan.String(),
		clientIP,
		time.Now(),
	); err != nil {
		return err
	}
	return nil
}

func (service *viewUseCase) GetTop5() ([]entities.GetViewGraduates, error) {
	result, err := service.Viewrepo.GetTop5()
	if err != nil {
		return result, err
	}
	return result, nil
}