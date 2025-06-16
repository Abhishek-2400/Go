package config

// Config holds Kafka configuration
type Config struct {
	Brokers []string
	Topic   string
	Group   string
}

// GetConfig returns the default Kafka configuration
func GetConfig() Config {
	return Config{
		Brokers: []string{"localhost:9092"},
		Topic:   "test-topic",
		Group:   "test-group",
	}
}
