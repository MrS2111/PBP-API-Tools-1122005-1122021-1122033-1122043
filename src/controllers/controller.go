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

func SendEmail() error {
	m := gomail.NewMessage()
	m.SetHeader("From", "jasonjeyys@gmail.com")
	m.SetHeader("To", "if-22021@students.ithb.ac.id")
	m.SetHeader("Subject", "This is your reminder email!")
	m.SetBody("text/plain", "This email is automated. Here is your reminder email. Have a great day!")

	d := gomail.NewDialer("smtp.gmail.com", 587, "jasonjeyys@gmail.com", "barr gggs fxck gkkv")

	err := d.DialAndSend(m)
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	return nil
}

func TaskScheduler(ctx context.Context, stop chan struct{}) {
	go func() {
		startTime := time.Now()
		interval := time.Minute * 5

		for {
			select {
			case <-ctx.Done():
				fmt.Println("Task scheduler stopped.")
				return
			default:
				now := time.Now()
				nextRun := startTime.Add(interval)

				if now.After(nextRun) || now.Equal(nextRun) {
					err := SendEmail()
					if err != nil {
						fmt.Printf("Error sending email: %v\n", err)
					}

					startTime = time.Now()
				}

				time.Sleep(time.Second)
			}
		}
	}()
}

func Caching(ctx context.Context) {
	go func() {
		client := redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})

		_, err := client.Ping(ctx).Result()
		if err != nil {
			log.Printf("Failed to ping Redis server: %v\n", err)
			return
		}

		type Person struct {
			ID         string `json:"id"`
			Name       string `json:"name"`
			Age        int    `json:"age"`
			Occupation string `json:"occupation"`
		}

		elliotID := "elliot"
		elliotKey := fmt.Sprintf("person:%s", elliotID)
		val, err := client.Get(ctx, elliotKey).Result()
		if err == nil {
			log.Println("Value retrieved from Redis cache:", val)
			var person Person
			if err := json.Unmarshal([]byte(val), &person); err != nil {
				log.Printf("Failed to unmarshal cached data: %v\n", err)
			}
			log.Println("Cached person:", person)
			err := SendEmail()
			if err != nil {
				log.Printf("Error sending email: %v\n", err)
			}
		} else if err == redis.Nil {
			log.Println("Key does not exist in Redis cache, fetching data...")
			jsonString, err := json.Marshal(Person{
				ID:         elliotID,
				Name:       "Elliot",
				Age:        25,
				Occupation: "Software Developer",
			})
			if err != nil {
				log.Printf("Failed to marshal the person struct: %v\n", err)
				return
			}

			err = client.Set(ctx, elliotKey, jsonString, 0).Err()
			if err != nil {
				log.Printf("Failed to set value in the Redis instance: %v\n", err)
				return
			}

			log.Println("Value stored in Redis cache.")
			err = SendEmail()
			if err != nil {
				log.Printf("Error sending email: %v\n", err)
			}
		} else {
			log.Printf("Failed to get value from the Redis instance: %v\n", err)
		}
	}()
}