# InoGo Notifier

> Easily get notified when your service is down.

# Setup

1. `git clone git@github.com:InoGo-Software/downtime-notifier.git`
1. `cp config.example.yaml config.yaml`
1. Write your config in `config.yaml`
1. Run & build
    1. with go: `go build . && ./downtime-notifier`
    1. with docker: `docker build -t downtime-notifier . && docker run -e SENDGRID_API_KEY= -e FCM_API_KEY= downtime-notifier`

# Config

## Sendgrid

In order to use sendgrid you MUST set the environment variable `SENDGRID_API_KEY` to your own sendgrid api key.

## FCM

In order to use FCM you MUST set the environment variable `FCM_API_KEY`. This is the 'server key' in firebase -> project overview settings -> cloud messaging.

Example config:
```yaml
healthchecks:
  - name: "Google"
    url: "https://google.com"
    interval: "* * * * *"  # Cron config. See https://godoc.org/github.com/robfig/cron#hdr-CRON_Expression_Format
    timeout: 300
    notifiers:
      - to: info@example.com
        type: sendgrid  # Valid types: 'sendgrid'
```
