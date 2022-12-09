package models

import "gorm.io/gorm"

type MessageStatus string

const (
	Pending  MessageStatus = "pending"
	Sending  MessageStatus = "sending"
	Sent     MessageStatus = "sent"
	Received MessageStatus = "received"
	Failed   MessageStatus = "failed"
)

type SMS struct {
	gorm.Model
	SenderType string        `json:"sender_type"`
	Receiver   string        `json:"receiver"`
	Status     MessageStatus `json:"status"`
	Text       string        `json:"text"`
}
