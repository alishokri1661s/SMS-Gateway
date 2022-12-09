package routers

import (
	"github.com/alishokri1661s/SMS-Gateway/API/internals/core/ports"
	"github.com/gin-gonic/gin"
)

func GetRoutes(router *gin.Engine, handler ports.IHandler) {
	sms := router.Group("/sms")
	//send message
	sms.POST("/send", handlerFunc(handler.SendSMS))
	//log message
	sms.GET("/log", handlerFunc(handler.LogSMS))
}

// This function gets "func(*gin.Context)" and returns "gin.HandlerFunc"
func handlerFunc(v interface{}) gin.HandlerFunc {
	switch v := v.(type) {
	case func(*gin.Context):
		return v
	case gin.HandlerFunc:
		return v
	default:
		panic("unexpected type")
	}
}
