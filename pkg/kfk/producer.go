package kfk

import (
	"log"
	"ptok/config"

	"github.com/segmentio/kafka-go"
)

func New(cfg config.Kafka) *kafka.Writer {

	return &kafka.Writer{
		Addr:     kafka.TCP(cfg.Host + ":" + cfg.Port),
		Topic:    cfg.Topic,
		Balancer: &kafka.LeastBytes{},
	}
}

func Close(w *kafka.Writer) error {
	if err := w.Close(); err != nil {
		log.Println("close kafka connection: ", err)
		return err
	}
	return nil
}
