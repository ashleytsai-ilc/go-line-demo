package utils

import (
	"go-line-demo/config"
	"net/http"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

var linebotClient *linebot.Client

func NewLinebot() {
	settings := config.GetConfig()
	channelSecret := settings.GetString("LINE_CHANNEL_SECRET")
	channelAccess := settings.GetString("LINE_CHANNEL_ACCESS")
	client := &http.Client{}
	var err error
	linebotClient, err = linebot.New(channelSecret, channelAccess, linebot.WithHTTPClient(client))
	if err != nil {
		panic(err)
	}
}

func GetLinebot() *linebot.Client {
	return linebotClient
}
