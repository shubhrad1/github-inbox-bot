package main

import (
	"fmt"
)

type Block struct {
	Type string `json:"type"`
	Text Text   `json:"text"`
}

type Text struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type Message struct {
	Blocks []Block `json:"blocks"`
}
type Underline struct {
	Type string `json:"type"`
}


func blockKitEncoding(repo, subject, notifType, reason, body string) {
	msg:= Message{
		Blocks: []Block{
			{
				Type: "section",
				Text: Text{
					Type: "mrkdwn",
					Text: ":rocket:  *New Notification on Github*",
				},
			},
			{
				Type: "section",
				Text: Text{
					Type: "mrkdwn",
					Text: fmt.Sprintf("*Repo:* %s", repo),
				},
			},
			{
				Type: "section",
				Text: Text{
					Type: "mrkdwn",
					Text: fmt.Sprintf("*Subject:* %s", subject),
				},
			},
			{
				Type: "section",
				Text: Text{
					Type: "mrkdwn",
					Text: fmt.Sprintf("*Type:* %s", notifType),
				},
			},
			{
				Type: "section",
				Text: Text{
					Type: "mrkdwn",
					Text: fmt.Sprintf("*Action:* %s", reason),
				},
			},
			{
				Type: "section",
				Text: Text{
					Type: "mrkdwn",
					Text: fmt.Sprintf("*Message:* %s", body),
				},
			},
			{
				Type:"section",
				Text: Text{
					Type: "mrkdwn",
					Text: fmt.Sprintln(""),
				},
			},
			

		},
	}
	
	sendNotification(msg)
}
