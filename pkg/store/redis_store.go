package store

import (
	"context"
	"errors"

	"github.com/go-redis/redis/v8"
)

type RedisStore struct{
	client *redis.Client
}

func NewRedisStore(addr, password string) *RedisStore {
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
		Password: password,
		DB: 0,
	})

	return &RedisStore{client: rdb}
}

func (s *RedisStore) Set(ctx context.Context, key, value string) error {
	if s.client == nil {
		return errors.New("redis client is nil")
	}

	return s.client.Set(ctx,key, value, 0).Err()
}

func (s *RedisStore) Get(ctx context.Context, key, value string) (string,error) {
	if s.client == nil {
		return "" , errors.New("redis client is nil")
	}

	return s.client.Get(ctx, key).Result()
}