package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func ExpiryAlert(wg *sync.WaitGroup, expiryChannel chan bool) {
	// This function will send an alert to the user when the token is about to expire
	defer wg.Done()

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	SLACK_TOKEN := os.Getenv("SLACK_TOKEN")
	SLACK_CHANNEL := os.Getenv("SLACK_CHANNEL")

	api := slack.New(SLACK_TOKEN)
	
	days:=90
	_, _, err = api.PostMessage(SLACK_CHANNEL, slack.MsgOptionText(fmt.Sprintf("[INFO] Token will expire in %d days",days), false))
	if err != nil {
		log.Println("Error sending message")
	}
	



	START_DATE := os.Getenv("GHTOKEN_START_DATE")

	startDate, err := time.Parse("2006-01-02", START_DATE)

	if err != nil {
		fmt.Println("Error parsing date")
	}
	EXPIRY_DATE := startDate.AddDate(0, 0, 90)


	for {
		daysUntilExpiry := int(time.Until(EXPIRY_DATE).Hours() / 24)
		if daysUntilExpiry <= 5 {
			_, _, err = api.PostMessage(SLACK_CHANNEL, slack.MsgOptionText(fmt.Sprintf("ALERT!!! Token will expire in %d days", daysUntilExpiry), false))
			if err != nil {
				log.Println("Error sending message")
			}

		}
		if daysUntilExpiry == 0 {
			fmt.Println("Token has expired")
			_, _, err = api.PostMessage(SLACK_CHANNEL, slack.MsgOptionText("ALERT!!! Token has expired. Renew GitHubToken", false))
			if err != nil {
				log.Println("Error sending message")
			}
			expiryChannel <- true
			break
		}
		
		time.Sleep(24 * time.Hour)
	}
	
}


