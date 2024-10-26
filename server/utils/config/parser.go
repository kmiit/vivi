package config

import (
	"os"

	"github.com/kmiit/vivi/types"
	"github.com/kmiit/vivi/utils/log"
	"github.com/pelletier/go-toml/v2"
)

func parse(c string) types.Config {
	if c == "" {
		log.F(TAG, "Config file not specified")
	}

	cByte, err := os.ReadFile(c)
	if err != nil {
		log.F(TAG, err)
	}

	var config types.Config
	err = toml.Unmarshal(cByte, &config)
	if err != nil {
		log.F(TAG, err)
	}

	return config
}
