package main

import (
	"fmt"
	"github.com/go-redis/redis"
	_ "github.com/go-redis/redis"
)

func get(key string, client redis.Client) {
	value, err := client.Get(key).Result()
	if err == redis.Nil {
		fmt.Printf("%s does not exist\n", key)
	} else if err != nil {
		panic(err)
	} else {
		fmt.Printf("%s : %s\n", key, value)
	}
}

func exists(key string, client redis.Client) {
	result, err := client.Exists(key).Result()
	if err != nil {
		panic(err)
	}
	if result == 1 {
		fmt.Printf("%s exist\n", key)
		return
	}
	fmt.Printf("%s does not exist\n", key)
}

func main() {
	students := map[string] int {
		"phong" : 23,
		"trung" : 22,
		"truong" : 21,
		"tu" : 20,
	}
	// new client
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	//set
	for key,value := range students {
		result, err := client.Set(key, value, 0).Result()
		if err != nil {
			panic(err)
		}
		fmt.Printf("set : %s\n", result)
	}


	//get
	fmt.Println("GET")
	get("an", *client)
	get("tu", *client)

	//del
	fmt.Println("DEL")
	err = client.Del("tu").Err()
	if err != nil {
		panic(err)
	}
	get("tu", *client)

	//exists
	fmt.Println("EXISTS")
	exists("tu", *client)
	exists("phong", *client)

	exist, err := client.Exists("phong", "tu", "truong", "trung").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Multiple keys (\"phong\", \"tu\", \"truong\", \"trung\") : ", exist)

	//hset and hget
	fmt.Println("HSET and HGET")
	_, err = client.HSet("user:0", "name", "phong").Result()
	if err != nil {
		panic(err)
	}

	hget, err := client.HGet("user:0", "name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("user:0 -> name: ", hget)

	//HExists
	fmt.Println("HExists")
	hexists, err := client.HExists("user:0", "name").Result()
	if err != nil {
		panic(err)
	}
	if hexists {
		fmt.Println("user:0 -> name exist")
	} else {
		fmt.Println("user:0 -> name does not exist")
	}


}