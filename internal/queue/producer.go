package queue

import (
	"context"
	"encoding/json"
	"log"

	"github.com/redis/go-redis/v9"
	"github.com/rishavbagri/go-job-queue/internal/model"
)

type Producer struct {
	client *redis.Client
}

func NewProducer(client *redis.Client) *Producer{
	return &Producer{
		client: client,	
	}
}

func (p Producer) Enqueue (ctx context.Context, job model.Job) error{

	data, err := json.Marshal(job)

	if err!=nil {
		return err
	}
	log.Printf("a job posted lol")
	return p.client.LPush(ctx, "key",data).Err()
}