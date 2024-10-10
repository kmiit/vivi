package main

import (
    "flag"
    "io/ioutil"
    "log"

    . "github.com/kmiit/vivi/types"
    "github.com/kmiit/vivi/utils"

    "github.com/pelletier/go-toml/v2"
)

func main () {
    var configPath string
    
    flag.StringVar(&configPath, "c", "", "配置文件路径")
    flag.Parse()
    
    if configPath == "" {
		log.Fatalln(0, "No config file specified!")
	}
	
	configFile, err := ioutil.ReadFile(configPath)
	if err != nil {
	    log.Fatalln(1, err)
	}

    var config ServerConfig
    err = toml.Unmarshal(configFile, &config)
    if err != nil {
        log.Fatalln(2, err)
    }
    
    utils.RunServer(config)
}
