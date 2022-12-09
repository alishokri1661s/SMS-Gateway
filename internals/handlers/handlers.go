package handlers

import (
	"log"
	"net/http"

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
		log.Println(err)
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	var err error
	sms, err = h.service.SendSMS(sms)

	if err == nil {
		ctx.IndentedJSON(http.StatusCreated, sms)
	} else {
		ctx.IndentedJSON(http.StatusBadRequest, err)
	}
}

// LogSMS implements ports.IHandler
func (h *Handler) LogSMS(ctx *gin.Context) {
	allSMS, err := h.service.LogSMS()

	if err == nil {
		ctx.IndentedJSON(http.StatusOK, allSMS)
	} else {
		ctx.IndentedJSON(http.StatusBadRequest, err)
	}

}
