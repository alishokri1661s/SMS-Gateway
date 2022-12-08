package server

import (
	"github.com/alishokri1661s/SMS-Gateway/internals/core/ports"
	"github.com/alishokri1661s/SMS-Gateway/internals/routers"
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
	router.Run(":8000")
}
