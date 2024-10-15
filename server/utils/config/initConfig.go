package config

import (
	"github.com/kmiit/vivi/cmd/flags"
	"github.com/kmiit/vivi/types"
	"github.com/kmiit/vivi/utils/log"
)

const TAG = "Config"

var (
	Config         types.Config
	DatabaseConfig types.DatabaseConfig
	ServerConfig   types.ServerConfig
	StorageConfig  types.StorageConfig
)

func InitConfig() {
	log.I(TAG, "Initializing config")
	Config = parse(flags.ConfigFile)
	DatabaseConfig = Config.Database
	ServerConfig = Config.Server
	StorageConfig = Config.Storage
}
