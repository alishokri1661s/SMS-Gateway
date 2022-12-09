package ports

import (
	"github.com/alishokri1661s/SMS-Gateway/Utils/models"
)

type IRepository interface {
	UpdateSmsStatus(sms models.SMS, status models.MessageStatus) (models.SMS, error)
}

type IService interface {
	SendSMS(models.SMS)
}

type IConsumer interface {
	ConsumeMessage()
}

type ISenderServcie interface {
	SendSMS(models.SMS) error
}
