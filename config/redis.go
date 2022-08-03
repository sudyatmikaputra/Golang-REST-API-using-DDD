package config

import (
	"strconv"
	"sync"

	"github.com/go-redis/redis"
)

var redisClient *redis.Client
var redisSingleton sync.Once

func Redis() *redis.Client {
	redisSingleton.Do(func() {
		var err error

		db, err := getRedisDB()

		if nil != err {
			GetLogger().Fatal("Redis DB should be Integer", err.Error())
		}

		redisClient = redis.NewClient(&redis.Options{
			Addr:     getRedisAddr(),
			Password: getRedisPwd(),
			DB:       db,
		})

		_, err = redisClient.Ping().Result()
		if nil != err {
			GetLogger().Fatal("Failed to create Redis DB Connection", err.Error())
		}
	})

	return redisClient
}

func getRedisDB() (int, error) {
	return strconv.Atoi(GetValue(REDIS_DB))
}

func getRedisAddr() string {
	return GetValue(REDIS_ADDRESS)
}

func getRedisPwd() string {
	return GetValue(REDIS_PASSWORD)
}
