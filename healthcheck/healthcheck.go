package healthcheck

import (
	"errors"
	"fmt"
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

func (h *HealthCheck) Notify() {
	subject := fmt.Sprintf("Service %s is experiencing problems", h.Name)
	body := fmt.Sprintf("Service %s failed a health check at %s.", h.Name, time.Now().Format(layoutISO))

	for _, notifier := range h.Notifiers {
		notifier.Notify(subject, body)
	}
}

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
