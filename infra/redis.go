package infra

import (
	"github.com/go-redis/redis"
)

func CreateDataBaseConnection() *redis.Client {
	configuration := configure()
	return redis.NewClient(&redis.Options{
		Addr:     configuration.DB.Address,
		Password: configuration.DB.Password,
		DB:       0,
	})
}
