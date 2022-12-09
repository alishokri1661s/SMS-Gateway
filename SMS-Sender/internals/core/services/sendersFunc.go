package services

import (
	"github.com/alishokri1661s/SMS-Gateway/SMS-Sender/internals/core/ports"
	"github.com/alishokri1661s/SMS-Gateway/Utils/models"
)

func GetSenderFunction(senderType string) func(models.SMS) error {
	var sender ports.ISenderServcie
	switch senderType {
	case "mock":
		sender = (*Mock)(nil)
	case "ghasedak":
		sender = (*Ghasedak)(nil)
	case "kavenegar":
		sender = (*Kavenegar)(nil)
	default:
		return nil
	}
	return sender.SendSMS
}
