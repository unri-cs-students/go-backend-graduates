package services

import (
	"github.com/SemmiDev/go-backend/entities"
	uuid "github.com/satori/go.uuid"
)

type messageUsecase struct {
	messagerepo entities.MessageRepository
}

func NewMessageUsecase(j entities.MessageRepository) entities.MessageUsecase {
	return &messageUsecase{messagerepo: j}
}

func ConvertEntityMessageToSerializer(x entities.Message) entities.GetMessageSerializer {
	return entities.GetMessageSerializer{
		ID:      x.ID,
		Sent:    x.CreatedAt.Format("15:04:05 02-01-2006"),
		Sender:  x.Sender,
		Message: x.Message,
	}
}

func (service *messageUsecase) CreateMessage(item entities.CreateMessageSerializer) error {
	if err := service.messagerepo.AddOne((item.Message), (item.Sender), (item.IdGraduates)); err != nil {
		return err
	}
	return nil
}

func (service *messageUsecase) DeleteMessage(idMessage uuid.UUID) error {
	if err := service.messagerepo.DeleteOne(idMessage.String()); err != nil {
		return err
	}
	return nil
}

func (service *messageUsecase) GetMessage(idWisudawan uuid.UUID) ([]entities.Message, error) {
	result, err := service.messagerepo.GetMessage(idWisudawan.String())
	if err != nil {
		return result, err
	}
	return result, nil
}