package app

import (
	"log"
	"os"
	"ptok/config"
	"ptok/internal/usecase"
	"ptok/pkg/kfk"
	"ptok/pkg/postgres"
)

func Run(cfg *config.Config, version *string) {

	pg, err := postgres.New(&cfg.Postgres)
	if err != nil {
		log.Fatal(err)
	}
	defer pg.Close()

	kfk := kfk.New(cfg.Kafka)
	profileUCase := usecase.New(*usecase.NewKFK(kfk), *usecase.NewPG(pg))

	switch *version {
	case "1":
		log.Println("cron: app v.1 start running ...")
		if err := profileUCase.TransportData(); err != nil {
			log.Println(err)
		}
	case "2":
		log.Println("cron: app v.2 start running ...")
		if err := profileUCase.TransportDataV2(); err != nil {
			log.Println(err)
		}
	default:
		log.Println("incorrect version")
		os.Exit(1)
	}

}
