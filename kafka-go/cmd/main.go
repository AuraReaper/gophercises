package main

import (
	"fmt"
	"time"

	"github.com/AuraReaper/kafka-go/internal/consumer"
	"github.com/AuraReaper/kafka-go/internal/producer"
)

type Server struct {
	producer *producer.KafkaProducer
	consumer *consumer.KafkaConsumer
	msgCh    chan string
}

func NewServer() *Server {
	msgCh := make(chan string, 64)
	c, err := consumer.NewKafkaConsumer(msgCh)
	if err != nil {
		panic(err)
	}
	return &Server{
		producer: producer.NewKafkaProducer(""),
		consumer: c,
		msgCh:    msgCh,
	}
}

func (s *Server) produceMsg() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	id := 0

	for t := range ticker.C {
		msg := fmt.Sprintf("hello from kafka, msgID=  %d, ts = %s", id, t.Format("15:20:20"))
		s.producer.Produce(msg)
		id++
	}
}

func (s *Server) handleMsg(msg string) {
	// db ops
	fmt.Printf("received msg = %s\n", msg)
}

func main() {
	s := NewServer()
	go s.produceMsg()

	for msg := range s.msgCh {
		go s.handleMsg(msg)
	}
}
