package usecase_test

import (
	"log"
	"ptok/config"
	"ptok/internal/usecase"
	"ptok/pkg/kfk"
	"ptok/pkg/postgres"
	"testing"
)

var cfg, _ = config.NewConfig()

func BenchmarkTransportData(t *testing.B) {
	pg, err := postgres.New(&cfg.Postgres)
	if err != nil {
		log.Println(err)
	}
	defer pg.Close()

	kfk := kfk.New(cfg.Kafka)
	profileUCase := usecase.New(*usecase.NewKFK(kfk), *usecase.NewPG(pg))

	for i := 0; i <= t.N; i++ {
		profileUCase.TransportData()
	}
}

func BenchmarkTransportDataV2(t *testing.B) {
	pg, err := postgres.New(&cfg.Postgres)
	if err != nil {
		log.Println(err)
	}
	defer pg.Close()

	kfk := kfk.New(cfg.Kafka)
	profileUCase := usecase.New(*usecase.NewKFK(kfk), *usecase.NewPG(pg))

	for i := 0; i <= t.N; i++ {
		profileUCase.TransportDataV2()
	}
}

// need check/repair path config before running benchmarks
// go test -bench=. -benchmem

/*
BenchmarkTransportData-4               1        14065277076 ns/op         607008 B/op       1594 allocs/op
BenchmarkTransportDataV2-4          1764            724593 ns/op           10047 B/op        144 allocs/op
PASS
ok      ptok/internal/usecase   16.251s
*/
