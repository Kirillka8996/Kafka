package main

import (
	"kafka/internal/handler"
	"kafka/internal/kafka"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const (
	topic         = "my-topic"
	consumerGroup = "my-consumer-group"
)

var address = []string{"localhost:9091", "localhost:9092", "localhost:9093"}

func main() {
	h := handler.NewHandler()
	c, err := kafka.NewConsumer(h, address, topic, consumerGroup)
	if err != nil {
		log.Fatal(err)

	}
	go func() {
		c.Start()
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan
	log.Fatal(c.Stop())
}
