package main

import (
	c "PBP-API/src/controllers"
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

func main() {

	s := gocron.NewScheduler(time.UTC)

	// CONTOH BUAT DIIMPLEMENT NANTI
	s.Every(1).Seconds().Do(func() {
		fmt.Println("1")
	})

	s.Every(30).Seconds().Do(func() {
		fmt.Println("30")
	})

	s.Every(1).Minute().Do(func() {
		fmt.Println("1 Minute")
	})

	// CONTOH KALO PAKE TIMER
	s.Every(1).Day().At("09:00").Do(func() {
		fmt.Println("Good morning! This task runs every day at 9:00 AM.")
	})

	s.StartBlocking()

	//go routine
	c.TaskScheduler()
	select {}
}
