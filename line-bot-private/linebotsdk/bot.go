package linebotsdk

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/line/line-bot-sdk-go/linebot"
)

type BotConfig struct {
	ChannelSecret      string
	ChannelAccessToken string
}

type Bot struct {
	config *BotConfig
	client *linebot.Client
}

func NewBot(config *BotConfig) *Bot {
	instance := Bot{config: config}
	client, err := linebot.New(config.ChannelSecret, config.ChannelAccessToken)
	if err != nil {
		return nil
	}
	instance.client = client
	return &instance
}

func convertLineBotRequest(requestBody string) (*Request, error) {
	var request Request
	err := json.Unmarshal([]byte(requestBody), &request)
	if err != nil {
		return nil, err
	}
	if len(request.Events) < 1 {
		return nil, errors.New("Not even one receive event")
	}
	return &request, nil
}

func (b *Bot) ReplyText(requestBody string, replying func(event *RequestEvent) string) error {
	request, err := convertLineBotRequest(requestBody)
	if err != nil {
		return err
	}
	eventDetail := request.Events[0]
	response, err := b.client.ReplyMessage(eventDetail.ReplyToken, linebot.NewTextMessage(replying(eventDetail))).Do()
	if err != nil {
		return err
	}
	fmt.Println(response)
	return nil
}
