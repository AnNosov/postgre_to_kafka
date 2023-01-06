package usecase_test

import (
	"log"
	"ptok/config"
	"ptok/internal/usecase"
	"ptok/pkg/kfk"
	"ptok/pkg/postgres"
	"testing"
)

// go test -bench=. -benchmem

func BenchmarkTransportData(t *testing.B) {
	var cfg, err = config.NewConfig()
	pg, err := postgres.New(&cfg.Postgres)
	if err != nil {
		log.Fatal(err)
	}
	defer pg.Close()

	kfk := kfk.New(cfg.Kafka)
	profileUCase := usecase.New(*usecase.NewKFK(kfk), *usecase.NewPG(pg))

	for i := 0; i <= t.N; i++ {
		profileUCase.TransportData()
	}
}

func BenchmarkTransportDataV2(t *testing.B) {
	var cfg, err = config.NewConfig()
	pg, err := postgres.New(&cfg.Postgres)
	if err != nil {
		log.Fatal(err)
	}
	defer pg.Close()

	kfk := kfk.New(cfg.Kafka)
	profileUCase := usecase.New(*usecase.NewKFK(kfk), *usecase.NewPG(pg))

	for i := 0; i <= t.N; i++ {
		profileUCase.TransportDataV2()
	}
}
