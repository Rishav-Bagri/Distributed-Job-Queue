package queue

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"

	"github.com/rishavbagri/go-job-queue/internal/model"
)

type Consumer struct {
	client *redis.Client
}

func NewConsumer(client *redis.Client) *Consumer {
	return &Consumer{
		client: client,
	}
}

func (c *Consumer) Dequeue(ctx context.Context) (*model.Job, error) {

	result, err := c.client.BRPop(ctx, 0, "key").Result()
	if err != nil {
		return nil, err
	}

	var job model.Job

	err = json.Unmarshal([]byte(result[1]), &job)
	if err != nil {
		return nil, err
	}

	return &job, nil
}