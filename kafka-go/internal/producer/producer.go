package producer

import (
	"fmt"

	"github.com/AuraReaper/kafka-go/internal/shared"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

// type User struct {
// 	Name           string `json:"name"`
// 	FavoriteNumber int    `json:"favorite_number"`
// }

type KafkaProducer struct {
	producer *kafka.Producer
	topic    string
}

func NewKafkaProducer(topic string) *KafkaProducer {
	cfg := shared.NewKakfkaComfig()
	if cfg.Topic == "" {
		topic = cfg.Topic
	}

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": cfg.Host})
	if err != nil {
		panic(err)
	}

	defer p.Close()

	// Delivery report handler for produced messages
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	return &KafkaProducer{
		producer: p,
		topic:    topic,
	}
}

func (p *KafkaProducer) Produce(msg string) {
	p.producer.Produce(
		&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &p.topic, Partition: kafka.PartitionAny},
			Value:          []byte(msg),
		},
		nil,
	)
}
