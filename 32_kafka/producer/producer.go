package main

import (
	"log"

	"github.com/Abhishek-2400/kafka/config"

	"github.com/IBM/sarama"
)

func main() {
	cfg := config.GetConfig()
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Retry.Max = 5

	producer, err := sarama.NewSyncProducer(cfg.Brokers, config)
	if err != nil {
		log.Fatalf("Error creating producer: %v", err)
	}
	defer producer.Close()

	message := &sarama.ProducerMessage{
		Topic: cfg.Topic,
		Value: sarama.StringEncoder("Hello Kafka!"),
	}

	partition, offset, err := producer.SendMessage(message)
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}

	log.Printf("Message sent to partition %d at offset %d\n", partition, offset)
}
