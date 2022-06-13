package utils

import (
	"go-line-demo/config"
	"net/http"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func GetLinebot() *linebot.Client {
	settings := config.GetConfig()
	channelSecret := settings.GetString("LINE_CHANNEL_SECRET")
	channelAccess := settings.GetString("LINE_CHANNEL_ACCESS")
	client := &http.Client{}
	bot, err := linebot.New(channelSecret, channelAccess, linebot.WithHTTPClient(client))
	if err != nil {
		panic(err)
	}

	return bot
}
