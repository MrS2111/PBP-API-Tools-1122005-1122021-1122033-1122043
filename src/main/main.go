package main

import (
	"PBP-API/src/controllers"
	"github.com/go-co-op/gocron"
	"log"
	"time"
)

func main() {
	scheduler := gocron.NewScheduler(time.FixedZone("WIB", 7*60*60))

	_, err := scheduler.Every(5).Minutes().Do(func() {
		log.Println("Executing email sending task...")
		controllers.SendEmail()
	})
	if err != nil {
		log.Fatal("Error scheduling email task:", err)
	}
	scheduler.StartAsync()
	select {}
}
