package services

import (
	"github.com/alishokri1661s/SMS-Gateway/API/internals/core/ports"
	"github.com/alishokri1661s/SMS-Gateway/Utils/logs"
	"github.com/alishokri1661s/SMS-Gateway/Utils/models"
)

type Service struct {
	repository ports.IRepository
}

// This line is for get feedback in case we are not implementing the interface correctly
var _ ports.IService = (*Service)(nil)

func NewService(repository ports.IRepository) *Service {
	createMessageBrokerConnection()
	return &Service{
		repository: repository,
	}
}

// SendSMS implements ports.IService
func (s *Service) SendSMS(sms models.SMS) (models.SMS, error) {
	sms, err := s.repository.SendSMS(sms)
	logs.LogOnError(err)

	//send to borker
	err = publishMessage(sms)

	return sms, err
}

// LogSMS implements ports.IService
func (s *Service) LogSMS() ([]models.SMS, error) {
	allsms, err := s.repository.LogSMS()
	logs.LogOnError(err)

	return allsms, err
}
