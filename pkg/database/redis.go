package database

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"time"
)

var redisClient *redis.Client

func InitRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err := redisClient.Ping().Result()
	if err != nil {
		panic(err)
	}
}

func StoreRedisValue[T any](user, function string, value T) error {
	stakeInfoJSON, err := json.Marshal(value)
	if err != nil {
		return err
	}
	key := "user:" + user + "function" + ":" + function
	err = redisClient.Set(key, stakeInfoJSON, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func StoreRedisValueTimeout[T any](user, function string, value T, duration time.Duration) error {
	stakeInfoJSON, err := json.Marshal(value)
	if err != nil {
		return err
	}
	key := "user:" + user + "function" + ":" + function
	err = redisClient.Set(key, stakeInfoJSON, duration).Err()
	if err != nil {
		return err
	}
	return nil
}

func GetRedisValue[T any](user, function string) (T, error) {
	key := "user:" + user + "function" + ":" + function
	val, err := redisClient.Get(key).Result()
	if err != nil {
		return getZero[T](), err
	}
	var data T
	err = json.Unmarshal([]byte(val), &data)
	if err != nil {
		return getZero[T](), err
	}
	return data, nil
}

func GetRedisArray[T any](user, function string) ([]T, error) {
	key := "user:" + user + "function" + ":" + function
	val, err := redisClient.Get(key).Result()
	if err != nil {
		return getZeroArray[T](), err
	}
	var data []T
	err = json.Unmarshal([]byte(val), &data)
	if err != nil {
		return getZeroArray[T](), err
	}
	return data, nil
}

func DelRedisValue(user, function string) error {
	key := "user:" + user + "function" + ":" + function
	err := redisClient.Del(key).Err()
	if err != nil {
		return err
	}
	return nil
}
