package controllers

import (
	"log"
	"time"

	"gopkg.in/gomail.v2"
)

func sendEmail() {
	m := gomail.NewMessage()
	m.SetHeader("From", "jasonjeyys@gmail.com")
	m.SetHeader("To", "elliezerchristian@gmail.com")
	m.SetHeader("Subject", "This is your reminder email!")
	m.SetBody("text/plain", "This email is automated. Here is your reminder email. Have a great day!")

	d := gomail.NewDialer("smtp.gmail.com", 587, "jasonjeyys@gmail.com", "testingemail")

	if err := d.DialAndSend(m); err != nil {
		log.Fatal(err)
	}
}

func TaskScheduler() {
	startTime := time.Now()

	interval := time.Minute * 5

	for {
		now := time.Now()
		nextRun := startTime.Add(interval)

		if now.After(nextRun) || now.Equal(nextRun) {
			sendEmail()

			startTime = time.Now()
		}

		time.Sleep(time.Second)
	}
}
