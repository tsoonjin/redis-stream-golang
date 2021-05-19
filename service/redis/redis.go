package service

import (
    "log"
    "fmt"
    "github.com/go-redis/redis"
)

type Redis struct {
    Endpoint string
    Password string
    client *redis.Client
}

func (r *Redis) Init() {
    log.Println(fmt.Sprintf("Connecting to redis server at: %s", r.Endpoint))
	redisClient := redis.NewClient(&redis.Options{
		Addr: r.Endpoint,
        Password: r.Password,
        DB: 0,
	})
    _, err := redisClient.Ping().Result()
    if err != nil {
        log.Fatal("Unable to connect to Redis", err)
    }
    log.Println("Connected to Redis server")
    r.client = redisClient
}
