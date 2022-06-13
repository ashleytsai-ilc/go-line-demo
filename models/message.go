package models

import (
	"encoding/json"
	"fmt"
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

func NewEvent(event CommonEvent) *Event {
	return &Event{
		UserId:    event.UserId,
		Timestamp: event.Timestamp,
	}
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

func (m *TextMessage) Marshal(cm CommonMessage) {
	jsonStr, err := json.Marshal(cm)
	if err != nil {
		fmt.Println(err)
	}
	if err := json.Unmarshal(jsonStr, &m); err != nil {
		fmt.Println(err)
	}
}

func (m *ImageMessage) Marshal(cm CommonMessage) {
	jsonStr, err := json.Marshal(cm)
	if err != nil {
		fmt.Println(err)
	}
	if err := json.Unmarshal(jsonStr, &m); err != nil {
		fmt.Println(err)
	}
}

func (m *TextMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type    linebot.MessageType `json:"type"`
		Text    string              `json:"text"`
		Emojis  []*linebot.Emoji    `json:"emojis,omitempty"`
		Mention *linebot.Mention    `json:"mention,omitempty"`
	}{
		Type:    m.Type,
		Text:    m.Text,
		Emojis:  m.Emojis,
		Mention: m.Mention,
	})
}

func (m *ImageMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type               linebot.MessageType      `json:"type"`
		OriginalContentURL string                   `json:"originalContentUrl"`
		PreviewImageURL    string                   `json:"previewImageUrl"`
		ContentProvider    *linebot.ContentProvider `json:"contentProvider,omitempty"`
		ImageSet           *linebot.ImageSet        `json:"imageSet,omitempty"`
	}{
		Type:               m.Type,
		OriginalContentURL: m.OriginalContentURL,
		PreviewImageURL:    m.PreviewImageURL,
		ContentProvider:    m.ContentProvider,
		ImageSet:           m.ImageSet,
	})
}
