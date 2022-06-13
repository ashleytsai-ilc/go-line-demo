package controllers

import (
	"context"
	"fmt"
	"go-line-demo/database"
	"go-line-demo/models"
	"go-line-demo/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func GetHi(request *gin.Context) {
	request.IndentedJSON(http.StatusOK, gin.H{
		"message": "Hi there!",
	})
}

func ReceiveMessage(request *gin.Context) {
	bot := utils.GetLinebot()
	events, err := bot.ParseRequest(request.Request)
	if err != nil {
		panic(err)
	}

	collection := database.GetClient().Database("chat").Collection("messages")
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			eventData := models.NewLineEvent(event)
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				eventData.SetType(linebot.MessageTypeText)
				_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do()
				if err != nil {
					log.Print(err)
				}
			case *linebot.ImageMessage:
				eventData.SetType(linebot.MessageTypeImage)
				_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("Got image(s)!")).Do()
				if err != nil {
					log.Print(err)
				}
			}

			result, err := collection.InsertOne(context.TODO(), eventData)
			if err != nil {
				log.Print(err)
			}
			fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
		}
	}
}