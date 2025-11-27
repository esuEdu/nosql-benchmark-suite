package db

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisDB struct {
	Addr string
	RDB  *redis.Client
}

func NewRedis(addr string) (*RedisDB, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return &RedisDB{
		Addr: addr,
		RDB:  rdb,
	}, nil
}

func (r *RedisDB) Name() string { return "redis" }

func (r *RedisDB) WriteTest(n int) (time.Duration, error) {
	ctx := context.Background()
	start := time.Now()
	for i := 0; i < n; i++ {
		key := fmt.Sprintf("bench:key:%d", i)
		if err := r.RDB.Set(ctx, key, fmt.Sprintf("value-%d", i), 0).Err(); err != nil {
			return 0, err
		}
	}
	return time.Since(start), nil
}

func (r *RedisDB) ReadTest(n int) (time.Duration, error) {
	ctx := context.Background()
	start := time.Now()
	for i := 0; i < n; i++ {
		key := fmt.Sprintf("bench:key:%d", i)
		_ = r.RDB.Get(ctx, key).Val()
	}
	return time.Since(start), nil
}
