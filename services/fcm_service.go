package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	FCM_API_KEY = "FCM_API_KEY"
	fcmPostUrl  = "https://fcm.googleapis.com/fcm/send"
	contentType = "application/json"
)

type fcmBody struct {
	Body string `json:"body"`
}

type fcmRequest struct {
	To           string  `json:"to"`
	Notification fcmBody `json:"notification"`
	Data         fcmBody `json:"data"`
}

var (
	FcmService fcmServiceInterface = &fcmService{}
)

type fcmService struct{}

type fcmServiceInterface interface {
	Notify(body string, receiver string)
}

func (s fcmService) Notify(body string, receiver string) {
	// Build the body data.
	data, err := json.Marshal(fcmRequest{
		To:           receiver,
		Notification: fcmBody{Body: body},
		Data:         fcmBody{Body: body},
	})
	if err != nil {
		fmt.Printf("unable to serialize fcm body: %v", err)
		return
	}

	log.Printf("serialized data: %s", data)
	// Create the request.
	req, err := http.NewRequest("POST", fcmPostUrl, bytes.NewReader(data))
	if err != nil {
		fmt.Printf("unable to build fcm request: %v", err)
		return
	}

	// Get the FCM_API_KEY.
	fcmApiKey := os.Getenv(FCM_API_KEY)
	req.Header.Set("Authorization", fmt.Sprintf("key=%s", fcmApiKey))
	req.Header.Set("Content-Type", contentType)

	// Perform the request.
	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		fmt.Printf("error while sending notification to fcm: %v", err)
		return
	}
}
