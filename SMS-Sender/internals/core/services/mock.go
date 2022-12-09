package services

import (
	"github.com/alishokri1661s/SMS-Gateway/SMS-Sender/internals/core/ports"
	"github.com/alishokri1661s/SMS-Gateway/Utils/models"
)

type Mock struct{}

var _ ports.ISenderServcie = (*Mock)(nil)

func (*Mock) SendSMS(sms models.SMS) error {
	//sending message
	return nil
}
