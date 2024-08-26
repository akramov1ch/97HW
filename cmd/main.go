package main

import (
	"97HW/config"
	cr "97HW/cron"
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	if err := config.LoadConfig("config/config.json"); err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	c := cron.New()

	c.AddFunc("0 9 * * *", func() {
		log.Println("Starting the daily data update and notification job")
		if err := cr.UpdateDataAndNotify(); err != nil {
			log.Printf("Error updating data and sending notifications: %v", err)
		} else {
			log.Println("Data updated and notifications sent successfully")
		}
	})

	c.Start()

	select {
	case <-time.After(time.Hour * 24 * 365): // Keeping the program alive for a year
	}
}
