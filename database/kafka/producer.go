package kafka

import (
	"encoding/json"
	"fmt"
	"user-service/database"
	"user-service/domain"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

const (
	_flashTimeout = 5000 // ms
)

type Producer struct {
	producer *kafka.Producer
	topic    string
}

func NewProducer(address, topic string) (*Producer, error) {
	cfg := &kafka.ConfigMap{
		"bootstrap.servers": address,
	}

	p, err := kafka.NewProducer(cfg)
	if err != nil {
		return nil, err
	}

	return &Producer{producer: p, topic: topic}, nil
}

func (p *Producer) Produce(message []database.ApprovedItem) error {
	msgBytes, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("не удалось сериализовать сообщение: %w", err)
	}

	kafkaMsg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &p.topic,
			Partition: kafka.PartitionAny,
		},
		Value: msgBytes,
		Key:   nil,
	}

	kafkaChan := make(chan kafka.Event)
	if err := p.producer.Produce(kafkaMsg, kafkaChan); err != nil {
		return err
	}

	kafkaEvent := <-kafkaChan

	switch event := kafkaEvent.(type) {
	case *kafka.Message:
		return nil
	case kafka.Error:
		return event
	default:
		return domain.UnknownEventType
	}
}

func (p *Producer) Close() {
	p.producer.Flush(_flashTimeout)
}
