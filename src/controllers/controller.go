package controllers

import (
	"fmt"
	"net/smtp"
	"time"
)

func sendEmail(to, subject, body string) {
	smtpServer := "smtp.gmail.com"
	smtpPort := 587
	sender := "jasonjeyys@gmail.com"
	password := "testingemail"

	message := []byte(fmt.Sprintf("To: %s\r\n"+
		"Subject: %s\r\n"+
		"\r\n"+
		"%s\r\n", to, subject, body))

	auth := smtp.PlainAuth("", sender, password, smtpServer)
	err := smtp.SendMail(fmt.Sprintf("%s:%d", smtpServer, smtpPort), auth, sender, []string{to}, message)
	if err != nil {
		fmt.Println("Failed to send email:", err)
		return
	}

	fmt.Println("Email sent successfully to", to)
}
func TaskScheduler() {
	startTime := time.Now()

	interval := time.Minute * 5

	for {
		now := time.Now()
		nextRun := startTime.Add(interval)

		if now.After(nextRun) || now.Equal(nextRun) {
			sendEmail("elliezerchristian@gmail.com", "Pemberitahuan!!!!", "Elli ganteng kga ada obat <3<3<3<3.")

			startTime = time.Now()
		}

		time.Sleep(time.Second)
	}
}
