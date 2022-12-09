package ports

import (
	"github.com/alishokri1661s/SMS-Gateway/Utils/models"
	"github.com/gin-gonic/gin"
)

type IRepository interface {
	SendSMS(sms models.SMS) (models.SMS, error)
	LogSMS() ([]models.SMS, error)
}

type IService interface {
	SendSMS(sms models.SMS) (models.SMS, error)
	LogSMS() ([]models.SMS, error)
}

type IHandler interface {
	SendSMS(ctx *gin.Context)
	LogSMS(ctx *gin.Context)
}

type IServer interface {
	Initialize()
}
