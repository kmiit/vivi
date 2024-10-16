package db

import (
	"context"
	"time"

	"github.com/kmiit/vivi/utils/config"
	"github.com/kmiit/vivi/utils/log"
	"github.com/redis/go-redis/v9"
)

const TAG = "Database"

var RDB *redis.Client

func InitDatabase() {
	address := config.DatabaseConfig.DbAddress + ":" + config.DatabaseConfig.DbPort
	ctx := context.Background()
	RDB = redis.NewClient(&redis.Options{
		Addr:     address,
		Password: config.DatabaseConfig.DbPassword,
		DB:       config.DatabaseConfig.DbNumber,
	})
	for {
		_, err := RDB.Ping(ctx).Result()
		if err != nil {
			log.E(TAG, "Ping: ", err)
		} else {
			log.I(TAG, "Database ready")
			break
		}
		log.W(TAG, "Database not ready, retrying...")
		time.Sleep(5 * time.Second)
	}
}
