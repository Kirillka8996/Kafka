package kafka

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

const (
	sessionTimeout = 7000 // ms
	noTimeout      = -1
)

type Handler interface {
	HandleMessage(message []byte, offset kafka.Offset) error
}
type Consumer struct {
	consumer *kafka.Consumer
	handler  Handler
	stop     bool
}

func NewConsumer(handler Handler, address []string, topic, consumerGroup string) (*Consumer, error) {
	// Документация конфига https://github.com/confluentinc/librdkafka/blob/master/CONFIGURATION.md
	cfg := kafka.ConfigMap{
		"bootstrap.servers":        strings.Join(address, ","),
		"group.id":                 consumerGroup,
		"session.timeout.ms":       sessionTimeout,
		"enable.auto.offset.store": false,
		"enable.auto.commit":       true,
		"auto.commit.interval.ms":  5000,
		"auto.offset.reset":        "earliest",
	}
	c, err := kafka.NewConsumer(&cfg)
	if err != nil {
		return nil, fmt.Errorf("error with consumer: %m", err)
	}

	if err := c.Subscribe(topic, nil); err != nil {
		return nil, fmt.Errorf("error with subscribe to topic: %m", err)
	}
	return &Consumer{consumer: c, handler: handler}, nil
}

func (c *Consumer) Start() {
	for {
		if c.stop {
			break
		}
		kafkaMessage, err := c.consumer.ReadMessage(noTimeout)
		if err != nil {
			slog.Error(err.Error())
		}
		if kafkaMessage == nil {
			continue
		}
		if err := c.handler.HandleMessage(kafkaMessage.Value, kafkaMessage.TopicPartition.Offset); err != nil {
			slog.Error(fmt.Sprintf("error with HandleMessage %m", err))
			continue
		}
		if _, err := c.consumer.StoreMessage(kafkaMessage); err != nil {
			slog.Error(fmt.Sprintf("error with StoreMessage %m", err))

		}

	}
}
func (c *Consumer) Stop() error {
	c.stop = true
	if _, err := c.consumer.Commit(); err != nil {
		return err
	}
	slog.Info("Commited offset")
	return c.consumer.Close()
}
