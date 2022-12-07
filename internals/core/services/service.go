package services

import (
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
func (s *Service) SendSMS(sms domain.SMS) error {
	err := s.repository.SendSMS(sms)
	if err != nil {
		return err
	}

	//send to borker

	return nil

}

// LogSMS implements ports.IService
func (s *Service) LogSMS() error {
	err := s.repository.LogSMS()
	if err != nil {
		return err
	}
	return nil
}
