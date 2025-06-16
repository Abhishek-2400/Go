package main

import (
	"log"
	// replace with actual module path
	"github.com/Abhishek-2400/kafka/config"
	"github.com/IBM/sarama"
)

func main() {
	cfg := config.GetConfig()

	config := sarama.NewConfig()

	consumer, err := sarama.NewConsumer(cfg.Brokers, config)
	if err != nil {
		log.Fatalf("Failed to start Kafka consumer: %v", err)
	}
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition(cfg.Topic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Failed to start partition consumer: %v", err)
	}
	defer partitionConsumer.Close()

	log.Println("Listening for messages...")

	for message := range partitionConsumer.Messages() {
		log.Printf("Received message: %s", string(message.Value))
	}
}
