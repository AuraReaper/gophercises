package consumer

import (
	"fmt"
	"time"

	"github.com/AuraReaper/kafka-go/internal/shared"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type KafkaConsumer struct {
	consumer *kafka.Consumer
	topic    string
	msgCh    chan<- string
}

func NewKafkaConsumer(msgCh chan<- string) (*KafkaConsumer, error) {
	cfg := shared.NewKakfkaComfig()

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": cfg.Host,
		"group.id":          cfg.ConsumerGroup,
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		return nil, err
	}

	err = c.SubscribeTopics([]string{cfg.Topic}, nil)

	if err != nil {
		return nil, err
	}

	consumer := &KafkaConsumer{
		consumer: c,
		topic:    cfg.Topic,
		msgCh:    msgCh,
	}

	go consumer.readMsgLoop()
	return consumer, nil
}

func (c *KafkaConsumer) readMsgLoop() {
	defer c.consumer.Close()

	for {
		msg, err := c.consumer.ReadMessage(time.Second)
		if err != nil && !err.(kafka.Error).IsTimeout() {
			// The client will automatically try to recover from all errors.
			// Timeout is not considered an error because it is raised by
			// ReadMessage in absence of messages.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			continue
		}

		payload := msg.Value
		c.msgCh <- string(payload)
	}
}
