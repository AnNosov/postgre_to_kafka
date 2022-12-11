package app

import (
	"log"
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
		if err := profileUCase.TransportData(); err != nil {
			log.Println(err)
		}
	case "2":
		if err := profileUCase.TransportDataV2(); err != nil {
			log.Println(err)
		}
	default:
		log.Println("incorrect version")
		return
	}

}
