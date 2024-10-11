package main

import (
	"flag"
	"log"
	"os"

	"github.com/kmiit/vivi/types"

	"github.com/pelletier/go-toml/v2"
)

func main() {
	var configPath string

	flag.StringVar(&configPath, "c", "", "配置文件路径")
	flag.Parse()

	if configPath == "" {
		log.Fatalln(-1, "No config file specified!")
	}

	configFile, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalln(1, err)
	}

	var config types.ServerConfig
	err = toml.Unmarshal(configFile, &config)
	if err != nil {
		log.Fatalln(2, err)
	}

	run(config)
}
