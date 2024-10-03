package main

import (
	"log"
	"sync"
	"time"
)
func main() {
	
	
	log.Println("Bot Started")

	expiryChannel:=make(chan bool)
	wg:=&sync.WaitGroup{}
	wg.Add(1)
	go ExpiryAlert(wg, expiryChannel)
	
	ticker:=time.NewTicker(1*time.Minute)
	for ;;<-ticker.C {
	select {
	case <-expiryChannel:
		log.Println("Token has expired")
		wg.Wait()
		return
	default:
			Scraper()
	}
}


}