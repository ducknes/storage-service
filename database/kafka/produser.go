package kafka

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"storage-service/database"
	"storage-service/tools/customerror"
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

	producer, err := kafka.NewProducer(cfg)
	if err != nil {
		return nil, err
	}

	return &Producer{producer: producer, topic: topic}, nil
}

func (p *Producer) Close() {
	p.producer.Flush(_flashTimeout)
}

func (p *Producer) Produce(message []database.ApproveMessage) error {
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
	}

	kafkaChan := make(chan kafka.Event)
	if err = p.producer.Produce(kafkaMsg, kafkaChan); err != nil {
		return err
	}

	kafkaEvent := <-kafkaChan

	switch event := kafkaEvent.(type) {
	case *kafka.Message:
		return nil
	case kafka.Error:
		return event
	default:
		return customerror.UnknownEventType
	}
}
