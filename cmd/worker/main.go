package main

import (
	"context"
	"log"

	"github.com/rishavbagri/go-job-queue/internal/queue"
)

func main() {

	client := queue.NewRedisClient()

	consumer := queue.NewConsumer(client)

	ctx := context.Background()

	for {
		job, err := consumer.Dequeue(ctx)
		if err != nil {
			log.Println(err)
			continue
		}

		
		log.Printf("Received Job: %+v\n", job)
	}
}