package db

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	FILE_GROUP_UNIQUE_ID = "file_group_unique_id"
	FILE_MAP_NAMESPACE = "file_map:"
	FILE_NAMESPACE    = "files:"
	FILE_UNIQUE_ID = "file_unique_id"
	STORAGE_UNIQUE_ID = "storage_unique_id"
)

// Set a pair of key-value to redis
func Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	if err := CheckDbAlive(ctx); err != nil {
		return err
	}
	if err:= rdb.Set(ctx, key, value, expiration).Err(); err != nil {
        return err
    }
	return nil
}

// Get value from redis
func Get(ctx context.Context, key string) (string, error) {
	if err := CheckDbAlive(ctx); err != nil {
		return "", err
	}
    val, err := rdb.Get(ctx, key).Result()
    if err == redis.Nil {
        return "", fmt.Errorf("key does not exist")
    } else if err != nil {
        return "", err
    }
    return val, nil
}

func CheckDbAlive(ctx context.Context) error {
	_, err := rdb.Ping(ctx).Result()
	return err
}


