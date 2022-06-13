package models

import (
	"time"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type Message interface {
	Marshal(CommonMessage)
}

type CommonMessage map[string]interface{}

type TextMessage struct {
	ID      string
	Type    linebot.MessageType
	Text    string
	Emojis  []*linebot.Emoji
	Mention *linebot.Mention
}

type ImageMessage struct {
	ID                 string
	Type               linebot.MessageType
	OriginalContentURL string
	PreviewImageURL    string
	ContentProvider    *linebot.ContentProvider
	ImageSet           *linebot.ImageSet
}

type Event struct {
	UserId    string    `json:"userid"`
	Timestamp time.Time `json:"timestamp"`
	Message   Message   `json:"message"`
}

type CommonEvent struct {
	UserId      string              `json:"userid"`
	Timestamp   time.Time           `json:"timestamp"`
	Message     CommonMessage       `json:"message"`
	MessageType linebot.MessageType `json:"messagetype"`
}

type LineEvent struct {
	UserId      string
	Timestamp   time.Time
	Message     linebot.Message
	MessageType linebot.MessageType
}

func NewLineEvent(lineEvent *linebot.Event) *LineEvent {
	return &LineEvent{
		UserId:    lineEvent.Source.UserID,
		Timestamp: lineEvent.Timestamp,
		Message:   lineEvent.Message,
	}
}

func (e *LineEvent) SetType(msgtype linebot.MessageType) {
	e.MessageType = msgtype
}
