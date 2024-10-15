package db

import (
	"context"

	"github.com/kmiit/vivi/utils/config"
	"github.com/kmiit/vivi/utils/log"
	"github.com/redis/go-redis/v9"
)

const TAG = "Database"

func InitDatabase() {
	address := config.DatabaseConfig.DbAddress + ":" + config.DatabaseConfig.DbPort

	rdb := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: config.DatabaseConfig.DbPassword,
		DB:       config.DatabaseConfig.DbNumber,
	})
	log.I(TAG, "Connected to Redis", rdb)

	ctx := context.Background()
	r, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.F(TAG, err)
	} else {
		log.I(TAG, "Ping: ", r)
	}
}
