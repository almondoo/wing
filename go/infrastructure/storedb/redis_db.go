package storedb

import (
	"context"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisService struct {
	client *redis.Client
}

func NewRedisDB() *RedisService {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	return &RedisService{client: rdb}
}

var ctx = context.Background()

//- 分単位
//- データを登録
func (rs *RedisService) Set(key string, value interface{}, ttl time.Duration) error {
	err := rs.client.Set(ctx, key, value, ttl).Err()
	if err != nil {
		return err
	}
	return nil
}

//- データを取得
func (rs *RedisService) Get(key string) (string, error) {
	val, err := rs.client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

//- データを削除
func (rs *RedisService) Delete(key string) error {
	err := rs.client.Del(ctx, key).Err()
	if err != nil {
		return err
	}
	return nil
}
