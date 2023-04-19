package testmemory

import (
	"context"
	"github.com/redis/go-redis/v9"
	"fmt"
)

var ctx = context.Background()

var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	fmt.Println("Init rdb")
}

func RedisSet(key string, value string) {
	err := rdb.Set(ctx, key, value, 0).Err()
	if err != nil {
		panic(err)
	}
}

func RedisGet(key string) string {
	val, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return "empty"
	} else if err != nil {
		return "error"
		panic(err)
	} else {
		return val
	}
}

func RedisDelete(key string) {
	_, err := rdb.Del(ctx, key).Result()
	if err != nil {
		panic(err)
	}
}
