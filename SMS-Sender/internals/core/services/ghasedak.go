package services

import (
	"fmt"
	"log"

	"github.com/alishokri1661s/SMS-Gateway/SMS-Sender/internals/core/ports"
	"github.com/alishokri1661s/SMS-Gateway/Utils/conf"
	"github.com/alishokri1661s/SMS-Gateway/Utils/models"

	ghasedak "github.com/ghasedakapi/ghasedak-go"
)

type Ghasedak struct{}

var _ ports.ISenderServcie = (*Ghasedak)(nil)

func (*Ghasedak) SendSMS(sms models.SMS) error {
	log.Println(conf.GhasedakAPIKey)
	c := ghasedak.NewClient(conf.GhasedakAPIKey, "")
	r := c.Send(sms.Text, sms.Receiver)
	log.Println(r)
	if r.Success {
		return nil
	} else {
		return fmt.Errorf("failed. code: %d", r.Code)
	}

}
