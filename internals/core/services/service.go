package services

import (
	"log"

	"github.com/alishokri1661s/SMS-Gateway/internals/core/domain"
	"github.com/alishokri1661s/SMS-Gateway/internals/core/ports"
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
func (s *Service) SendSMS(sms domain.SMS) (domain.SMS, error) {
	var err error
	sms, err = s.repository.SendSMS(sms)
	if err != nil {
		log.Println(err)
	}

	//send to borker

	return sms, err

}

// LogSMS implements ports.IService
func (s *Service) LogSMS() ([]domain.SMS, error) {
	allsms, err := s.repository.LogSMS()
	if err != nil {
		log.Println(err)
	}
	return allsms, err
}
