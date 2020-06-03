package cron

import (
	"fmt"
	"github.com/InoGo-Software/downtime-notifier/config"
	"github.com/InoGo-Software/downtime-notifier/healthcheck"
	"github.com/robfig/cron/v3"
	"log"
)

func register(c *cron.Cron, ht healthcheck.HealthCheck) {
	_, err := c.AddFunc(ht.Interval, func() {
		work(ht)
	})
	if err != nil {
		fmt.Printf("error registering cron job %v: %s", ht, err)
		return
	}
	log.Printf("[%s] %s %s", ht.Name, ht.Interval, ht.Url)
	for _, notifier := range ht.Notifiers {
		log.Printf(" - Notifier %s via %s", notifier.To, notifier.Type)
	}
}

func StartCron(config config.Config) {
	// Create the scheduler.
	c := cron.New()

	// Register the jobs.
	log.Printf("Registering jobs")
	for _, healthCheck := range config.HealthChecks {
		register(c, healthCheck)
	}

	log.Println("Scheduler running...")
	c.Run()
}
