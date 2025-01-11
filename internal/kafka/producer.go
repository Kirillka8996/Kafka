package kafka

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

const (
	flushTimeout = 5000 // ms
)

var errUnKnownType = errors.New("unknown event type")

type Producer struct {
	producer *kafka.Producer
}

func NewProducer(address []string) (*Producer, error) {
	conf := kafka.ConfigMap{
		"bootstrap.servers": strings.Join(address, ","),
	}
	p, err := kafka.NewProducer(&conf)
	if err != nil {
		return nil, fmt.Errorf("error with producer: %m", err)
	}

	return &Producer{producer: p}, nil
}
func (p *Producer) Produce(message, topic string) error {
	kafkaMessage := &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value:         []byte(message),
		Key:           nil,
		Timestamp:     time.Time{},
		TimestampType: 0,
		Opaque:        nil,
		Headers:       nil,
	}
	kafkaCh := make(chan kafka.Event)
	if err := p.producer.Produce(kafkaMessage, kafkaCh); err != nil {
		return err
	}
	e := <-kafkaCh

	switch event := e.(type) {
	case *kafka.Message:
		return nil
	case *kafka.Error:
		return event
	default:
		return errUnKnownType
	}

}
func (p *Producer) Close() {
	p.producer.Flush(flushTimeout)
	p.producer.Close()
}
