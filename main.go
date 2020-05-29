package main

import (
	"github.com/InoGo-Software/downtime-notifier/config"
	"github.com/InoGo-Software/downtime-notifier/cron"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		return
	}

	cron.StartCron(*cfg)
}
