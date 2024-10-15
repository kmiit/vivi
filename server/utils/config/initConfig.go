package config

import (
	"github.com/kmiit/vivi/cmd/flags"
	"github.com/kmiit/vivi/types"
	"github.com/kmiit/vivi/utils/log"
)

const TAG = "Config"

var (
	Config       types.Config
	ServerConfig types.ServerConfig
)

func InitConfig() {
	log.I(TAG, "Initializing config")
	Config = parse(flags.ConfigFile)
	ServerConfig = Config.Server
}
