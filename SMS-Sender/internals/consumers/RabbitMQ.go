package consumers

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/alishokri1661s/SMS-Gateway/SMS-Sender/internals/core/ports"
	"github.com/alishokri1661s/SMS-Gateway/Utils/conf"
	"github.com/alishokri1661s/SMS-Gateway/Utils/logs"
	"github.com/alishokri1661s/SMS-Gateway/Utils/models"
	amqp "github.com/rabbitmq/amqp091-go"
)

var connectRabbitMQ *amqp.Connection

type Consumer struct {
	service ports.IService
}

var _ ports.IConsumer = (*Consumer)(nil)

func NewConsumer(service ports.IService) *Consumer {
	createMessageBrokerConnection()
	return &Consumer{
		service: service,
	}
}

func createMessageBrokerConnection() {
	mbSetting := conf.MessageBrokerSetting
	amqpURL := fmt.Sprintf("amqp://%s:%s@%s:%s/", mbSetting.User, mbSetting.Password, mbSetting.Host, mbSetting.Port)
	conn, err := amqp.Dial(amqpURL)
	logs.PanicOnError(err)
	connectRabbitMQ = conn
}

func (c *Consumer) ConsumeMessage() {
	ch, err := connectRabbitMQ.Channel()
	logs.LogOnError(err)
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"sms", // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	logs.LogOnError(err)

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	logs.LogOnError(err)

	var forever chan struct{}
	var sms models.SMS
	go func() {
		for d := range msgs {
			json.Unmarshal(d.Body, &sms)
			c.service.SendSMS(sms)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
