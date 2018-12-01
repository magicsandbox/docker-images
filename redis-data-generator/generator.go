package main

import (
	"github.com/go-redis/redis"
	"github.com/Pallinder/go-randomdata"
	"os"
	"time"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
	})

	for {
		client.Set(randomdata.Email(), randomdata.Address(), 5 * time.Minute)
		<- time.After(2 * time.Second)
	}
}
