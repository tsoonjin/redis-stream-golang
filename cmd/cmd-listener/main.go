package main

import (
    "log"
    "os"
    "fmt"
    "net/http"
    "github.com/joho/godotenv"
    "github.com/tsoonjin/redis-stream-golang/service/redis"
    "github.com/tsoonjin/redis-stream-golang/handler/redis"
)

type Config struct {
    RedisEndpoint string
    RedisPassword string
    Port string
}

func loadConfig() Config {
    if err := godotenv.Load(); err != nil {
        log.Println("Error loading env file")
    }
    return Config{
        RedisEndpoint: os.Getenv("REDIS_ENDPOINT"),
        RedisPassword: os.Getenv("REDIS_PASSWORD"),
        Port: os.Getenv("PORT"),
    }
}

func main() {
    log.Println("Loading Config")
    config := loadConfig()
    redisClient := &service.Redis{
        Endpoint: config.RedisEndpoint,
        Password: config.RedisPassword,
    }
    redisClient.Init()
    commandHandler := handler.CommandHandler{redisClient}
    http.HandleFunc("/command", commandHandler.Handler)
    log.Println(fmt.Sprintf("Listening for commands on port: %s", config.Port))
    http.ListenAndServe(fmt.Sprintf(":%s", config.Port), nil)
}
