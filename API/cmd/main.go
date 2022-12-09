package main

import (
	"github.com/alishokri1661s/SMS-Gateway/API/internals/core/services"
	"github.com/alishokri1661s/SMS-Gateway/API/internals/handlers"
	"github.com/alishokri1661s/SMS-Gateway/API/internals/repositories"
	"github.com/alishokri1661s/SMS-Gateway/API/internals/server"
	"github.com/alishokri1661s/SMS-Gateway/Utils/conf"
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
