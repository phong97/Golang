package cache

import (
	"log"
	"sync"
	"time"

	"github.com/go-redis/redis"
)

type singletonRedisClient struct {
	redis.Client
}

type redisConfig struct {
	Addr     string
	Password string
	DB       int
}

var config = redisConfig{
	"localhost:6379",
	"",
	0,
}

var instance *singletonRedisClient
var once sync.Once

func setupRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.DB,
	})
}

func GetInstance() *singletonRedisClient {
	once.Do(func() {
		instance = &singletonRedisClient{*setupRedis()}
	})
	return instance
}

func Set(key, value string, time time.Duration, client redis.Client) bool {
	err := client.Set(key, value, time).Err()
	if err != nil {
		return false
	}
	return true
}

func Get(username string, client redis.Client) string {
	token, err := client.Get(username).Result()
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return token
}
