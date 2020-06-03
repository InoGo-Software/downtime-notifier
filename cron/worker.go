package cron

import (
	"github.com/InoGo-Software/downtime-notifier/healthcheck"
	"github.com/InoGo-Software/downtime-notifier/services"
	"log"
	"time"
)

const delayForNextCall = 5 * time.Second

func isHealthCheckPassing(healthCheck *healthcheck.HealthCheck, isUnhealthy bool) bool {
	// Make the first request.
	if healthCheck.PerformRequest() {
		services.FileSystemService.Save(healthCheck.Name, false)

		if isUnhealthy {
			healthCheck.Notify(true)
		}

		log.Printf("[%s]: request to %s was successful", healthCheck.Name, healthCheck.Url)
		return true
	}
	log.Printf("[%s]: Failed to make successful GET request to %s", healthCheck.Name, healthCheck.Url)
	return false
}

func work(healthCheck healthcheck.HealthCheck) {
	isUnhealthy := services.FileSystemService.IsFailing(healthCheck.Name)

	if isHealthCheckPassing(&healthCheck, isUnhealthy) {
		return
	}

	// Delay before making next call.
	time.Sleep(delayForNextCall)

	if isHealthCheckPassing(&healthCheck, isUnhealthy) {
		return
	}

	// Check if the service was already failing.
	if isUnhealthy {
		log.Printf("[%s]: Already in failing state. Skipping notify", healthCheck.Name)
		return
	}

	healthCheck.Notify(false)

	// Save the failing state.
	services.FileSystemService.Save(healthCheck.Name, true)
}
