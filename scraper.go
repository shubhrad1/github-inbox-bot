package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var GITHUB_TOKEN string
var BASE_URL string
var result []map[string]interface{}

func Scraper() {

	

	err:=godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	GITHUB_TOKEN=os.Getenv("GITHUB_TOKEN")
	BASE_URL="https://api.github.com/notifications"

	if GITHUB_TOKEN == "" {	
		log.Fatal("GITHUB_TOKEN is required")
	}
	
	client:=http.Client{}
	req, err := http.NewRequest("GET", BASE_URL, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", "token "+GITHUB_TOKEN)
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body,err:=io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(body, &result)
	if len(result) == 0 {
		visited=make([]string, 0)
	}
	jsonResult, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}
	jsonString := string(jsonResult)
	
	
	makeNotification(jsonString)

	

}