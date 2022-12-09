package services

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/alishokri1661s/SMS-Gateway/Utils/conf"
	"github.com/alishokri1661s/SMS-Gateway/Utils/logs"
	"github.com/alishokri1661s/SMS-Gateway/Utils/models"
	amqp "github.com/rabbitmq/amqp091-go"
)

var connectRabbitMQ *amqp.Connection

func createMessageBrokerConnection() {
	mbSetting := conf.MessageBrokerSetting
	amqpURL := fmt.Sprintf("amqp://%s:%s@%s:%s/", mbSetting.User, mbSetting.Password, mbSetting.Host, mbSetting.Port)
	conn, err := amqp.Dial(amqpURL)
	logs.PanicOnError(err)
	connectRabbitMQ = conn
}

func publishMessage(sms models.SMS) error {
	ch, err := connectRabbitMQ.Channel()
	logs.LogOnError(err)
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"sms", // name
		true,    // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	logs.LogOnError(err)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	smsBytes, err := json.Marshal(sms)
	logs.LogOnError(err)

	message := amqp.Publishing{
		ContentType:  "application/json",
		Body:         smsBytes,
		DeliveryMode: amqp.Persistent,
	}

	err = ch.PublishWithContext(ctx,
		"",      // exchange
		q.Name,  // queue name
		false,   // mandatory
		false,   // immediate
		message, // message to publish
	)
	logs.LogOnError(err)

	return err
}
