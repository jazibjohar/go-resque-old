package main

import (
	"context"

	"github.com/go-redis/redis/v9"            // Redis client from go-redis package
	"github.com/skaurus/go-resque"            // Import this package
	_ "github.com/skaurus/go-resque/redis.v9" // Use go-redis v9
)

func main() {
	var err error

	// Create new Redis client to use for enqueuing
	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		DB:   0,
	})
	// Create enqueuer instance
	enqueuer := resque.NewRedisEnqueuer(context.Background(), "redis.v9", client, "resque:")

	// Enqueue the job into the "go" queue with appropriate client
	_, err = enqueuer.Enqueue(context.Background(), "go", "Demo::Job")
	if err != nil {
		panic(err)
	}

	// Enqueue into the "default" queue with passing one parameter to the Demo::Job.perform
	_, err = enqueuer.Enqueue(context.Background(), "default", "Demo::Job", 1)
	if err != nil {
		panic(err)
	}

	// Enqueue into the "extra" queue with passing multiple
	// parameters to the Demo::Job.perform so it will fail
	_, err = enqueuer.Enqueue(context.Background(), "extra", "Demo::Job", 1, 2, "woot")
	if err != nil {
		panic(err)
	}
}
