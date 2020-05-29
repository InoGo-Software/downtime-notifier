# InoGo Notifier

> Easily get notified when your service is down.

# Setup

1. Write your config in `config.yaml`
1. Run & build
    1. with go: `go build . && ./downtime-notifier`
    1. with docker: `docker build -t downtime-notifier . && docker run -e SENDGRID_API_KEY=xxx downtime-notifier`

# Config

## Sendgrid

In order to use sendgrid you MUST set the environment variable `SENDGRID_API_KEY` to your own sendgrid api key.

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