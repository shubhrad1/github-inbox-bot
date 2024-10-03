package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

var SLACK_TOKEN string
var SLACK_CHANNEL string


func sendNotification(notification Message) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	SLACK_TOKEN = os.Getenv("SLACK_TOKEN")
	SLACK_CHANNEL = os.Getenv("SLACK_CHANNEL")

	var blocks []slack.Block
	for _, block := range notification.Blocks {
		blocks = append(blocks, slack.NewSectionBlock(slack.NewTextBlockObject(block.Text.Type, block.Text.Text, false, false), nil, nil))
	}

	api:=slack.New(SLACK_TOKEN)
	

	_, _, err = api.PostMessage(SLACK_CHANNEL, slack.MsgOptionBlocks(blocks...))
	if err != nil {
		log.Fatal(err)
	} 
}
