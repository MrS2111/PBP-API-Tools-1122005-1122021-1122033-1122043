package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
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

func Caching() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	ping, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println(ping)

	type Person struct {
		ID         string `json:"id"`
		Name       string `json:"name"`
		Age        int    `json:"age"`
		Occupation string `json:"occupation"`
	}

	elliotID := "elliot"
	elliotKey := fmt.Sprintf("person:%s", elliotID)
	val, err := client.Get(context.Background(), elliotKey).Result()
	if err == nil {
		log.Println("Value retrieved from Redis cache:", val)
		var person Person
		if err := json.Unmarshal([]byte(val), &person); err != nil {
			log.Fatal("Failed to unmarshal cached data:", err)
		}
		log.Println("Cached person:", person)
		sendEmail()
	} else if err == redis.Nil {
		log.Println("Key does not exist in Redis cache, fetching data...")
		jsonString, err := json.Marshal(Person{
			ID:         elliotID,
			Name:       "Elliot",
			Age:        25,
			Occupation: "Software Developer",
		})
		if err != nil {
			log.Fatal("Failed to marshal the person struct:", err)
		}

		err = client.Set(context.Background(), elliotKey, jsonString, 0).Err()
		if err != nil {
			log.Fatal("Failed to set value in the Redis instance:", err)
		}

		log.Println("Value stored in Redis cache.")

		sendEmail()
	} else {
		log.Fatal("Failed to get value from the Redis instance:", err)
	}
}
