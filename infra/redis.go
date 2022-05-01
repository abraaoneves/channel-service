package infra

import "github.com/go-redis/redis"

func CreateDataBaseConnection() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "Redis1234",
		DB:       0,
	})
}
