package handlers

import (
	"log"

	"github.com/alishokri1661s/SMS-Gateway/internals/core/domain"
	"github.com/alishokri1661s/SMS-Gateway/internals/core/ports"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service ports.IService
}

var _ ports.IHandler = (*Handler)(nil)

func NewHandlers(service ports.IService) *Handler {
	return &Handler{
		service: service,
	}
}

// SendSMS implements ports.IHandler
func (h *Handler) SendSMS(ctx *gin.Context) {
	var sms domain.SMS
	if err := ctx.BindJSON(&sms); err != nil {
		log.Fatal(err)
	}

	err := h.service.SendSMS(sms)
	if err != nil {
		log.Fatal(err)
	}
}

// LogSMS implements ports.IHandler
func (h *Handler) LogSMS(ctx *gin.Context) {
	panic("unimplemented")
}
