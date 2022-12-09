package ports

import (
	"github.com/alishokri1661s/SMS-Gateway/internals/core/domain"
	"github.com/gin-gonic/gin"
)

type IRepository interface {
	SendSMS(sms domain.SMS) (domain.SMS, error)
	LogSMS() ([]domain.SMS, error)
}

type IService interface {
	SendSMS(sms domain.SMS) (domain.SMS, error)
	LogSMS() ([]domain.SMS, error)
}

type IHandler interface {
	SendSMS(ctx *gin.Context)
	LogSMS(ctx *gin.Context)
}

type IServer interface {
	Initialize()
}
