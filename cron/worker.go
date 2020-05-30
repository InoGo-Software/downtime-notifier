package cron

import (
	"github.com/InoGo-Software/downtime-notifier/healthcheck"
	"log"
	"net/http"
	"time"
)

func work(healthCheck healthcheck.HealthCheck) {
	// Initialize client.
	client := http.Client{
		Timeout: time.Duration(healthCheck.Timeout) * time.Millisecond,
	}

	// Perform the GET request.
	resp, err := client.Get(healthCheck.Url)
	if err != nil || resp.StatusCode != 200 {
		log.Printf("Failed to make successful GET request to %s", healthCheck.Url)
		healthCheck.Notify()
		return
	}

	// Print result.
	log.Printf("[%s]: %s returned %d\n", healthCheck.Name, healthCheck.Url, resp.StatusCode)
}
