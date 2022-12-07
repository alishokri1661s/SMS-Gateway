package main

import (
	"github.com/alishokri1661s/SMS-Gateway/conf"
	"github.com/alishokri1661s/SMS-Gateway/internals/core/services"
	"github.com/alishokri1661s/SMS-Gateway/internals/handlers"
	"github.com/alishokri1661s/SMS-Gateway/internals/repositories"
	"github.com/alishokri1661s/SMS-Gateway/internals/server"
)

func init() {
	conf.Setup()
	
}

func main() {
	//repositories
	repository := repositories.CreateDabaseConnection()

	//services
	service := services.NewService(repository)

	//handlers
	handler := handlers.NewHandlers(service)

	//server
	server.NewServer(handler).Initialize()
}
