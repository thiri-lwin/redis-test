package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {
	fmt.Println("Go Redis Tutorial")

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	ctx := context.Background()
	pong, err := client.Ping(ctx).Result()
	fmt.Println(pong, err)

	err = client.Set(ctx, "name", "Elliot", 1*time.Minute).Err()

	if err != nil {
		fmt.Println(err)
	}

	res, err := client.Get(ctx, "name").Result()
	fmt.Println(res, err)
}
