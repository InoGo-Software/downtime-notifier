package cron

import (
	"fmt"
	"github.com/InoGo-Software/downtime-notifier/config"
	"github.com/robfig/cron/v3"
	"log"
)

func StartCron(config config.Config) {
	// Create the scheduler.
	c := cron.New()

	// Register the jobs.
	log.Printf("Registering jobs")
	for _, healthCheck := range config.HealthChecks {
		_, err := c.AddFunc(healthCheck.Interval, func() {
			work(healthCheck)
		})
		if err != nil {
			fmt.Printf("error registering cron job %v: %s", healthCheck, err)
			return
		}
		log.Printf("[%s] %s %s", healthCheck.Name, healthCheck.Interval, healthCheck.Url)
		for _, notifier := range healthCheck.Notifiers {
			log.Printf(" - Notifier %s via %s", notifier.To, notifier.Type)
		}
	}

	log.Println("Scheduler running...")
	c.Run()
}
