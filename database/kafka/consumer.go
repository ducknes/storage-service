package kafka

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"storage-service/tools/storagecontext"
)

const (
	_sessionTimeout     = 7000 // ms
	_autoCommitInterval = 5000
	_consumerTimeout    = -1
)

type Consumer struct {
	consumer       *kafka.Consumer
	messageHandler MessageHandler
	stop           bool
}

func NewConsumer(handler MessageHandler, address, topic, consumerGroup string) (*Consumer, error) {
	cfg := &kafka.ConfigMap{
		"bootstrap.servers":        address,
		"group.id":                 consumerGroup,
		"session.timeout.ms":       _sessionTimeout,
		"enable.auto.offset.store": false,
		"enable.auto.commit":       true,
		"auto.commit.interval.ms":  _autoCommitInterval,
	}

	consumer, err := kafka.NewConsumer(cfg)
	if err != nil {
		return nil, err
	}

	if err = consumer.Subscribe(topic, nil); err != nil {
		return nil, err
	}

	return &Consumer{consumer: consumer, messageHandler: handler}, nil
}

func (c *Consumer) Stop() error {
	c.stop = true
	return c.consumer.Close()
}

func (c *Consumer) Consume(ctx storagecontext.StorageContext) {
	for {
		if c.stop {
			break
		}

		kafkaMsg, err := c.consumer.ReadMessage(_consumerTimeout)
		if err != nil {
			ctx.Log().Error(fmt.Sprintf("не удалось прочитать сообщение из кафки: %v", err))
			continue
		}

		if kafkaMsg == nil {
			continue
		}

		if err = c.messageHandler.HandleMessage(ctx, kafkaMsg.Value); err != nil {
			ctx.Log().Error(fmt.Sprintf("не удалось обработать сообщение: %v", err))
			continue
		}

		if _, err = c.consumer.StoreMessage(kafkaMsg); err != nil {
			ctx.Log().Error(fmt.Sprintf("не удалось сохранить информацию о прочтении: %v", err))
			continue
		}
	}
}
