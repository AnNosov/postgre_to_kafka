package usecase

import (
	"context"
	"encoding/json"
	"log"
	"ptok/internal/entity"
	"strconv"

	"github.com/segmentio/kafka-go"
)

type ProfileKafka struct {
	*kafka.Writer
}

func NewKFK(kfk *kafka.Writer) *ProfileKafka {
	return &ProfileKafka{kfk}
}

func (w ProfileKafka) Write(profile entity.Profile) {
	msg, err := json.Marshal(profile)
	if err != nil {
		log.Println("kafka writer: ", err)
	}
	w.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte(strconv.Itoa(profile.Id)),
		Value: []byte(msg),
	})

}

func (w ProfileKafka) WriteV2(profileChan chan entity.Profile) {
	go func() {
		for profile := range profileChan {

			msg, err := json.Marshal(profile)
			//log.Println(msg, err)
			if err != nil {
				log.Println("kafka writer: ", err)
			}
			w.WriteMessages(context.Background(), kafka.Message{
				Key:   []byte(strconv.Itoa(profile.Id)),
				Value: []byte(msg),
			})
		}
	}()
}
