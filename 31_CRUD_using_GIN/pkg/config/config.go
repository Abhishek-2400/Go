package config

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var redisClient *redis.Client
var ctx = context.Background()

func Connect() { //establish connection with db
	dsn := "host=localhost user=postgres password=arun196821803001 dbname=stocks port=5432 sslmode=disable"
	dbInstance, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = dbInstance
}

func ConnectRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379", // or from os.Getenv("REDIS_ADDR")
		Password: "",               // no password set
		DB:       0,                // use default DB
	})
	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
	log.Println("Connected to Redis 🎯")
}

func GetDB() *gorm.DB { //transfer the db string
	return db
}

func GetRedisClient() *redis.Client {
	return redisClient
}

func GetContext() context.Context {
	return ctx
}
