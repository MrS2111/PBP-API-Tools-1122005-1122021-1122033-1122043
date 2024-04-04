package main

import (
	c "PBP-API/src/controllers"
	"fmt"
	"log"
	"time"

	"github.com/go-co-op/gocron"
	"gopkg.in/gomail.v2"
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

	select {}

	m := gomail.NewMessage()
	m.SetHeader("From", "jasonjeyys@gmail.com")
	m.SetHeader("To", "elliezerchristian@gmail.com")
	m.SetHeader("Subject", "Hello, Golang Email!")
	m.SetBody("text/plain", "This is the body of the email.")

	d := gomail.NewDialer("smtp.gmail.com", 587, "jasonjeyys@gmail.com", "testingemail")

	if err := d.DialAndSend(m); err != nil {
		log.Fatal(err)
	}
	c.TaskScheduler();
	select {}
}
