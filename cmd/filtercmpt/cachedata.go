package main

import (
	redis "github.com/go-redis/redis/v7"
)

var client *redis.Client

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "192.168.190.150:6379",
		Password: "",
		DB:       0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		panic("Redis Connect Fail!")
	}
}

func getCache(key string) (string, error) {
	val, err := client.Get(key).Result()
	if err != nil; err.Error() == "redis: nil" {
		return nil, nil
	}
	return val, err
}

func setCache(key string, value interface{}) (string, error) {
	val, err := client.Set(key, value, 0).Result()
	return val, err
}
