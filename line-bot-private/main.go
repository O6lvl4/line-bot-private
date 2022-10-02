package main

import (
	"encoding/json"
	"fmt"
	"line-bot-private/linebotsdk"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func printStructure(structure interface{}) {
	fmt.Printf("(%%#v) %#v\n", structure)
}

func convertLineBotRequest(requestBody string) (*linebotsdk.RequestEvent, error) {
	var event linebotsdk.RequestEvent
	err := json.Unmarshal([]byte(requestBody), &event)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	printStructure(request)
	// LINE_BOT_CHANNEL_ACCESS_TOKEN=SCUMtq0OiYG2PamQZAI9PXpVSo0ZAvWP/R2NgkzLVHp1XXy2Ns2E5acfZr0VIhnkARM0A6rO8rsx26z3+6jBdSvo5PD1YwaFa/X8LxXZasuRlVpmw72U1q5Z3dCcAjFelLgk+UYWFxg8a2h0AybqiAdB04t89/1O/w1cDnyilFU=
	// LINE_BOT_CHANNEL_SECRET=4946fec807651d6dec07ceb6c70f2404
	channelSecret := os.Getenv("LINE_BOT_CHANNEL_SECRET")
	channelAccessToken := os.Getenv("LINE_BOT_CHANNEL_ACCESS_TOKEN")
	bot := linebotsdk.NewBot(&linebotsdk.BotConfig{
		ChannelSecret:      channelSecret,
		ChannelAccessToken: channelAccessToken,
	})
	printStructure(bot)
	err := bot.ReplyText(request.Body, func(event *linebotsdk.RequestEvent) string {
		return `こんにちは、あなたのコンシェルジュです。
今はまだ準備中なので、お願いしたいことがあればリクエストしてくださいね`
	})
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("Succeeded reply message"),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
