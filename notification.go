package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

	type Notification struct{
	Repo string
	Sub string
	Type string
	Action string
	Message string
}
	var resolver map[string]interface{}
	var visited []string

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func makeNotification(jsondata string)  {

	
	
	var data []map[string]interface{}
	err := json.Unmarshal([]byte(jsondata), &data)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range data {
		id:=v["id"].(string)
		if contains(visited, id) {
			continue
		}

		notification:=Notification{}
		notification.Repo = v["repository"].(map[string]interface{})["full_name"].(string)
		notification.Sub = v["subject"].(map[string]interface{})["title"].(string)
		notification.Type = v["subject"].(map[string]interface{})["type"].(string)
		notification.Action = v["reason"].(string)
		bodyURL := v["subject"].(map[string]interface{})["latest_comment_url"].(string)
		notification.Message = getBody(bodyURL)
		visited = append(visited, id)

		blockKitEncoding(notification.Repo, notification.Sub, notification.Type, notification.Action, notification.Message)

	}

	

	
	
}

func getBody(url string) string  {
	err:=godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	GITHUB_TOKEN:=os.Getenv("GITHUB_TOKEN")

	client:=http.Client{}
	req, err := http.NewRequest("GET", url, nil)
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

	json.Unmarshal(body, &resolver)

	comment:=resolver["body"].(string)
	return comment
	
	
}