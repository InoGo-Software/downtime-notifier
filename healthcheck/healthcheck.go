package healthcheck

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

const layoutISO = "2006-01-02T15:04:05-0700"

type HealthCheck struct {
	Name      string
	Url       string
	Interval  string
	Timeout   uint16
	Notifiers []Notifier
}

// Make the HTTP call. Returns true if the request was successful.
func (h *HealthCheck) PerformRequest() bool {
	// Initialize client.
	client := http.Client{
		Timeout: time.Duration(h.Timeout) * time.Millisecond,
	}

	// Perform the GET request.
	resp, err := client.Get(h.Url)
	return err == nil && resp.StatusCode == 200
}

// Notify all listeners of the health check.
func (h *HealthCheck) Notify(isHealthy bool) {
	var subject string
	var body string

	if !isHealthy {
		subject = fmt.Sprintf("Service %s is experiencing problems", h.Name)
		body = fmt.Sprintf("Service %s failed a health check at %s.", h.Name, time.Now().Format(layoutISO))
	} else {
		subject = fmt.Sprintf("Service %s is healthy again", h.Name)
		body = fmt.Sprintf("Service %s back online at %s!", h.Name, time.Now().Format(layoutISO))
	}

	for _, notifier := range h.Notifiers {
		notifier.Notify(subject, body)
	}
}

// Validate that the data in the struct is correct.
func (h *HealthCheck) Validate() error {
	if h.Name == "" {
		return errors.New("name is required")
	}
	if h.Url == "" {
		return errors.New("url is required")
	}
	if h.Interval == "" {
		return errors.New("interval is required")
	}
	if h.Timeout == 0 {
		return errors.New("timeout is required")
	}
	if len(h.Notifiers) == 0 {
		return errors.New("notifiers are required")
	}
	for _, notifier := range h.Notifiers {
		if err := notifier.Validate(); err != nil {
			return err
		}
	}
	return nil
}
