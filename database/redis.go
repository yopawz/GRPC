package database

import (
	"github.com/go-redis/redis"
)

type RedisDatabase struct {
	Client *redis.Client
}

func CreateRedisDatabase() *RedisDatabase {
	client := redis.NewClient(
		&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})
	_, err := client.Ping().Result()

	HandleError("Cannot Connect to Redis Database", err)

	return &RedisDatabase{
		Client: client,
	}
}

func (redis *RedisDatabase) Set(key string, value string) string {
	_, err := redis.Client.Set(key, value, 0).Result()

	HandleError("Failed to set the key", err)
	return key
}

func (redis *RedisDatabase) Get(key string) string {
	value, err := redis.Client.Get(key).Result()

	HandleError("Failed to get the value from the key: %v", err)

	// if err != nil && err != rd.Nil {
	// log.Fatalf("Failed to get the value from the key: %v", err)
	// }
	return value
}

func (redis *RedisDatabase) Delete(key string) string {
	_, err := redis.Client.Del(key).Result()

	HandleError("Failed to delete data from the key", err)

	return key
}
