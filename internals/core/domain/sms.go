package domain

import "gorm.io/gorm"

type MessageStatus string

const (
	Pending  MessageStatus = "pending"
	Sending  MessageStatus = "sending"
	Sent     MessageStatus = "sent"
	Received MessageStatus = "received"
)

type SMS struct {
	gorm.Model
	Id       int           `json:"id" gorm:"primary_key"`
	Sender   string        `json:"sender"`
	Receiver string        `json:"receiver"`
	Status   MessageStatus `json:"status"`
	Text     string        `json:"text"`
}
