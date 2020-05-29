package services

import (
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"log"
	"os"
)

const (
	SENDGRID_API_KEY = "SENDGRID_API_KEY"
	SUBJECT          = "InoGo Notifier"
	FROM             = "noreply@inogo.nl"
)

var (
	SendgridService sendgridServiceInterface = &sendgridService{}
)

type sendgridService struct{}

type sendgridServiceInterface interface {
	Notify(subject string, body string, receiver string)
}

func (s sendgridService) Notify(subject string, body string, receiver string) {
	from := mail.NewEmail(SUBJECT, FROM)
	to := mail.NewEmail("", receiver)
	message := mail.NewSingleEmail(from, subject, to, body, body)
	client := sendgrid.NewSendClient(os.Getenv(SENDGRID_API_KEY))

	_, err := client.Send(message)

	if err != nil {
		log.Printf("error while sending mail to %s: %v", receiver, err)
		return
	}
	log.Printf("successfully notified %s via mail", receiver)
}
