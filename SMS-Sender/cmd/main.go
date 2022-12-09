package main

import (
	"github.com/alishokri1661s/SMS-Gateway/SMS-Sender/internals/consumers"
	"github.com/alishokri1661s/SMS-Gateway/SMS-Sender/internals/core/services"
	"github.com/alishokri1661s/SMS-Gateway/SMS-Sender/internals/repositories"
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

	//cosumers
	consumer := consumers.NewConsumer(service)
	consumer.ConsumeMessage()

}
