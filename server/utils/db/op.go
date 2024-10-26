package db

import (
	"context"
	"strconv"
	"time"
)

const (
	FILE_GROUP_UNIQUE_ID = "file_group_unique_id"
	FILE_MAP_NAMESPACE = "file_map:"
	FILE_NAMESPACE    = "files:"
	FILE_UNIQUE_ID = "file_unique_id"
	STORAGE_UNIQUE_ID = "storage_unique_id"
)

// Func to check db whether alive, via ping/pong
func CheckDbAlive(ctx context.Context) error {
	_, err := rdb.Ping(ctx).Result()
	return err
}

// Func to delete a key.
func Del(ctx context.Context, key string) error {
	if err := CheckDbAlive(ctx); err != nil {
		return err
	}
	return rdb.Del(ctx, key).Err()
}

// Get value from redis
func Get(ctx context.Context, key string) (string, error) {
	if err := CheckDbAlive(ctx); err != nil {
		return "", err
	}
    val, err := rdb.Get(ctx, key).Result()
	if err != nil {
        return "", err
    }
    return val, nil
}

// Get the id from map namespace
func GetIdByPath(ctx context.Context, p string) (int64, error) {
	val, err := Get(ctx, FILE_MAP_NAMESPACE + p)
	if err != nil {
        return 0, err
    }
    return strconv.ParseInt(val, 10, 64)
}

// Set a pair of key-value to redis
func Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	if err := CheckDbAlive(ctx); err != nil {
		return err
	}
	return rdb.Set(ctx, key, value, expiration).Err()
}
