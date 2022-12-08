package main

import (
	"log"
	"ptok/config"
	"ptok/internal/app"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	customLocation, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		log.Println(err)
	}

	cronHandler := cron.New(cron.WithLocation(customLocation))

	cronHandler.AddFunc("0 */5 * ? * *", func() { // every 5 minutes
		log.Println("cron start running...")
		app.Run(cfg)
		log.Println("cron stop running...")
	})
}
