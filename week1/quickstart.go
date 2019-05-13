package quickstart

import "github.com/go-redis/redis"

func ExampleNewClient() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0
	})

	pong,err := client.Ping().Result()
	fmt.Println(pong, err)
}

