include .env

.PHONY: build
build:
	sam build

.PHONY: deploy
deploy: build
	sam deploy --parameter-overrides \
		LineBotChannelSecret=${LINE_BOT_CHANNEL_SECRET} LineBotChannelAccessToken=${LINE_BOT_CHANNEL_ACCESS_TOKEN}
