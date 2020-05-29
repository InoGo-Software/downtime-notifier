package healthcheck

import (
	"errors"
	"fmt"
	"github.com/InoGo-Software/downtime-notifier/services"
	"os"
)

type NotifierType string

const (
	SENDGRID NotifierType = "sendgrid"
)

type Notifier struct {
	Type NotifierType
	To   string
}

func (n *Notifier) Notify(subject string, body string) {
	if n.Type == SENDGRID {
		services.SendgridService.Notify(subject, body, n.To)
	}
}

func (n *Notifier) Validate() error {
	if n.To == "" {
		return errors.New("notifier.to is required")
	}
	if n.Type != SENDGRID {
		return errors.New("notifier.type is not a valid type")
	}
	if n.Type == SENDGRID && os.Getenv(services.SENDGRID_API_KEY) == "" {
		return errors.New(fmt.Sprintf("sendgrid type is used but %s is not set", services.SENDGRID_API_KEY))
	}
	return nil
}
