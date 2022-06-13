package controllers

import (
	"context"
	"fmt"
	"go-line-demo/database"
	"go-line-demo/models"
	"go-line-demo/utils"
	"go-line-demo/validators"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"go.mongodb.org/mongo-driver/bson"
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

func PushMessage(request *gin.Context) {
	var pushMsg validators.PushMessage

	if err := request.ShouldBindWith(&pushMsg, binding.JSON); err != nil {
		request.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		message := linebot.NewTextMessage(pushMsg.Content)
		bot := utils.GetLinebot()
		_, err := bot.PushMessage(pushMsg.UserID, message).Do()
		if err != nil {
			log.Print(err)
		}

		request.IndentedJSON(http.StatusOK, gin.H{"message": "Push message successfully."})
	}
}

func GetMessages(request *gin.Context) {
	var getMsg validators.GetMessages

	if err := request.ShouldBindUri(&getMsg); err != nil {
		request.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		var results = []*models.Event{}
		collection := database.GetClient().Database("chat").Collection("messages")
		cursor, err := collection.Find(context.TODO(), bson.D{{"userid", getMsg.UserID}})
		if err != nil {
			panic(err)
		}
		for cursor.Next(context.TODO()) {
			var commonEvent models.CommonEvent
			if err := cursor.Decode(&commonEvent); err != nil {
				panic(err)
			}

			newEvent := models.NewEvent(commonEvent)
			msgType := commonEvent.MessageType
			var message models.Message
			switch msgType {
			case linebot.MessageTypeText:
				message = &models.TextMessage{Type: msgType}
			case linebot.MessageTypeImage:
				message = &models.ImageMessage{Type: msgType}
			}
			message.Marshal(commonEvent.Message)
			newEvent.Message = message
			results = append(results, newEvent)
		}
		if err := cursor.Err(); err != nil {
			panic(err)
		}
		request.IndentedJSON(http.StatusOK, results)
	}
}
