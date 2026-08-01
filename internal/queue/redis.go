package queue

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var ctx= context.Background()

func NewRedisClient() *redis.Client{
	client:= redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	_,err:= client.Ping(ctx).Result()
	if err!=nil {
		log.Fatal(err)
	}

	log.Println("Connected to Redis 🚀")

	return client

}