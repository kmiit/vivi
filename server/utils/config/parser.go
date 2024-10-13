package config

import (
	"log"
	"os"

	"github.com/kmiit/vivi/types"

	"github.com/pelletier/go-toml/v2"
)

func Parse(c string) (types.ServerConfig) {
	if c == "" {
		log.Fatalln(-1, "Config file not specified")
	}

	cByte, err := os.ReadFile(c)
	if err != nil {
		log.Fatalln(1, err)
	}

	var config types.ServerConfig
	err = toml.Unmarshal(cByte, &config)
	if err != nil {
		log.Fatalln(2, err)
	}

	return config
}