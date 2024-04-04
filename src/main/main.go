package main

import (
	"PBP-API/src/controllers"
	"log"
	"time"

	"github.com/go-co-op/gocron"
)

func main() {
	scheduler := gocron.NewScheduler(time.FixedZone("WIB", 7*60*60))

	_, err := scheduler.Every(10).Seconds().Do(func() {
		log.Println("Executing email sending task...")
		controllers.SendEmail()
	})
	if err != nil {
		log.Fatal("Error scheduling email task:", err)
	}
	scheduler.StartAsync()
	time.Sleep(20 * time.Second)
	scheduler.Stop()
	time.Sleep(2 * time.Second)
	log.Println("Program selesai.")
}
