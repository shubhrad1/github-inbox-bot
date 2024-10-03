package main

import (
	"log"
	"time"
)
func main() {
	
	ticker:=time.NewTicker(1*time.Minute)
	log.Println("Bot Started")
	for ;;<-ticker.C {
		Scraper()
	}
}