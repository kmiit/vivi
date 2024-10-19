package db

import (
	"context"
	"time"

	"github.com/kmiit/vivi/utils/config"
	"github.com/kmiit/vivi/utils/log"
	"github.com/redis/go-redis/v9"
)

const TAG = "Database"

var rdb *redis.Client

func InitDatabase() {
	address := config.DatabaseConfig.DbAddress + ":" + config.DatabaseConfig.DbPort
	ctx := context.Background()
	rdb = redis.NewClient(&redis.Options{
		Addr:     address,
		Password: config.DatabaseConfig.DbPassword,
		DB:       config.DatabaseConfig.DbNumber,
	})
	for {	
		if err := CheckDbAlive(ctx); err != nil {
			log.E(TAG, "Check alive failed: ", err)
		} else {
			log.I(TAG, "Database ready")
			break
		}
		log.W(TAG, "Database not ready, retrying...")
		time.Sleep(5 * time.Second)
	}
}
