package handler

import (
	"fmt"
	"log/slog"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Handler struct {
}

func NewHandler() *Handler {

	return &Handler{}
}

func (h *Handler) HandleMessage(message []byte, offset kafka.Offset) error {
	slog.Info(fmt.Sprintf("Messsage from kafka with offset %d %s", offset, string(message)))
	return nil
}
