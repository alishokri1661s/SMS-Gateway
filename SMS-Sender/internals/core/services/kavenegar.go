package services

import (
	"fmt"

	"github.com/alishokri1661s/SMS-Gateway/SMS-Sender/internals/core/ports"
	"github.com/alishokri1661s/SMS-Gateway/Utils/conf"
	"github.com/alishokri1661s/SMS-Gateway/Utils/models"
	"github.com/kavenegar/kavenegar-go"
)

type Kavenegar struct{}

var _ ports.ISenderServcie = (*Kavenegar)(nil)

func (*Kavenegar) SendSMS(sms models.SMS) error {
	api := kavenegar.New(conf.KavenegarAPIKey)
	sender := ""
	receptor := []string{sms.Receiver}
	message := sms.Text
	if res, err := api.Message.Send(sender, receptor, message, nil); err != nil {
		switch err := err.(type) {
		case *kavenegar.APIError:
			fmt.Println(err.Error())
		case *kavenegar.HTTPError:
			fmt.Println(err.Error())
		default:
			fmt.Println(err.Error())
		}
		return err
	} else {
		for _, r := range res {
			fmt.Println("MessageID 	= ", r.MessageID)
			fmt.Println("Status    	= ", r.Status)
		}
		return nil
	}
}
