package app

import (
	"log"
	"ptok/config"
	"ptok/internal/usecase"
	"ptok/pkg/kfk"
	"ptok/pkg/postgres"
)

func Run(cfg *config.Config) {

	pg, err := postgres.New(&cfg.Postgres)
	if err != nil {
		log.Fatal(err)
	}
	defer pg.Close()

	kfk := kfk.New(cfg.Kafka)

	profileUCase := usecase.New(*usecase.NewKFK(kfk), *usecase.NewPG(pg))

	if err := profileUCase.TransportData(); err != nil {
		log.Println(err)
	}

}
