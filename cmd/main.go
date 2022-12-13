package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"ptok/config"
	"ptok/internal/app"
	"syscall"
	"time"

	"github.com/robfig/cron/v3"
)

var version *string

func init() {
	version = flag.String("version", "1", "loader version") // 1 or 2
}

func main() {

	flag.Parse()
	log.Println("starting...")

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	customLocation, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		log.Println(err)
	}

	cronHandler := cron.New(cron.WithLocation(customLocation))

	defer cronHandler.Stop()

	cronHandler.AddFunc("*/2 * * * *", func() { // every 2 minutes
		app.Run(cfg, version)
	})

	go cronHandler.Start()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}
