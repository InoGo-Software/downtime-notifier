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
	FCM      NotifierType = "fcm"
)

type Notifier struct {
	Type NotifierType
	To   string
}

func (n *Notifier) Notify(subject string, body string) {
	if n.Type == SENDGRID {
		go services.SendgridService.Notify(subject, body, n.To)
	}
	if n.Type == FCM {
		go services.FcmService.Notify(body, n.To)
	}
}

func (n *Notifier) Validate() error {
	if n.To == "" {
		return errors.New("notifier.to is required")
	}
	if n.Type != SENDGRID && n.Type != FCM {
		return errors.New(fmt.Sprintf("%s is not a valid notifier type", n.Type))
	}
	if n.Type == SENDGRID && os.Getenv(services.SENDGRID_API_KEY) == "" {
		return errors.New(fmt.Sprintf("sendgrid type is used but %s is not set", services.SENDGRID_API_KEY))
	}
	if n.Type == FCM && os.Getenv(services.FCM_API_KEY) == "" {
		return errors.New(fmt.Sprintf("fcm type is used but %s is not set", services.FCM_API_KEY))
	}
	return nil
}
