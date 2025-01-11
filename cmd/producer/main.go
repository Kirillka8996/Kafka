package main

import (
	"fmt"
	"log"
	"log/slog"

	k "kafka/internal/kafka"
)

const (
	topic = "my-topic"
)

var address = []string{"localhost:9091", "localhost:9092", "localhost:9093"}

func main() {
	p, err := k.NewProducer(address)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < 100; i++ {
		msg := fmt.Sprintf("Kafka msg %d", i)
		if err := p.Produce(msg, topic); err != nil {
			slog.Error(err.Error())
		}
	}

}
