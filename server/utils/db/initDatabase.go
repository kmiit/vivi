package db

import (
	"github.com/kmiit/vivi/utils/config"
	"github.com/kmiit/vivi/utils/log"
	"github.com/redis/go-redis/v9"
)

const TAG = "Database"

var RDB *redis.Client

func InitDatabase() {
	address := config.DatabaseConfig.DbAddress + ":" + config.DatabaseConfig.DbPort

	RDB = redis.NewClient(&redis.Options{
		Addr:     address,
		Password: config.DatabaseConfig.DbPassword,
		DB:       config.DatabaseConfig.DbNumber,
	})
	log.I(TAG, "Connected to Redis", RDB)
}
