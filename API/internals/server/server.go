package server

import (
	"fmt"
	"log"

	"github.com/alishokri1661s/SMS-Gateway/API/internals/core/ports"
	"github.com/alishokri1661s/SMS-Gateway/API/internals/routers"
	"github.com/alishokri1661s/SMS-Gateway/Utils/conf"
	"github.com/gin-gonic/gin"
)

type Server struct {
	handler ports.IHandler
}

var _ ports.IServer = (*Server)(nil)

func NewServer(handler ports.IHandler) *Server {
	return &Server{
		handler: handler,
	}
}

// Initialize implements ports.IServer
func (s *Server) Initialize() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	routers.GetRoutes(router, s.handler)
	serverAddress := fmt.Sprintf("%s:%s", conf.ServerSetting.Host, conf.ServerSetting.Port)
	log.Printf("Server started on %s\n", serverAddress)
	router.Run(serverAddress)
}
