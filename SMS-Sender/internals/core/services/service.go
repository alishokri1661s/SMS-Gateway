package services

import (
	"log"

	"github.com/alishokri1661s/SMS-Gateway/SMS-Sender/internals/core/ports"
	"github.com/alishokri1661s/SMS-Gateway/Utils/logs"
	"github.com/alishokri1661s/SMS-Gateway/Utils/models"
)

type Service struct {
	repository ports.IRepository
}

// This line is for get feedback in case we are not implementing the interface correctly
var _ ports.IService = (*Service)(nil)

func NewService(repository ports.IRepository) *Service {
	return &Service{
		repository: repository,
	}
}

// SendSMS implements ports.IService
func (s *Service) SendSMS(sms models.SMS) {
	//update status to sending
	sms, err := s.repository.UpdateSmsStatus(sms, models.Sending)
	logs.LogOnError(err)

	//send the message
	sendFunction := GetSenderFunction(sms.SenderType)
	if sendFunction != nil {
		err := sendFunction(sms)
		if err == nil {
			sms, err = s.repository.UpdateSmsStatus(sms, models.Sent)
		} else {
			sms, err = s.repository.UpdateSmsStatus(sms, models.Failed)
		}
		logs.LogOnError(err)
	} else {
		log.Println("Send Function is null")
	}

}
