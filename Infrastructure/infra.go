package infrastrucutre

import (
	"context"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

func RedisConnect() *redis.Client {
	
	ctx := context.Background()
	redisAddress  := os.Getenv("REDIS_ADDRESS")
	redisUsername := os.Getenv("REDIS_USERNAME")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Username: redisUsername,
		Password: redisPassword,
		DB:       0,
	})

	rdb.Set(ctx, "foo", "bar", 0)
	result, err := rdb.Get(ctx, "foo").Result()

	if err != nil {
		panic(err)
	}

	fmt.Println(result) 
	return rdb
}